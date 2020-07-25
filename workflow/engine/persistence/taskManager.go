package persistence

import (
	"github.com/heartlhj/go-learning/workflow/db"
	"github.com/heartlhj/go-learning/workflow/engine"
	. "github.com/heartlhj/go-learning/workflow/model"
	"github.com/prometheus/common/log"
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
	actinst, err := historicActinstManager.FindUnfinishedHistoricActivityInstancesByExecutionAndActivityId(entity.GetProcessInstanceId(), task.Id)
	if err != nil {
		actinst.TaskId = task.Id
		actinst.Assignee = task.Assignee
		historicActinstManager.HistoricActinst = actinst
		historicActinstManager.Update()
	}
}

func (taskManager TaskManager) createHistoricTask(task *Task) HistoricTask {
	historicTask := HistoricTask{}
	historicTask.TaskEntity = task.TaskEntity
	historicTask.TaskId = task.Id
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

func (taskManager TaskManager) DeleteTask(taskId int64) {
	task := Task{}
	_, err := db.MasterDB.Id(taskId).Delete(task)
	if err != nil {
		log.Infoln("Delete Task Err ", err)
	}
}
