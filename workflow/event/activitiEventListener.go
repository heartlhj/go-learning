package event

type ActivitiEventListener interface {
	onEvent(event ActivitiEvent)
}
