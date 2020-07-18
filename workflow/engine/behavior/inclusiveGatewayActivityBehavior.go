package behavior

import (
	"github.com/heartlhj/go-learning/workflow/engine"
	"github.com/heartlhj/go-learning/workflow/engine/utils"
)

type InclusiveGatewayActivityBehavior struct {
}

//包容网关
func (exclusive InclusiveGatewayActivityBehavior) Execute(execution engine.ExecutionEntity) {
	exclusive.Leave(execution)
}

func (exclusive InclusiveGatewayActivityBehavior) Leave(execution engine.ExecutionEntity) {
	processInstanceId := execution.GetProcessInstanceId()
	taskManager := GetTaskManager()
	//查询当前执行节点
	tasks := taskManager.FindByProcessInstanceId(processInstanceId)
	var oneExecutionCanReachGateway = false
	if tasks != nil && len(tasks) > 0 {
		bytearry := GetDefineManager().GetBytearry(execution.GetProcessDefineId())
		process := GetBpmn(bytearry)
		for _, task := range tasks {
			if task.TaskDefineKey != execution.GetCurrentActivityId() {
				oneExecutionCanReachGateway = utils.IsReachable(process, task.TaskDefineKey, execution.GetCurrentActivityId())
			} else {
				oneExecutionCanReachGateway = true
			}
		}
	}
	if !oneExecutionCanReachGateway {
		GetAgenda().PlanTakeOutgoingSequenceFlowsOperation(execution, true)
	}
}
