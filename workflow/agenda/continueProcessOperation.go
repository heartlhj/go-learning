package agenda

import (
	"github.com/heartlhj/go-learning/workflow/context"
	. "github.com/heartlhj/go-learning/workflow/model"
)

type ContinueProcessOperation struct {
	AbstractOperation
}

func (cont *ContinueProcessOperation) run() {
	element := cont.Execution.GetCurrentFlowElement()
	if element != nil {
		flow, ok := element.(SequenceFlow)
		if !ok {
			cont.continueThroughFlowNode(element)
		} else {
			cont.continueThroughSequenceFlow(flow)
		}
	}
}

func (cont *ContinueProcessOperation) continueThroughSequenceFlow(sequenceFlow SequenceFlow) {
	flowElement := sequenceFlow.TargetFlowElement
	cont.Execution.SetCurrentFlowElement(*flowElement)
	context.GetAgenda().PlanContinueProcessOperation(cont.Execution)
}

func (cont *ContinueProcessOperation) continueThroughFlowNode(element FlowElement) {
	behavior := element.GetBehavior()
	behavior.Execute(cont.Execution)
}
