package behavior

import (
	"github.com/heartlhj/go-learning/workflow/engine"
	"github.com/heartlhj/go-learning/workflow/engine/utils"
)

type ExclusiveGatewayActivityBehavior struct {
}

//排他网关
func (exclusive ExclusiveGatewayActivityBehavior) Execute(execution engine.ExecutionEntity) {
	exclusive.Leave(execution)
}

func (exclusive ExclusiveGatewayActivityBehavior) Leave(execution engine.ExecutionEntity) {
	element := execution.GetCurrentFlowElement()
	exclusiveGateway, ok := element.(engine.Gateway)
	var outgoingSequenceFlow *engine.FlowElement
	var defaultSequenceFlow *engine.FlowElement
	if ok {
		defaultSequenceFlowId := exclusiveGateway.DefaultFlow
		sequenceFlowIterator := exclusiveGateway.GetOutgoing()
		if sequenceFlowIterator != nil && len(sequenceFlowIterator) > 0 {
			for _, sequenceFlow := range sequenceFlowIterator {
				flow := (*sequenceFlow).(engine.SequenceFlow)
				conditionEvaluatesToTrue := utils.HasTrueCondition(flow, execution)
				if conditionEvaluatesToTrue && defaultSequenceFlowId != "" && defaultSequenceFlowId != flow.Id {
					outgoingSequenceFlow = sequenceFlow
				}
				if defaultSequenceFlowId != "" && defaultSequenceFlowId == flow.Id {
					defaultSequenceFlow = sequenceFlow
				}
			}

		}
	}
	if outgoingSequenceFlow != nil {
		execution.SetCurrentFlowElement(*outgoingSequenceFlow)
	} else {
		if defaultSequenceFlow != nil {
			execution.SetCurrentFlowElement(*defaultSequenceFlow)
		}
	}
	GetAgenda().PlanTakeOutgoingSequenceFlowsOperation(execution, true)

}
