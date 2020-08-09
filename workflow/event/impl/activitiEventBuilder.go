package impl

import (
	"github.com/heartlhj/go-learning/workflow/errs"
	. "github.com/heartlhj/go-learning/workflow/event"
)

type ActivitiEventBuilder struct {
}

func CreateEvent() ActivitiEvent {
	return nil
}

func CreateEntityEvent(eventType ActivitiEventType, entity interface{}) (ActivitiEntityEvent, error) {
	entityEventImpl := ActivitiEntityEventImpl{}
	entityEventImpl.ActivitiEventImpl = ActivitiEventImpl{}
	entityEventImpl.EventType = eventType
	var err error = nil
	if entity == nil {
		err = errs.ProcessError{Msg: "Entity cannot be null."}
	}
	entityEventImpl.Entity = entity
	return entityEventImpl, err
}
