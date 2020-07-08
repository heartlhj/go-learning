package engine

type ExecutionEntity interface {
	SetBusinessKey(businessKey string)

	GetCurrentFlowElement() FlowElement

	SetCurrentFlowElement(flow FlowElement)

	GetDeploymentId() int

	SetDeploymentId(deploymentId int)

	GetProcessInstanceId() int64

	SetProcessInstanceId(processInstanceId int64)
}
