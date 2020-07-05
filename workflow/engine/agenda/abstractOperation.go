package agenda

import (
	"github.com/heartlhj/go-learning/workflow/engine"
	"github.com/heartlhj/go-learning/workflow/engine/interceptor"
)

type AbstractOperation struct {
	CommandContext interceptor.CommandContext
	Execution      engine.ExecutionEntity
}
