package behavior

import (
	"github.com/heartlhj/go-learning/workflow/engine"
)

type TriggerableActivityBehavior interface {
	Trigger(entity engine.ExecutionEntity)
}
