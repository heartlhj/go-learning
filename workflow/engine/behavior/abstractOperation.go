package behavior

import (
	"github.com/heartlhj/go-learning/workflow/engine"
)

type AbstractOperation struct {
	CommandContext CommandContext
	Execution      engine.ExecutionEntity
}
