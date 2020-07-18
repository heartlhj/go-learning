package engine

type ExecutionEntity interface {
	SetBusinessKey(businessKey string)

	GetCurrentFlowElement() FlowElement

	SetCurrentFlowElement(flow FlowElement)

	GetDeploymentId() int

	SetDeploymentId(deploymentId int)

	GetProcessInstanceId() int64

	SetProcessInstanceId(processInstanceId int64)

	GetProcessDefineId() int64

	SetProcessDefineId(processDefineId int64)

	GetCurrentActivityId() string

	SetCurrentActivityId(currentActivityId string)

	SetVariable(variables map[string]interface{}) error

	GetVariable() map[string]interface{}

	GetTaskId() int64

	SetTaskId(taskId int64)
}
