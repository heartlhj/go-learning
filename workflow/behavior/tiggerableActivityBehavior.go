package behavior

import . "github.com/heartlhj/go-learning/workflow/model"

type TriggerableActivityBehavior interface {
	Trigger(entity ExecutionEntity)
}
