package entity

import (
	. "github.com/heartlhj/go-learning/workflow/engine/behavior"
	. "github.com/heartlhj/go-learning/workflow/engine/persistence"
	. "github.com/heartlhj/go-learning/workflow/engine/variable"
)

type TaskEntityImpl struct {
	ExecutionEntityImpl
	TaskId    int64
	Variables map[string]interface{}
}

func (task *TaskEntityImpl) GetTaskId() int64 {
	return task.TaskId
}

func (task *TaskEntityImpl) SetTaskId(taskId int64) {
	task.TaskId = taskId
}

func (task *TaskEntityImpl) GetVariable() map[string]interface{} {
	variableManager := GetVariableManager()
	variables, err := variableManager.SelectByTaskId(task.TaskId)
	if err != nil {
		return task.HandleVariable(variables)
	}
	return nil
}

func (task *TaskEntityImpl) getSpecificVariable(variableName string, variableManager VariableManager) (Variable, error) {
	return variableManager.SelectTaskId(variableName, task.TaskId)
}

func (task *TaskEntityImpl) SetScope(variable *Variable) {
	variable.TaskId = task.TaskId
}
