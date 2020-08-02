package behavior

import (
	"github.com/heartlhj/go-learning/workflow/engine"
)

type FlowNodeActivityBehavior interface {
	Leave(execution engine.ExecutionEntity) error
}
