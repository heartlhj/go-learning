package engine

import (
	"github.com/heartlhj/go-learning/workflow/context"
	. "github.com/heartlhj/go-learning/workflow/interceptor"
	. "github.com/heartlhj/go-learning/workflow/model"
)

type TaskServiceImpl struct {
}

//流程审批完成
func (task TaskServiceImpl) Complete(taskId int, variables map[string]interface{}) Task {

	manager := GetTaskManager()
	tasks := manager.FindById(taskId)
	if len(tasks) > 0 {
		task := tasks[0]
		executeTaskComplete(task)
		return task
	}
	return Task{}
}

func executeTaskComplete(task Task) {
	manager := GetTaskManager()
	manager.DeleteTask(task.Id)
	execution := ExecutionEntityImpl{}
	context.GetAgenda().PlanTriggerExecutionOperation(execution)
}
