package behavior

import (
	. "github.com/heartlhj/go-learning/workflow/engine"
	"github.com/heartlhj/go-learning/workflow/engine/utils"
)

type TakeOutgoingSequenceFlowsOperation struct {
	AbstractOperation
	EvaluateConditions bool
}

func (task TakeOutgoingSequenceFlowsOperation) Run() {
	currentFlowElement := task.getCurrentFlowElement()
	_, ok := currentFlowElement.(SequenceFlow)
	if ok {
		task.handleSequenceFlow()
	} else {
		task.handleFlowNode()
	}
}

//处理节点
func (task TakeOutgoingSequenceFlowsOperation) handleFlowNode() {
	execution := task.Execution
	currentFlowElement := task.Execution.GetCurrentFlowElement()
	task.handleActivityEnd(currentFlowElement)
	gateway, ok := currentFlowElement.(Gateway)
	var defaultSequenceFlowId = ""
	if ok {
		defaultSequenceFlowId = gateway.DefaultFlow
	}
	flowElements := currentFlowElement.GetOutgoing()
	var outgoingSequenceFlows = make([]FlowElement, 0)
	if len(flowElements) > 0 {
		for _, flowElement := range flowElements {
			sequenceFlow := (*flowElement).(SequenceFlow)
			if !task.EvaluateConditions || utils.HasTrueCondition(sequenceFlow, execution) {
				outgoingSequenceFlows = append(outgoingSequenceFlows, *flowElement)
			}
		}
		if outgoingSequenceFlows != nil && len(outgoingSequenceFlows) == 0 {
			if defaultSequenceFlowId != "" {
				for _, flowElement := range flowElements {
					if defaultSequenceFlowId == (*flowElement).GetId() {
						outgoingSequenceFlows = append(outgoingSequenceFlows, *flowElement)
					}
				}
			}
		}
	}

	if len(outgoingSequenceFlows) == 0 {
		if flowElements == nil || len(flowElements) == 0 {
			GetAgenda().PlanEndExecutionOperation(execution)
		} else {
			panic("No outgoing sequence flow of element")
		}
	} else {
		for _, outgoingExecution := range outgoingSequenceFlows {
			execution.SetCurrentFlowElement(outgoingExecution)
			GetAgenda().PlanContinueProcessOperation(execution)
		}
	}
}

//处理连线
func (task TakeOutgoingSequenceFlowsOperation) handleSequenceFlow() {
	GetAgenda().PlanContinueProcessOperation(task.Execution)
}

//获取当前活动节点
func (task TakeOutgoingSequenceFlowsOperation) getCurrentFlowElement() FlowElement {
	execution := task.Execution
	currentFlowElement := execution.GetCurrentFlowElement()
	if currentFlowElement != nil {
		return currentFlowElement
	}
	return nil
}

func (task TakeOutgoingSequenceFlowsOperation) handleActivityEnd(element FlowElement) {
	historicActinstManager := GetHistoricActinstManager()
	historicActinstManager.RecordTaskCreated(element, task.Execution)
}
