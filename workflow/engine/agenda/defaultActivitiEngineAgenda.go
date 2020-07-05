package agenda

import (
	"container/list"
	"github.com/heartlhj/go-learning/workflow/engine"
)

type DefaultActivitiEngineAgenda struct {
	Operations list.List
}

//设置后续操作
func (agenda *DefaultActivitiEngineAgenda) PlanOperation(operation engine.Operation) {
	agenda.Operations.PushFront(operation)
}

//设置后续操作
func (agenda *DefaultActivitiEngineAgenda) GetNextOperation() engine.Operation {
	value := agenda.Operations.Front()
	return value.Value.(engine.Operation)
}

func (agenda *DefaultActivitiEngineAgenda) getNextOperation() engine.Operation {
	return agenda.Operations.Front().Value.(engine.Operation)
}

//连线继续执行
func (agenda *DefaultActivitiEngineAgenda) PlanContinueProcessOperation(execution engine.ExecutionEntity) {
	agenda.PlanOperation(&ContinueProcessOperation{AbstractOperation{Execution: execution}})
}

//任务出口执行
func (agenda *DefaultActivitiEngineAgenda) PlanTriggerExecutionOperation(execution engine.ExecutionEntity) {
	agenda.PlanOperation(&TriggerExecutionOperation{AbstractOperation{Execution: execution}})
}

//任务出口执行
func (agenda *DefaultActivitiEngineAgenda) PlanTakeOutgoingSequenceFlowsOperation(execution engine.ExecutionEntity, valuateConditions bool) {
	agenda.PlanOperation(&TakeOutgoingSequenceFlowsOperation{AbstractOperation{Execution: execution}})
}

//任务出口执行
func (agenda *DefaultActivitiEngineAgenda) PlanEndExecutionOperation(execution engine.ExecutionEntity) {
	agenda.PlanOperation(&TakeOutgoingSequenceFlowsOperation{AbstractOperation{Execution: execution}})
}
