package agenda

import (
	. "github.com/heartlhj/go-learning/workflow/interceptor"
	. "github.com/heartlhj/go-learning/workflow/model"
)

type AbstractOperation struct {
	CommandContext CommandContext
	Execution      ExecutionEntity
}
