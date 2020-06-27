package agenda

import . "github.com/heartlhj/go-learning/workflow/model"

type ContinueProcessOperation struct {
	AbstractOperation
}

func (cont ContinueProcessOperation) run() {
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

func (cont ContinueProcessOperation) continueThroughSequenceFlow(flow SequenceFlow) {

}

func (cont ContinueProcessOperation) continueThroughFlowNode(element FlowElement) {
	behavior := element.GetBehavior()
	behavior.Execute(cont.Execution)
}
