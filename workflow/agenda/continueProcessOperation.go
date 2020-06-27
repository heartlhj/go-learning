package agenda

type ContinueProcessOperation struct {
	AbstractOperation
}

func (cont ContinueProcessOperation) run() {
	element := cont.Execution.GetCurrentFlowElement()
	if element != nil {

	}
}
