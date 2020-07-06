package behavior

type EndExecutionOperation struct {
	AbstractOperation
}

func (end EndExecutionOperation) run() {
	manager := GetProcessInstanceManager()
	manager.DeleteProcessInstance(end.Execution.GetProcessInstanceId())
}
