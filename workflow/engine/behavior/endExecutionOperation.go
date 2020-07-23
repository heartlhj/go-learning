package behavior

import "github.com/heartlhj/go-learning/workflow/engine"

type EndExecutionOperation struct {
	AbstractOperation
}

func (end *EndExecutionOperation) Run() {
	deleteDataForExecution(end.Execution)
	manager := GetProcessInstanceManager()
	manager.DeleteProcessInstance(end.Execution.GetProcessInstanceId())
}

func deleteDataForExecution(entity engine.ExecutionEntity) {
	taskManager := GetTaskManager()
	tasks := taskManager.FindByProcessInstanceId(entity.GetProcessInstanceId())
	if tasks != nil && len(tasks) > 0 {
		for _, task := range tasks {
			taskManager.DeleteTask(task.Id)
		}
	}
	variableManager := GetVariableManager()
	variables, err := variableManager.SelectByProcessInstanceId(entity.GetProcessInstanceId())
	if err == nil {
		for _, variable := range variables {
			variableManager.Delete(variable.Id)
		}
	}
}
