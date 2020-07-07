package behavior

import (
	"github.com/heartlhj/go-learning/workflow/engine"
)

type TriggerExecutionOperation struct {
	AbstractOperation
}

func (trigger TriggerExecutionOperation) Run() {
	element := trigger.getCurrentFlowElement(trigger.Execution)
	behavior := element.GetBehavior()
	operation := behavior.(TriggerableActivityBehavior)
	operation.Trigger(trigger.Execution)
}

func (trigger TriggerExecutionOperation) getCurrentFlowElement(execut engine.ExecutionEntity) engine.FlowElement {
	return nil
}