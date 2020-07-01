package model

type ExecutionEntityImpl struct {
	BusinessKey        string
	CurrentFlowElement FlowElement
	DeploymentId       int
	ProcessInstanceId  int
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

func (execution ExecutionEntityImpl) GetDeploymentId() int {
	return execution.DeploymentId
}

func (execution ExecutionEntityImpl) SetDeploymentId(deploymentId int) {
	execution.DeploymentId = deploymentId
}

func (execution ExecutionEntityImpl) GetProcessInstanceId() int {
	return execution.ProcessInstanceId
}

func (execution ExecutionEntityImpl) SetProcessInstanceId(processInstanceId int) {
	execution.ProcessInstanceId = processInstanceId
}
