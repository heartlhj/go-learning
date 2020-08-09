package persistence

import (
	"github.com/heartlhj/go-learning/workflow/db"
	"github.com/heartlhj/go-learning/workflow/engine"
	"github.com/heartlhj/go-learning/workflow/errs"
	. "github.com/heartlhj/go-learning/workflow/model"
	"github.com/prometheus/common/log"
	"time"
)

type TaskManager struct {
	Task *Task
}

func (taskManager TaskManager) Insert(execution engine.ExecutionEntity) (err error) {
	err = db.DB().Create(taskManager.Task).Error
	if err == nil {
		err = taskManager.recordTaskCreated(taskManager.Task, execution)
	}
	//dispatcher := event.GetEventDispatcher()
	//dispatcher.DispatchEvent(CreateEntityEvent())
	return err
}

func (taskManager TaskManager) recordTaskCreated(task *Task, entity engine.ExecutionEntity) (err error) {
	historicTaskManager := HistoricTaskManager{}
	historicTask := taskManager.createHistoricTask(task)
	historicTaskManager.HistoricTask = historicTask
	err = historicTaskManager.Insert()
	if err != nil {
		historicActinstManager := HistoricActinstManager{}
		actinst, err := historicActinstManager.FindUnfinishedHistoricActivityInstancesByExecutionAndActivityId(entity.GetProcessInstanceId(), task.TaskDefineKey)
		if err == nil {
			actinst.Assignee = task.Assignee
			actinst.TaskId = task.Id
			historicActinstManager.HistoricActinst = actinst
			err = historicActinstManager.Update()
		}
	}
	return err
}

func (taskManager TaskManager) createHistoricTask(task *Task) HistoricTask {
	historicTask := HistoricTask{}
	//historicTask.TaskEntity = task.TaskEntity
	historicTask.TaskId = task.Id
	historicTask.ProcessInstanceId = task.ProcessInstanceId
	historicTask.StartTime = task.StartTime
	historicTask.TenantId = task.TenantId
	historicTask.Assignee = task.Assignee
	historicTask.TaskDefineKey = task.TaskDefineKey
	historicTask.DeploymentId = task.DeploymentId
	historicTask.TaskDefineName = task.TaskDefineName
	return historicTask
}

func (taskManager TaskManager) FindById(taskId int) (Task, error) {
	task := Task{}
	err := db.DB().Where("id= ?", taskId).First(&task).Error
	if err != nil {
		log.Infoln("Select FindById Err ", err)
		return task, err
	}
	return task, nil
}

func (taskManager TaskManager) FindByProcessInstanceId(processInstanceId int64) (task []Task, err error) {
	task = make([]Task, 0)
	err = db.DB().Where("proc_inst_id=?", processInstanceId).Find(&task).Error
	if err != nil {
		log.Infoln("Select FindByProcessInstanceId err ", err)
	}
	if task == nil || len(task) <= 0 {
		return task, errs.ProcessError{Code: "1001", Msg: "Not find"}
	}
	return task, err
}

func (taskManager TaskManager) DeleteTask(task Task) (err error) {
	err = db.DB().Where("id = ?", task.Id).Delete(&task).Error
	if err != nil {
		return err
	}
	identityLinkManager := IdentityLinkManager{}
	identityLinks, errSelect := identityLinkManager.SelectByTaskId(task.Id)
	if errSelect == nil {
		for _, identityLink := range identityLinks {
			identityLinkManager.Delete(identityLink.Id)
		}
	}
	variableManager := VariableManager{}
	variables, errSelect := variableManager.SelectByTaskId(task.Id)
	if errSelect == nil {
		for _, variable := range variables {
			variableManager.Delete(variable.Id)
		}
	}
	err = recordTaskEnd(task)
	return err
}

func recordTaskEnd(task Task) (err error) {
	historicTaskManager := HistoricTaskManager{}
	historicTask := HistoricTask{}
	historicTask.TaskId = task.Id
	historicTask.EndTime = time.Now()
	historicTaskManager.HistoricTask = historicTask
	historicTaskManager.MarkEnded()

	historicActinst := HistoricActinst{}
	historicActinst.EndTime = historicTask.EndTime
	historicActinst.TaskId = historicTask.TaskId
	historicActinstManager := HistoricActinstManager{}
	historicActinstManager.HistoricActinst = historicActinst
	return historicActinstManager.UpdateTaskId()
}
