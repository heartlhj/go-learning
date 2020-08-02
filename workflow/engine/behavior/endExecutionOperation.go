package behavior

import "github.com/heartlhj/go-learning/workflow/engine"

type EndExecutionOperation struct {
	AbstractOperation
}

func (end *EndExecutionOperation) Run() (err error) {
	err = deleteDataForExecution(end.Execution)
	if err != nil {
		return err
	}
	manager := GetProcessInstanceManager()
	err = manager.DeleteProcessInstance(end.Execution.GetProcessInstanceId())
	return err
}

func deleteDataForExecution(entity engine.ExecutionEntity) (err error) {
	taskManager := GetTaskManager()
	tasks, errSelect := taskManager.FindByProcessInstanceId(entity.GetProcessInstanceId())
	if errSelect == nil {
		for _, task := range tasks {
			taskManager.DeleteTask(task)
		}
	}

	identityLinkManager := GetIdentityLinkManager()
	identityLinks, errSelect := identityLinkManager.SelectByProcessInstanceId(entity.GetProcessInstanceId())
	if errSelect == nil {
		for _, identityLink := range identityLinks {
			identityLinkManager.Delete(identityLink.Id)
		}
	}
	variableManager := GetVariableManager()
	variables, err := variableManager.SelectByProcessInstanceId(entity.GetProcessInstanceId())
	if err == nil {
		for _, variable := range variables {
			variableManager.Delete(variable.Id)
		}
	}
	return err
}
