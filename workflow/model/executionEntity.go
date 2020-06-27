package model

type ExecutionEntity interface {
	SetBusinessKey(businessKey string)

	GetCurrentFlowElement() FlowElement

	SetCurrentFlowElement(flow *FlowElement)
}
