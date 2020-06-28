package behavior

import (
	. "github.com/heartlhj/go-learning/workflow/model"
)

type ExclusiveGatewayActivityBehavior struct {
}

//排他网关
func (exclusive ExclusiveGatewayActivityBehavior) Execute(execution ExecutionEntity) {
	exclusive.Leave(execution)
}

func (exclusive ExclusiveGatewayActivityBehavior) Leave(execution ExecutionEntity) {

}
