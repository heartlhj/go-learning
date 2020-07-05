package agenda

import (
	"github.com/heartlhj/go-learning/workflow/engine/interceptor"
)

type EndExecutionOperation struct {
	AbstractOperation
}

func (end EndExecutionOperation) run() {
	manager := interceptor.GetProcessInstanceManager()
	manager.DeleteProcessInstance(end.Execution.GetProcessInstanceId())
}
