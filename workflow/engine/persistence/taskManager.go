package persistence

import (
	"github.com/heartlhj/go-learning/workflow/db"
	"github.com/heartlhj/go-learning/workflow/engine"
	. "github.com/heartlhj/go-learning/workflow/model"
	"github.com/prometheus/common/log"
	"time"
)

type TaskManager struct {
	Task *Task
}

func (taskManager TaskManager) Insert(execution engine.ExecutionEntity) {
	_, err := db.MasterDB.Insert(taskManager.Task)
	if err != nil {
		log.Infoln("Create Task Err ", err)
	}
	taskManager.recordTaskCreated(taskManager.Task, execution)
}

func (taskManager TaskManager) recordTaskCreated(task *Task, entity engine.ExecutionEntity) {
	historicTaskManager := HistoricTaskManager{}
	historicTask := taskManager.createHistoricTask(task)
	historicTaskManager.HistoricTask = historicTask
	historicTaskManager.Insert()

	historicActinstManager := HistoricActinstManager{}
	actinst, err := historicActinstManager.FindUnfinishedHistoricActivityInstancesByExecutionAndActivityId(entity.GetProcessInstanceId(), task.TaskDefineKey)
	if err == nil {
		actinst.Assignee = task.Assignee
		actinst.TaskId = task.Id
		historicActinstManager.HistoricActinst = actinst
		historicActinstManager.Update()
	}
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

func (taskManager TaskManager) FindById(taskId int) []Task {
	task := make([]Task, 0)
	err := db.MasterDB.Where("id=?", taskId).Find(&task)
	if err != nil {
		log.Infoln("Select FindById Err ", err)
	}
	return task
}

func (taskManager TaskManager) FindByProcessInstanceId(processInstanceId int64) []Task {
	task := make([]Task, 0)
	err := db.MasterDB.Where("proc_inst_id=?", processInstanceId).Find(&task)
	if err != nil {
		log.Infoln("Select FindByProcessInstanceId err ", err)
	}
	return task
}

func (taskManager TaskManager) DeleteTask(task Task) {
	_, err := db.MasterDB.Id(task.Id).Delete(task)
	if err != nil {
		log.Infoln("Delete Task Err ", err)
	}
	identityLinkManager := IdentityLinkManager{}
	identityLinks, err := identityLinkManager.SelectByTaskId(task.Id)
	if err == nil {
		for _, identityLink := range identityLinks {
			identityLinkManager.Delete(identityLink.Id)
		}
	}
	variableManager := VariableManager{}
	variables, err := variableManager.SelectByTaskId(task.Id)
	if err == nil {
		for _, variable := range variables {
			variableManager.Delete(variable.Id)
		}
	}
	recordTaskEnd(task)
}

func recordTaskEnd(task Task) {
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
	historicActinstManager.UpdateTaskId()
}
