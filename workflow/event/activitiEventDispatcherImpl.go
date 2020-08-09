package event

var eventDispatcher ActivitiEventDispatcherImpl

type ActivitiEventDispatcherImpl struct {
	EventSupport *ActivitiEventSupport
	Enabled      bool
}

func SetEventDispatcher(eventDispatcher ActivitiEventDispatcher) {
	eventDispatcher = eventDispatcher
}

func GetEventDispatcher() ActivitiEventDispatcher {
	return eventDispatcher
}
func (eventDispatcher ActivitiEventDispatcherImpl) AddEventListener(listenerToAdd ActivitiEventListener) {
	eventDispatcher.EventSupport.AddEventListener(listenerToAdd)
}

func (eventDispatcher ActivitiEventDispatcherImpl) RemoveEventListener(listenerToRemove ActivitiEventListener) {
}

func (eventDispatcher ActivitiEventDispatcherImpl) DispatchEvent(event ActivitiEvent) {
	if eventDispatcher.Enabled {
		eventDispatcher.EventSupport.DispatchEvent(event)
	}
}

func (eventDispatcher ActivitiEventDispatcherImpl) SetEnabled(enabled bool) {
	eventDispatcher.Enabled = enabled
}

func (eventDispatcher ActivitiEventDispatcherImpl) IsEnabled() bool {
	return eventDispatcher.Enabled
}
