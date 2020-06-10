package model

import (
	"encoding/xml"
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
	Flow                   []FlowElement
	InitialFlowElement     FlowElement
}

//子流程
type SubProcess struct {
	*Process
	SubProcessName xml.Name `xml:"subProcess"`
}

//消息订阅
type Message struct {
	*BaseElement
	MessageName xml.Name `xml:"message"`
}

//通用字段
type BaseElement struct {
	Id   string `xml:"id,attr"`
	Name string `xml:"name,attr"`
}

//父类实现体
type Flow struct {
	BaseElement
	Id                string `xml:"id,attr"`
	Name              string `xml:"name,attr"`
	IncomingFlow      []*FlowElement
	OutgoingFlow      []*FlowElement
	SourceFlowElement *FlowElement
	TargetFlowElement *FlowElement
}

//开始节点
type StartEvent struct {
	*Flow
	StartEventName xml.Name `xml:"startEvent"`
	Initiator      string   `xml:"initiator,attr"`
	FormKey        string   `xml:"formKey,attr"`
}

//结束节点
type EndEvent struct {
	*Flow
	EndEventName xml.Name `xml:"endEvent"`
}

//用户任务
type UserTask struct {
	*Flow
	UserTaskName   xml.Name `xml:"userTask"`
	Assignee       string   `xml:"assignee,attr"`
	CandidateUsers string   `xml:"candidateUsers,attr"`
}

//连线
type SequenceFlow struct {
	*Flow
	SequenceFlowName    xml.Name `xml:"sequenceFlow"`
	Id                  string   `xml:"id,attr"`
	SourceRef           string   `xml:"sourceRef,attr"`
	TargetRef           string   `xml:"targetRef,attr"`
	ConditionExpression string   `xml:"conditionExpression"`
}

//排他网关
type ExclusiveGateway struct {
	*Flow
}

//包容网关
type InclusiveGateway struct {
	*Flow
}

//并行网关
type ParallelGateway struct {
	*Flow
}

//边界事件
type BoundaryEvent struct {
	*Flow
	BoundaryEventName    xml.Name             `xml:"boundaryEvent"`
	AttachedToRef        string               `xml:"attachedToRef,attr"`
	CancelActivity       string               `xml:"cancelActivity,attr"`
	TimerEventDefinition TimerEventDefinition `xml:"timerEventDefinition"`
}

//定时任务
type TimerEventDefinition struct {
	TimerEventDefinitionName xml.Name `xml:"timerEventDefinition"`
	TimeDuration             string   `xml:"timeDuration"`
}

//中间抛出事件
type IntermediateCatchEvent struct {
	*Flow
	IntermediateCatchEventName xml.Name               `xml:"intermediateCatchEvent"`
	MessageEventDefinition     MessageEventDefinition `xml:"messageEventDefinition"`
}

//消息事件
type MessageEventDefinition struct {
	MessageEventDefinitionName xml.Name `xml:"messageEventDefinition"`
	MessageRef                 string   `xml:"messageRef,attr"`
}

//接口
type FlowElement interface {
	SetIncoming(f []*FlowElement)
	SetOutgoing(f []*FlowElement)
	GetIncoming() []*FlowElement
	GetOutgoing() []*FlowElement

	SetSourceFlowElement(f *FlowElement)
	SetTargetFlowElement(f *FlowElement)
	GetSourceFlowElement() *FlowElement
	GetTargetFlowElement() *FlowElement
}

func (flow *Flow) SetIncoming(f []*FlowElement) {
	flow.IncomingFlow = f
}
func (flow *Flow) SetOutgoing(f []*FlowElement) {
	flow.OutgoingFlow = f
}

func (flow *Flow) GetIncoming() []*FlowElement {
	return flow.IncomingFlow
}
func (flow *Flow) GetOutgoing() []*FlowElement {
	return flow.OutgoingFlow
}

func (flow *Flow) SetSourceFlowElement(f *FlowElement) {
	flow.SourceFlowElement = f
}
func (flow *Flow) SetTargetFlowElement(f *FlowElement) {
	flow.TargetFlowElement = f
}

func (flow *Flow) GetSourceFlowElement() *FlowElement {
	return flow.SourceFlowElement
}
func (flow *Flow) GetTargetFlowElement() *FlowElement {
	return flow.TargetFlowElement
}
