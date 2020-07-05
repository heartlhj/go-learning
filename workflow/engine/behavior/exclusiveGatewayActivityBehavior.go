package behavior

import (
	"github.com/heartlhj/go-learning/workflow/engine"
)

type ExclusiveGatewayActivityBehavior struct {
}

//排他网关
func (exclusive ExclusiveGatewayActivityBehavior) Execute(execution engine.ExecutionEntity) {
	exclusive.Leave(execution)
}

func (exclusive ExclusiveGatewayActivityBehavior) Leave(execution engine.ExecutionEntity) {

}
