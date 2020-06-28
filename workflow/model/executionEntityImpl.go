package model

type ExecutionEntityImpl struct {
	BusinessKey        string
	CurrentFlowElement FlowElement
	DeploymentId       string
}

func (execution ExecutionEntityImpl) SetBusinessKey(businessKey string) {
	execution.BusinessKey = businessKey
}

func (execution ExecutionEntityImpl) GetCurrentFlowElement() FlowElement {
	return execution.CurrentFlowElement
}

func (execution ExecutionEntityImpl) SetCurrentFlowElement(flow FlowElement) {
	execution.CurrentFlowElement = flow
}

func (execution ExecutionEntityImpl) GetDeploymentId() string {
	return execution.DeploymentId
}

func (execution ExecutionEntityImpl) SetDeploymentId(deploymentId string) {
	execution.DeploymentId = deploymentId
}
