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

//任务出口执行
func (agenda *DefaultActivitiEngineAgenda) PlanTriggerExecutionOperation(execution ExecutionEntity) {
	agenda.planOperation(&TriggerExecutionOperation{AbstractOperation{Execution: execution}})
}

//任务出口执行
func (agenda *DefaultActivitiEngineAgenda) PlanTakeOutgoingSequenceFlowsOperation(execution ExecutionEntity, valuateConditions bool) {
	agenda.planOperation(&TakeOutgoingSequenceFlowsOperation{AbstractOperation{Execution: execution}})
}

//任务出口执行
func (agenda *DefaultActivitiEngineAgenda) PlanEndExecutionOperation(execution ExecutionEntity, valuateConditions bool) {
	agenda.planOperation(&TakeOutgoingSequenceFlowsOperation{AbstractOperation{Execution: execution}})
}
