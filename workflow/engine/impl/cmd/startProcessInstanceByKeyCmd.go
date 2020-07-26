package cmd

import (
	"github.com/heartlhj/go-learning/workflow/engine"
	"github.com/heartlhj/go-learning/workflow/engine/behavior"
	. "github.com/heartlhj/go-learning/workflow/engine/entityImpl"
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
	instance := ProcessInstance{}
	instance.BusinessKey = start.BusinessKey
	instance.TenantId = start.TenantId
	instance.StartTime = time.Now()
	instance.Key = process.Id
	instance.Name = process.Name
	instance.ProcessDefineId = bytearries[0].Id
	//生成流程实例
	manager := ProcessInstanceManager{Instance: &instance}
	manager.CreateProcessInstance()
	//获取开始节点
	flowElement := process.InitialFlowElement
	element := flowElement.(engine.StartEvent)
	execution := ExecutionEntityImpl{ProcessInstanceId: instance.Id}
	execution.SetCurrentFlowElement(element)
	execution.SetProcessDefineId(bytearries[0].Id)
	execution.SetCurrentActivityId(element.GetId())
	//保存流程变量
	SetVariable(&execution, start.Variables)
	context, e := behavior.GetCommandContext()
	if e != nil {

	}
	context.Agenda.PlanContinueProcessOperation(&execution)
	return process
}
