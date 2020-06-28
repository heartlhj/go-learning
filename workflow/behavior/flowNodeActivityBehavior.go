package behavior

import . "github.com/heartlhj/go-learning/workflow/model"

type FlowNodeActivityBehavior interface {
	Leave(execution ExecutionEntity)
}
