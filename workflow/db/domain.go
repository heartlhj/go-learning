package db

import (
	"encoding/xml"
)

var (
	flowMap = make(map[string]flow, 0)
)

type Bytearry struct {
	Id    int    `json:"id" xorm:"pk autoincr"`
	Name  string `json:"name"`
	Bytes string `json:"bytes"`
}
type Definitions struct {
	DefinitionsName    xml.Name  `xml:"definitions"`
	Xmlns              string    `xml:"xmlns,attr"`
	Xsi                string    `xml:"xsi,attr"`
	Xsd                string    `xml:"xsd,attr"`
	Activiti           string    `xml:"activiti,attr"`
	Bpmndi             string    `xml:"bpmndi,attr"`
	Omgdc              string    `xml:"omgdc,attr"`
	Omgdi              string    `xml:"omgdi,attr"`
	TypeLanguage       string    `xml:"typeLanguage,attr"`
	RgetNamespace      string    `xml:"rgetNamespace,attr"`
	ExpressionLanguage string    `xml:"expressionLanguage,attr"`
	TargetNamespace    string    `xml:"targetNamespace,attr"`
	Process            []Process `xml:"process"`
	Message            []Message `xml:"message"`
}
type Process struct {
	ProcessName            xml.Name `xml:"process"`
	Id                     string   `xml:"id,attr"`
	Name                   string
	Documentation          string                   `xml:"documentation"`
	IsExecutable           string                   `xml:"isExecutable,attr"`
	StartEvent             []StartEvent             `xml:"startEvent"`
	EndEvent               []EndEvent               `xml:"endEvent"`
	UserTask               []UserTask               `xml:"userTask"`
	SequenceFlow           []SequenceFlow           `xml:"sequenceFlow"`
	ExclusiveGateway       []ExclusiveGateway       `xml:"exclusiveGateway"`
	InclusiveGateway       []InclusiveGateway       `xml:"inclusiveGateway"`
	ParallelGateway        []ParallelGateway        `xml:"parallelGateway"`
	BoundaryEvent          []BoundaryEvent          `xml:"boundaryEvent"`
	IntermediateCatchEvent []IntermediateCatchEvent `xml:"intermediateCatchEvent"`
	SubProcess             []SubProcess             `xml:"subProcess"`
	Flow                   []flow
}

type SubProcess struct {
	*Process
	SubProcessName xml.Name `xml:"subProcess"`
}

type Message struct {
	*BaseElement
	MessageName xml.Name `xml:"message"`
}

type BaseElement struct {
	Id   string `xml:"id,attr"`
	Name string `xml:"name,attr"`
}

type Flow struct {
	BaseElement
	Id           string `xml:"id,attr"`
	Name         string `xml:"name,attr"`
	IncomingFlow []*flow
	OutgoingFlow []*flow
}

type StartEvent struct {
	*Flow
	StartEventName xml.Name `xml:"startEvent"`
	Initiator      string   `xml:"initiator,attr"`
	FormKey        string   `xml:"formKey,attr"`
}

type EndEvent struct {
	*Flow
	EndEventName xml.Name `xml:"endEvent"`
}

type UserTask struct {
	*Flow
	UserTaskName   xml.Name `xml:"userTask"`
	Assignee       string   `xml:"assignee,attr"`
	CandidateUsers string   `xml:"candidateUsers,attr"`
}
type SequenceFlow struct {
	*Flow
	SequenceFlowName    xml.Name `xml:"sequenceFlow"`
	Id                  string   `xml:"id,attr"`
	SourceRef           string   `xml:"sourceRef,attr"`
	TargetRef           string   `xml:"targetRef,attr"`
	ConditionExpression string   `xml:"conditionExpression"`
	SourceFlowElement   *flow
	TargetFlowElement   *flow
}

type ExclusiveGateway struct {
	*Flow
}

type InclusiveGateway struct {
	*Flow
}

type ParallelGateway struct {
	*Flow
}

type BoundaryEvent struct {
	*Flow
	BoundaryEventName    xml.Name             `xml:"boundaryEvent"`
	AttachedToRef        string               `xml:"attachedToRef,attr"`
	CancelActivity       string               `xml:"cancelActivity,attr"`
	TimerEventDefinition TimerEventDefinition `xml:"timerEventDefinition"`
}

type TimerEventDefinition struct {
	TimerEventDefinitionName xml.Name `xml:"timerEventDefinition"`
	TimeDuration             string   `xml:"timeDuration"`
}

type IntermediateCatchEvent struct {
	*Flow
	IntermediateCatchEventName xml.Name               `xml:"intermediateCatchEvent"`
	MessageEventDefinition     MessageEventDefinition `xml:"messageEventDefinition"`
}

type MessageEventDefinition struct {
	MessageEventDefinitionName xml.Name `xml:"messageEventDefinition"`
	MessageRef                 string   `xml:"messageRef,attr"`
}
type flow interface {
	setIncoming(f []*flow)
	setOutgoing(f []*flow)
	getIncoming() []*flow
	getOutgoing() []*flow
}

func (flow *Flow) setIncoming(f []*flow) {
	flow.IncomingFlow = f
}
func (flow *Flow) setOutgoing(f []*flow) {
	flow.OutgoingFlow = f
}

func (flow *Flow) getIncoming() []*flow {
	return flow.IncomingFlow
}
func (flow *Flow) getOutgoing() []*flow {
	return flow.OutgoingFlow
}

func Converter(d *Definitions) {
	processes := d.Process
	if processes != nil {
		for j, p := range processes {
			start := p.StartEvent
			if start != nil {
				for i, sta := range start {
					flowMap[sta.Id] = start[i]
				}
			}
			se := p.SequenceFlow
			if se != nil {
				for i, s := range se {
					flowMap[s.Id] = se[i]
				}
			}
			user := p.UserTask
			if user != nil {
				for i, u := range user {
					flowMap[u.Id] = user[i]
				}
			}
			end := p.EndEvent
			if end != nil {
				for i, e := range end {
					flowMap[e.Id] = end[i]
				}
			}
			flows := make([]flow, len(flowMap))
			count := 0
			for _, v := range flowMap {
				flows[count] = v
				count++
			}
			processes[j].Flow = flows
		}
	}
}

func ConvertXMLToElement(model *Definitions) {
	processes := model.Process
	if processes != nil {
		for _, p := range processes {
			flows := p.Flow
			for i := range flows {
				value, ok := flows[i].(SequenceFlow)
				if ok {
					SourceRef := value.SourceRef
					//上一个节点
					lastFlow := flowMap[SourceRef]
					if lastFlow != nil {
						var outgoing = lastFlow.getOutgoing()
						if outgoing == nil {
							outgoing = make([]*flow, 0)
						}
						newOut := append(outgoing, &flows[i])
						//设置上一个节点出口
						lastFlow.setOutgoing(newOut)
						//设置当前连线入口
						value.SourceFlowElement = &lastFlow

					}
					//下一个节点
					TargetRef := value.TargetRef
					nextFlow := flowMap[TargetRef]
					if nextFlow != nil {
						incoming := nextFlow.getIncoming()
						if incoming == nil {
							incoming = make([]*flow, 0)
						}
						newIn := append(incoming, &flows[i])
						m := make([]*flow, 1)
						m[0] = &nextFlow
						//设置当前连线出口
						value.TargetFlowElement = &nextFlow
						//设置写一个节点入口
						nextFlow.setIncoming(newIn)
					}
					flows[i] = &value
				}
			}
		}
	}
}
