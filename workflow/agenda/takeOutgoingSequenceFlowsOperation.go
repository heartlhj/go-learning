package agenda

import (
	"github.com/heartlhj/go-learning/workflow/context"
	. "github.com/heartlhj/go-learning/workflow/model"
)

type TakeOutgoingSequenceFlowsOperation struct {
	AbstractOperation
}

func (task TakeOutgoingSequenceFlowsOperation) run() {
	execution := task.Execution
	element := task.getCurrentFlowElement(execution)
	flowElements := element.GetOutgoing()
	if len(flowElements) > 0 {
		for _, flowElement := range flowElements {
			execution.SetCurrentFlowElement(*flowElement)
			context.GetAgenda().PlanContinueProcessOperation(execution)
		}
	} else {

	}
}

func (task TakeOutgoingSequenceFlowsOperation) getCurrentFlowElement(execut ExecutionEntity) FlowElement {
	return nil
}
