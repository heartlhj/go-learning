package cmd

import (
	"encoding/xml"
	"github.com/heartlhj/go-learning/workflow/engine"
	"github.com/heartlhj/go-learning/workflow/engine/behavior"

	//"github.com/heartlhj/go-learning/workflow/engine/behavior"
	"github.com/heartlhj/go-learning/workflow/engine/entity"
	. "github.com/heartlhj/go-learning/workflow/engine/interceptor"
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

func (start StartProcessInstanceByKeyCmd) Execute(interceptor CommandContext) interface{} {
	bytearries := FindDeployedProcessDefinitionByKey(start.ProcessDefinitionKey)
	//解析xml数据
	define := new(engine.Definitions)
	xml.Unmarshal([]byte(bytearries[0].Bytes), &define)
	behavior.Converter(define)
	instance := ProcessInstance{BusinessKey: start.BusinessKey, TenantId: start.TenantId, StartTime: time.Now()}
	manager := ProcessInstanceManager{Instance: &instance}
	manager.CreateProcessInstance()
	process := define.Process[0]
	flowElement := process.InitialFlowElement
	element := flowElement.(engine.StartEvent)
	outgoing := element.GetOutgoing()
	execution := entity.ExecutionEntityImpl{}
	execution.SetCurrentFlowElement(*outgoing[0])
	context, e := GetCommandContext()
	if e != nil {

	}
	context.Agenda.PlanContinueProcessOperation(execution)
	return process
}
