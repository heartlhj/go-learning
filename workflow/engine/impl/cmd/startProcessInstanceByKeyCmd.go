package cmd

import (
	"github.com/heartlhj/go-learning/workflow/engine"
	"github.com/heartlhj/go-learning/workflow/engine/behavior"
	"github.com/heartlhj/go-learning/workflow/engine/entity"
	. "github.com/heartlhj/go-learning/workflow/engine/persistence"
	. "github.com/heartlhj/go-learning/workflow/model"
	"time"
)

type StartProcessInstanceByKeyCmd struct {
	ProcessDefinitionKey string
	Variables            map[string]interface{}
	BusinessKey          string
	TenantId             string
}

func (start StartProcessInstanceByKeyCmd) Execute(interceptor behavior.CommandContext) interface{} {
	defineManager := behavior.GetDefineManager()
	bytearries := defineManager.FindDeployedProcessDefinitionByKey(start.ProcessDefinitionKey)
	//解析xml数据
	process := behavior.GetBpmn(*bytearries[0])
	instance := ProcessInstance{BusinessKey: start.BusinessKey, TenantId: start.TenantId, StartTime: time.Now()}
	instance.Key = process.Id
	instance.Name = process.Name
	manager := ProcessInstanceManager{Instance: &instance}
	manager.CreateProcessInstance()
	flowElement := process.InitialFlowElement
	element := flowElement.(engine.StartEvent)
	outgoing := element.GetOutgoing()
	execution := entity.ExecutionEntityImpl{ProcessInstanceId: instance.Id}
	execution.SetCurrentFlowElement(*outgoing[0])
	execution.SetProcessDefineId(bytearries[0].Id)
	execution.SetCurrentActivityId((*outgoing[0]).GetId())
	execution.SetVariable(start.Variables)
	context, e := behavior.GetCommandContext()
	if e != nil {

	}
	context.Agenda.PlanContinueProcessOperation(&execution)
	return process
}
