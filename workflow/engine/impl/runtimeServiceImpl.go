package engine

import (
	"github.com/heartlhj/go-learning/workflow/engine/impl/cmd"
	. "github.com/heartlhj/go-learning/workflow/engine/interceptor"
)

type RuntimeService struct {
	ServiceImpl
}

//发起流程
func (runtime RuntimeService) StartProcessInstanceByKey(processDefinitionKey string, variables map[string]interface{},
	businessKey string, tenantId string) {
	GetServiceImpl().CommandExecutor.Exe(cmd.StartProcessInstanceByKeyCmd{ProcessDefinitionKey: processDefinitionKey,
		Variables: variables, TenantId: tenantId, BusinessKey: businessKey})
}
