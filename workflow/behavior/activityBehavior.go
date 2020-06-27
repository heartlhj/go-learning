package behavior

import . "github.com/heartlhj/go-learning/workflow/model"

type ActivityBehavior interface {
	execute(execution ExecutionEntity)
}
