package agenda

import . "github.com/heartlhj/go-learning/workflow/interceptor"

type EndExecutionOperation struct {
	AbstractOperation
}

func (end EndExecutionOperation) run() {
	manager := GetProcessInstanceManager()
	manager.DeleteProcessInstance(end.Execution.GetProcessInstanceId())
}
