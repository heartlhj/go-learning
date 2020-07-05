package agenda

import (
	"github.com/heartlhj/go-learning/workflow/engine"
	"github.com/heartlhj/go-learning/workflow/engine/interceptor"
)

type ContinueProcessOperation struct {
	AbstractOperation
}

func (cont *ContinueProcessOperation) Run() {
	element := cont.Execution.GetCurrentFlowElement()
	if element != nil {
		flow, ok := element.(engine.SequenceFlow)
		if !ok {
			cont.continueThroughFlowNode(element)
		} else {
			cont.continueThroughSequenceFlow(flow)
		}
	}
}

func (cont *ContinueProcessOperation) continueThroughSequenceFlow(sequenceFlow engine.SequenceFlow) {
	flowElement := sequenceFlow.TargetFlowElement
	cont.Execution.SetCurrentFlowElement(*flowElement)
	interceptor.GetAgenda().PlanContinueProcessOperation(cont.Execution)
}

func (cont *ContinueProcessOperation) continueThroughFlowNode(element engine.FlowElement) {
	behavior := element.GetBehavior()
	behavior.Execute(cont.Execution)
}
