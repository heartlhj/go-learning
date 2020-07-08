package behavior

import (
	"github.com/heartlhj/go-learning/workflow/engine"
)

type TakeOutgoingSequenceFlowsOperation struct {
	AbstractOperation
}

func (task TakeOutgoingSequenceFlowsOperation) Run() {
	execution := task.Execution
	element := task.getCurrentFlowElement(execution)
	flowElements := element.GetOutgoing()
	if len(flowElements) > 0 {
		for _, flowElement := range flowElements {
			execution.SetCurrentFlowElement(*flowElement)
			GetAgenda().PlanContinueProcessOperation(execution)
		}
	} else {

	}
}

func (task TakeOutgoingSequenceFlowsOperation) getCurrentFlowElement(execut engine.ExecutionEntity) engine.FlowElement {
	currentFlowElement := execut.GetCurrentFlowElement()
	if currentFlowElement != nil {
		return currentFlowElement
	}
	return nil
}
