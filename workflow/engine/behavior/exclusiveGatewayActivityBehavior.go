package behavior

import (
	"github.com/heartlhj/go-learning/workflow/engine"
)

type ExclusiveGatewayActivityBehavior struct {
}

//排他网关
func (exclusive ExclusiveGatewayActivityBehavior) Execute(execution engine.ExecutionEntity) {
	processInstanceId := execution.GetProcessInstanceId()
	taskManager := GetTaskManager()
	//查询当前执行节点
	tasks := taskManager.FindByProcessInstanceId(processInstanceId)
	if tasks != nil && len(tasks) > 0 {

	}
	exclusive.Leave(execution)
}

func (exclusive ExclusiveGatewayActivityBehavior) Leave(execution engine.ExecutionEntity) {

}
