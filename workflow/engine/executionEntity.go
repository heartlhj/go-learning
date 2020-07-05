package engine

type ExecutionEntity interface {
	SetBusinessKey(businessKey string)

	GetCurrentFlowElement() FlowElement

	SetCurrentFlowElement(flow FlowElement)

	GetDeploymentId() int

	SetDeploymentId(deploymentId int)

	GetProcessInstanceId() int

	SetProcessInstanceId(processInstanceId int)
}
