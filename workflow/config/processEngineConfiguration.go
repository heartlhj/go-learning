package config

import (
	. "github.com/heartlhj/go-learning/workflow/event"
	. "github.com/heartlhj/go-learning/workflow/interceptor"
)

var EventListeners = []ActivitiEventListener{}

type ProcessEngineConfiguration struct {
}

func init() {
	processEngineConfiguration := ProcessEngineConfiguration{}
	initCommandContext(processEngineConfiguration)
}

func initCommandContext(configuration ProcessEngineConfiguration) {
	context := CommandContext{}
	context.SetProcessEngineConfiguration(configuration)
}

func AddEventListeners(eventListeners []ActivitiEventListener) {
	if len(eventListeners) > 0 {
		for _, listener := range eventListeners {
			EventListeners = append(EventListeners, listener)
		}
	}

}
