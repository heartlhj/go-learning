package behavior

import . "github.com/heartlhj/go-learning/workflow/model"

type ActivityBehavior interface {
	Execute(execution ExecutionEntity)
}
