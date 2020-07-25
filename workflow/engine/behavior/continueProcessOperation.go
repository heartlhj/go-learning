package behavior

import (
	"github.com/heartlhj/go-learning/workflow/engine"
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
	GetAgenda().PlanContinueProcessOperation(cont.Execution)
}

func (cont *ContinueProcessOperation) continueThroughFlowNode(element engine.FlowElement) {
	historicActinstManager := GetHistoricActinstManager()
	historicActinstManager.RecordActivityStart(cont.Execution)
	behavior := element.GetBehavior()
	if behavior != nil {
		behavior.Execute(cont.Execution)
	} else {
		GetAgenda().PlanTakeOutgoingSequenceFlowsOperation(cont.Execution, true)
	}

}
