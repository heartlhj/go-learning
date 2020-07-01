package agenda

import . "github.com/heartlhj/go-learning/workflow/model"

type ActivitiEngineAgenda interface {
	PlanContinueProcessOperation(execution ExecutionEntity)

	//planContinueProcessSynchronousOperation(execution ExecutionEntity)
	//
	//planContinueProcessInCompensation(execution ExecutionEntity)
	//
	//planContinueMultiInstanceOperation(execution ExecutionEntity)

	PlanTakeOutgoingSequenceFlowsOperation(execution ExecutionEntity, evaluateConditions bool)
	//
	//planEndExecutionOperation(execution ExecutionEntity)
	PlanTriggerExecutionOperation(execution ExecutionEntity)
	//
	//planDestroyScopeOperation(execution ExecutionEntity)
	//
	//planExecuteInactiveBehaviorsOperation()

}
