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
	defineManager := behavior.GetDefineManager()
	bytearry := defineManager.FindProcessByTask(task.ProcessInstanceId)
	currentTask := behavior.FindCurrentTask(bytearry.Bytes, task.TaskDefineKey)
	execution := entity.ExecutionEntityImpl{}
	execution.SetCurrentFlowElement(currentTask)
	execution.SetProcessInstanceId(task.ProcessInstanceId)
	interceptor.Agenda.PlanTriggerExecutionOperation(&execution)
}
