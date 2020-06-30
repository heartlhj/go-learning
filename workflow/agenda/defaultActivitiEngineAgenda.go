package agenda

import (
	"container/list"
	. "github.com/heartlhj/go-learning/workflow/model"
)

type DefaultActivitiEngineAgenda struct {
	Operations list.List
}

//设置后续操作
func (agenda *DefaultActivitiEngineAgenda) planOperation(operation Operation) {
	agenda.Operations.PushFront(operation)
}

func (agenda *DefaultActivitiEngineAgenda) getNextOperation() Operation {
	return agenda.Operations.Front().Value.(Operation)
}

//连线继续执行
func (agenda *DefaultActivitiEngineAgenda) PlanContinueProcessOperation(execution ExecutionEntity) {
	agenda.planOperation(&ContinueProcessOperation{AbstractOperation{Execution: execution}})
}

func (agenda *DefaultActivitiEngineAgenda) PlanTriggerExecutionOperation(execution ExecutionEntity) {
	agenda.planOperation(&TriggerExecutionOperation{AbstractOperation{Execution: execution}})
}
