package interceptor

import (
	. "github.com/heartlhj/go-learning/workflow/agenda"
	. "github.com/heartlhj/go-learning/workflow/persistence"
)

type CommandContext struct {
	Agenda ActivitiEngineAgenda
}

func GetTaskManager() TaskManager {
	return TaskManager{}
}
