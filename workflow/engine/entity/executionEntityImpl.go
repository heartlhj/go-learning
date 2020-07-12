package entity

import (
	"github.com/heartlhj/go-learning/workflow/engine"
	. "github.com/heartlhj/go-learning/workflow/engine/behavior"
	. "github.com/heartlhj/go-learning/workflow/engine/persistence"
	. "github.com/heartlhj/go-learning/workflow/engine/variable"
	"github.com/heartlhj/go-learning/workflow/errs"
	"reflect"
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

func (execution *ExecutionEntityImpl) SetVariable(variables map[string]interface{}) error {
	engineConfiguration := GetProcessEngineConfiguration()
	variableTypes := engineConfiguration.VariableTypes
	variableManager := GetVariableManager()
	if variables != nil && len(variables) > 0 {
		for k, v := range variables {
			kind := reflect.TypeOf(v).Kind()
			variableType := variableTypes.GetVariableType(kind.String())
			if variableType == nil {
				return errs.ProcessError{Code: "1001", Msg: "no type"}
			}
			variable := variableManager.Create(k, variableType, v)
			//存在更新
			specificVariable, e := execution.getSpecificVariable(k, variableManager)
			if e != nil {
				variable.Version = specificVariable.Version + 1
			}
			variable.ProcessInstanceId = execution.ProcessInstanceId
			variableManager.Insert(variable)
		}
	}
	return nil
}

func (execution *ExecutionEntityImpl) getSpecificVariable(variableName string, variableManager VariableManager) (Variable, error) {
	return variableManager.SelectProcessInstanceId(variableName, execution.ProcessInstanceId)
}
