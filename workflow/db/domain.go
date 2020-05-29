package db

import "encoding/xml"

var (
	flowMap = make([]Flow, 0)
)

type Bytearry struct {
	Id    int    `json:"id" xorm:"pk autoincr"`
	Name  string `json:"name"`
	Bytes string `json:"bytes"`
}
type Definitions struct {
	Definitionsname    xml.Name  `xml:"definitions"`
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
}
type Process struct {
	ProcessName      xml.Name           `xml:"process"`
	Id               string             `xml:"id,attr"`
	Name             string             `xml:"name,attr"`
	IsExecutable     string             `xml:"isExecutable,attr"`
	StartEvent       []StartEvent       `xml:"startEvent"`
	EndEvent         []EndEvent         `xml:"endEvent"`
	UserTask         []UserTask         `xml:"userTask"`
	SequenceFlow     []SequenceFlow     `xml:"sequenceFlow"`
	ExclusiveGateway []ExclusiveGateway `xml:"ExclusiveGateway"`
}

type FlowElement struct {
	Id   string `xml:"id,attr"`
	Name string `xml:"name,attr"`
}

type Flow struct {
	FlowElement
	Id           string `xml:"id,attr"`
	Name         string `xml:"name,attr"`
	IncomingFlow *FlowElement
	OutgoingFlow *FlowElement
}

type StartEvent struct {
	Flow
	StartEventName xml.Name `xml:"startEvent"`
	Initiator      string   `xml:"initiator,attr"`
	FormKey        string   `xml:"formKey,attr"`
}

type EndEvent struct {
	Flow
	EndEventName xml.Name `xml:"endEvent"`
}

type UserTask struct {
	Flow
	UserTaskName   xml.Name `xml:"userTask"`
	Assignee       string   `xml:"assignee,attr"`
	CandidateUsers string   `xml:"candidateUsers,attr"`
}
type SequenceFlow struct {
	SequenceFlowName xml.Name `xml:"sequenceFlow"`
	Id               string   `xml:"id,attr"`
	SourceRef        string   `xml:"sourceRef,attr"`
	TargetRef        string   `xml:"targetRef,attr"`
}
type ExclusiveGateway struct {
	Flow
}

func (flow Flow) setIncoming(f *FlowElement) {
	flow.IncomingFlow = f
}
func (flow Flow) setOutgoing(f *FlowElement) {
	flow.OutgoingFlow = f
}

func setFow() {

}
