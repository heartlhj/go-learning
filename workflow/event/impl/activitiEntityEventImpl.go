package impl

import . "github.com/heartlhj/go-learning/workflow/event"

type ActivitiEntityEventImpl struct {
	ActivitiEntityEvent
	ActivitiEventImpl
	Entity interface{}
}

func (ActivitiEntityEventImpl) GetType() ActivitiEventType {
	return TASK_CREATED
}
