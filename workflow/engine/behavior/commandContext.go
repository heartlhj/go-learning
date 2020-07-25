package behavior

import (
	"github.com/heartlhj/go-learning/workflow/engine"
	. "github.com/heartlhj/go-learning/workflow/engine/persistence"
)

type CommandContext struct {
	Command                    Command
	Agenda                     engine.ActivitiEngineAgenda
	ProcessEngineConfiguration ProcessEngineConfiguration
}

func GetProcessInstanceManager() ProcessInstanceManager {
	return ProcessInstanceManager{}
}

func GetTaskManager() TaskManager {
	return TaskManager{}
}

func GetDefineManager() DefineManager {
	return DefineManager{}
}
func GetVariableManager() VariableManager {
	return VariableManager{}
}

func GetIdentityLinkManager() IdentityLinkManager {
	return IdentityLinkManager{}
}

func GetHistoricActinstManager() HistoricActinstManager {
	return HistoricActinstManager{}
}

func GetHistoricTaskManager() HistoricTaskManager {
	return HistoricTaskManager{}
}

func GetHistoricProcessManager() HistoricProcessManager {
	return HistoricProcessManager{}
}
