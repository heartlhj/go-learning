package engine

import (
	. "github.com/heartlhj/go-learning/workflow/model"
	. "github.com/heartlhj/go-learning/workflow/persistence"
)

type RuntimeServiceImpl struct {
}

//发起流程
func (runtime RuntimeServiceImpl) StartProcessInstanceByKey(processDefinitionKey string, variables map[string]interface{},
	businessKey string, tenantId string) {
	process := FindDeployedProcessDefinitionByKey(processDefinitionKey)
	instance := ProcessInstance{BusinessKey: businessKey, TenantId: tenantId}
	CreateProcessInstance(&instance)
	flowElement := process.InitialFlowElement
	element := flowElement.(StartEvent)
	element.GetOutgoing()

}
