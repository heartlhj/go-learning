package interceptor

import (
	"github.com/heartlhj/go-learning/workflow/engine"
	"github.com/heartlhj/go-learning/workflow/engine/persistence"
)

type CommandContext struct {
	Command                    Command
	Agenda                     engine.ActivitiEngineAgenda
	ProcessEngineConfiguration ProcessEngineConfiguration
}

func GetProcessInstanceManager() persistence.ProcessInstanceManager {
	return persistence.ProcessInstanceManager{}
}

func GetTaskManager() persistence.TaskManager {
	return persistence.TaskManager{}
}
