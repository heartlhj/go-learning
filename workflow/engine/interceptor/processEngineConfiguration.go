package interceptor

import (
	. "github.com/heartlhj/go-learning/workflow/event"
)

var processEngineConfiguration ProcessEngineConfiguration

type ProcessEngineConfiguration struct {
	CommandInvoker        CommandInterceptor
	CommandInterceptors   []CommandInterceptor
	EventListeners        []ActivitiEventListener
	Service               ServiceImpl
	CommandExecutor       CommandExecutor
	CommandContextFactory CommandContextFactory
}

func init() {
	processEngineConfiguration = ProcessEngineConfiguration{}
	initCommandContextFactory()
	initCommandInvoker()
	initCommandInterceptors()
	initCommandExecutor()
	initService()
	initCommandContext(processEngineConfiguration)
}

func initCommandContext(configuration ProcessEngineConfiguration) {
	//context := CommandContext{}
}

func AddEventListeners(eventListeners []ActivitiEventListener) {
	var EventListeners []ActivitiEventListener
	if len(eventListeners) > 0 {
		for _, listener := range eventListeners {
			EventListeners = append(EventListeners, listener)
		}
	}
	processEngineConfiguration.EventListeners = EventListeners
}

func getDefaultCommandInterceptors() []CommandInterceptor {
	var interceptors []CommandInterceptor
	interceptors = append(interceptors, &CommandContextInterceptor{CommandContextFactory: processEngineConfiguration.CommandContextFactory})
	//interceptors = append(interceptors, CommandInvoker{})
	return interceptors
}
func initCommandInvoker() {
	commandInvoker := CommandInvoker{}
	processEngineConfiguration.CommandInvoker = &commandInvoker
}

func initCommandInterceptors() {
	interceptors := getDefaultCommandInterceptors()
	interceptors = append(interceptors, processEngineConfiguration.CommandInvoker)
	processEngineConfiguration.CommandInterceptors = interceptors
}

func initCommandExecutor() {
	if processEngineConfiguration.CommandExecutor == nil {
		first := initInterceptorChain(processEngineConfiguration.CommandInterceptors)
		commandExecutor := CommandExecutorImpl{First: first}
		processEngineConfiguration.CommandExecutor = commandExecutor
	}
}
func initService() {
	serviceImpl := ServiceImpl{CommandExecutor: processEngineConfiguration.CommandExecutor}
	SetServiceImpl(serviceImpl)
	processEngineConfiguration.Service = serviceImpl
}

func initInterceptorChain(interceptors []CommandInterceptor) CommandInterceptor {
	if len(interceptors) > 0 {
		for i := 0; i < len(interceptors)-1; i++ {
			interceptor := interceptors[i]
			interceptor.SetNext(interceptors[i+1])
		}
	}
	return interceptors[0]
}

func initCommandContextFactory() {
	factory := CommandContextFactory{}
	processEngineConfiguration.CommandContextFactory = factory
}