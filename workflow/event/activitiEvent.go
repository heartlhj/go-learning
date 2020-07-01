package event

type ActivitiEventType string

type ActivitiEvent interface {
	GetType() ActivitiEventType
}
