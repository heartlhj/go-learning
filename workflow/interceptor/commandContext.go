package interceptor

import (
	. "github.com/heartlhj/go-learning/workflow/agenda"
	. "github.com/heartlhj/go-learning/workflow/config"
	. "github.com/heartlhj/go-learning/workflow/persistence"
)

type CommandContext struct {
	Agenda                     ActivitiEngineAgenda
	ProcessEngineConfiguration ProcessEngineConfiguration
}

func GetTaskManager() TaskManager {
	return TaskManager{}
}

func (commandContext *CommandContext) SetProcessEngineConfiguration(processEngineConfiguration ProcessEngineConfiguration) {
	commandContext.ProcessEngineConfiguration = processEngineConfiguration
}

func (commandContext CommandContext) GetProcessEngineConfiguration() ProcessEngineConfiguration {
	return commandContext.ProcessEngineConfiguration
}
