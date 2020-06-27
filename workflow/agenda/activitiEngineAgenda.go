package agenda

import . "github.com/heartlhj/go-learning/workflow/model"

type ActivitiEngineAgenda interface {
	PlanContinueProcessOperation(execution ExecutionEntity)

	//planContinueProcessSynchronousOperation(execution ExecutionEntity)
	//
	//planContinueProcessInCompensation(execution ExecutionEntity)
	//
	//planContinueMultiInstanceOperation(execution ExecutionEntity)
	//
	//planTakeOutgoingSequenceFlowsOperation(execution ExecutionEntity, evaluateConditions bool )
	//
	//planEndExecutionOperation(execution ExecutionEntity)
	//
	//planTriggerExecutionOperation(execution ExecutionEntity)
	//
	//planDestroyScopeOperation(execution ExecutionEntity)
	//
	//planExecuteInactiveBehaviorsOperation()

}
