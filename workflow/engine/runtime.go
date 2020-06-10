package engine

type Runtime interface {
	StartProcessInstanceByKey(processDefinitionKey string, variables map[string]interface{}, businessKey string, tenantId string)
}
