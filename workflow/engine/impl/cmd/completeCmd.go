package cmd

import (
	"github.com/heartlhj/go-learning/workflow/engine/behavior"
	"github.com/heartlhj/go-learning/workflow/engine/entity"
	. "github.com/heartlhj/go-learning/workflow/model"
)

type CompleteCmd struct {
	TaskId     int
	Variables  map[string]interface{}
	LocalScope bool
}

func (taskCmd CompleteCmd) Execute(interceptor behavior.CommandContext) interface{} {

	manager := behavior.GetTaskManager()
	tasks := manager.FindById(taskCmd.TaskId)
	if len(tasks) > 0 {
		task := tasks[0]
		taskCmd.executeTaskComplete(task, interceptor)
		return task
	}
	return Task{}
}

func (taskCmd CompleteCmd) executeTaskComplete(task Task, interceptor behavior.CommandContext) {
	manager := behavior.GetTaskManager()
	manager.DeleteTask(task.Id)
	defineManager := behavior.GetDefineManager()
	bytearry := defineManager.FindProcessByTask(task.ProcessInstanceId)
	currentTask := behavior.FindCurrentTask(*bytearry, task.TaskDefineKey)
	taskExecution := entity.TaskEntityImpl{}
	execution := entity.ExecutionEntityImpl{}
	execution.SetCurrentFlowElement(currentTask)
	execution.SetCurrentActivityId(task.TaskDefineKey)
	processInstanceManager := behavior.GetProcessInstanceManager()
	execution.SetProcessDefineId(processInstanceManager.GetProcessInstance(task.ProcessInstanceId).ProcessDefineId)
	execution.SetProcessInstanceId(task.ProcessInstanceId)
	taskExecution.SetTaskId(task.Id)
	taskExecution.ExecutionEntityImpl = execution
	if taskCmd.LocalScope {
		taskExecution.SetVariable(taskCmd.Variables)
	} else {
		taskExecution.ExecutionEntityImpl.SetVariable(taskCmd.Variables)
	}
	interceptor.Agenda.PlanTriggerExecutionOperation(&taskExecution)
}
