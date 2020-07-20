package behavior

type EndExecutionOperation struct {
	AbstractOperation
}

func (end *EndExecutionOperation) Run() {
	manager := GetProcessInstanceManager()
	manager.DeleteProcessInstance(end.Execution.GetProcessInstanceId())
}
