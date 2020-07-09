package entity

import (
	"github.com/heartlhj/go-learning/workflow/engine"
	. "github.com/heartlhj/go-learning/workflow/engine/behavior"
	"github.com/heartlhj/go-learning/workflow/model"
)

type ExecutionEntityImpl struct {
	BusinessKey        string
	CurrentFlowElement engine.FlowElement
	DeploymentId       int
	ProcessInstanceId  int64
}

func (execution *ExecutionEntityImpl) SetBusinessKey(businessKey string) {
	execution.BusinessKey = businessKey
}

func (execution ExecutionEntityImpl) GetCurrentFlowElement() engine.FlowElement {
	return execution.CurrentFlowElement
}

func (execution *ExecutionEntityImpl) SetCurrentFlowElement(flow engine.FlowElement) {
	execution.CurrentFlowElement = flow
}

func (execution ExecutionEntityImpl) GetDeploymentId() int {
	return execution.DeploymentId
}

func (execution *ExecutionEntityImpl) SetDeploymentId(deploymentId int) {
	execution.DeploymentId = deploymentId
}

func (execution ExecutionEntityImpl) GetProcessInstanceId() int64 {
	return execution.ProcessInstanceId
}

func (execution *ExecutionEntityImpl) SetProcessInstanceId(processInstanceId int64) {
	execution.ProcessInstanceId = processInstanceId
}

func (execution *ExecutionEntityImpl) SetVariable(variables map[string]interface{}) {
	engineConfiguration := GetProcessEngineConfiguration()
	executor := engineConfiguration.CommandExecutor
	if executor != nil {

	}
	variableManager := GetVariableManager()
	variable := model.Variable{ProcessInstanceId: execution.ProcessInstanceId}
	variableManager.Variable = variable
	variableManager.Insert("")
}
