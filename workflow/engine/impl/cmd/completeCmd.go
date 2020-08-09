package cmd

import (
	"github.com/heartlhj/go-learning/workflow/engine/behavior"
	. "github.com/heartlhj/go-learning/workflow/engine/entityImpl"
	"github.com/heartlhj/go-learning/workflow/event"
	"github.com/heartlhj/go-learning/workflow/event/impl"
	. "github.com/heartlhj/go-learning/workflow/model"
)

type CompleteCmd struct {
	TaskId     int
	Variables  map[string]interface{}
	LocalScope bool
}

func (taskCmd CompleteCmd) Execute(interceptor behavior.CommandContext) (interface{}, error) {
	manager := behavior.GetTaskManager()
	task, err := manager.FindById(taskCmd.TaskId)
	if err != nil {
		return task, err
	}
	taskCmd.executeTaskComplete(task, interceptor)
	return task, err
}

func (taskCmd CompleteCmd) executeTaskComplete(task Task, interceptor behavior.CommandContext) (err error) {

	// All properties set, now firing 'create' events
	if event.GetEventDispatcher().IsEnabled() {
		activitiEntityEvent, err := impl.CreateEntityEvent(event.TASK_COMPLETED, task)
		if err != nil {
			return err
		}
		event.GetEventDispatcher().DispatchEvent(activitiEntityEvent)
	}
	err = deleteTask(task)
	if err != nil {
		return err
	}
	defineManager := behavior.GetDefineManager()
	bytearry, err := defineManager.FindProcessByTask(task.ProcessInstanceId)
	if err != nil {
		return err
	}
	currentTask := behavior.FindCurrentTask(bytearry, task.TaskDefineKey)
	taskExecution := TaskEntityImpl{}
	execution := ExecutionEntityImpl{}
	execution.SetCurrentFlowElement(currentTask)
	execution.SetCurrentActivityId(task.TaskDefineKey)
	processInstanceManager := behavior.GetProcessInstanceManager()
	execution.SetProcessDefineId(processInstanceManager.GetProcessInstance(task.ProcessInstanceId).ProcessDefineId)
	execution.SetProcessInstanceId(task.ProcessInstanceId)
	taskExecution.SetTaskId(task.Id)
	taskExecution.ExecutionEntityImpl = execution
	if taskCmd.LocalScope {
		err = SetVariable(&taskExecution, taskCmd.Variables)
	} else {
		err = SetVariable(&execution, taskCmd.Variables)
	}
	if err != nil {
		return err
	}
	interceptor.Agenda.PlanTriggerExecutionOperation(&taskExecution)
	return nil
}

func deleteTask(task Task) (err error) {
	manager := behavior.GetTaskManager()
	return manager.DeleteTask(task)
}
