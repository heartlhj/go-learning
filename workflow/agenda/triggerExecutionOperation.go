package agenda

import (
	. "github.com/heartlhj/go-learning/workflow/behavior"
	. "github.com/heartlhj/go-learning/workflow/model"
)

type TriggerExecutionOperation struct {
	AbstractOperation
}

func (trigger TriggerExecutionOperation) run() {
	element := trigger.getCurrentFlowElement(trigger.Execution)
	behavior := element.GetBehavior()
	operation := behavior.(TriggerableActivityBehavior)
	operation.Trigger(trigger.Execution)
}

func (trigger TriggerExecutionOperation) getCurrentFlowElement(execut ExecutionEntity) FlowElement {
	return nil
}
