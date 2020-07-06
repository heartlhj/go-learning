package cmd

import (
	"github.com/heartlhj/go-learning/workflow/engine/behavior"
	"github.com/heartlhj/go-learning/workflow/engine/entity"
	. "github.com/heartlhj/go-learning/workflow/model"
)

type CompleteCmd struct {
	TaskId    int
	Variables map[string]interface{}
}

func (task CompleteCmd) Execute(interceptor behavior.CommandContext) interface{} {

	manager := behavior.GetTaskManager()
	tasks := manager.FindById(task.TaskId)
	if len(tasks) > 0 {
		task := tasks[0]
		executeTaskComplete(task, interceptor)
		return task
	}
	return Task{}
}

func executeTaskComplete(task Task, interceptor behavior.CommandContext) {
	manager := behavior.GetTaskManager()
	manager.DeleteTask(task.Id)
	execution := entity.ExecutionEntityImpl{}
	interceptor.Agenda.PlanTriggerExecutionOperation(&execution)
}
