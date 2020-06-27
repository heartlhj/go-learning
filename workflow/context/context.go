package context

import (
	"container/list"
	. "github.com/heartlhj/go-learning/workflow/agenda"
	. "github.com/heartlhj/go-learning/workflow/interceptor"
)

var (
	Stack list.List
)

type Context struct {
}

func SetCommandContext(commandContext CommandContext) {
	Stack.PushFront(commandContext)
}

func GetCommandContext() CommandContext {
	return Stack.Front().Value.(CommandContext)
}

func GetAgenda() ActivitiEngineAgenda {
	return Stack.Front().Value.(CommandContext).Agenda
}
