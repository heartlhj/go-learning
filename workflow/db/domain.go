package db

import "encoding/xml"

var (
	flowMap = make(map[string]flow, 0)
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
	Folw             flow
}

type FlowElement struct {
	Id   string `xml:"id,attr"`
	Name string `xml:"name,attr"`
}

type Flow struct {
	FlowElement
	Id           string `xml:"id,attr"`
	Name         string `xml:"name,attr"`
	IncomingFlow *[]flow
	OutgoingFlow *[]flow
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
	Flow
	SequenceFlowName xml.Name `xml:"sequenceFlow"`
	Id               string   `xml:"id,attr"`
	SourceRef        string   `xml:"sourceRef,attr"`
	TargetRef        string   `xml:"targetRef,attr"`
}
type ExclusiveGateway struct {
	Flow
}

type flow interface {
	setIncoming(f []flow)
}

func (flow *Flow) setIncoming(f []flow) {
	flow.IncomingFlow = &f
}
func (flow *Flow) setOutgoing(f []flow) {
	flow.OutgoingFlow = &f
}

func setFow() {
	task := UserTask{}
	sequence := SequenceFlow{}
	m := []flow{}
	m[0] = &sequence
	task.setIncoming(m)
}

func Converter(d *Definitions) {
	processes := d.Process
	if processes != nil {
		for _, p := range processes {
			start := p.StartEvent
			if start != nil {
				for _, sta := range start {
					flowMap[sta.Id] = &sta
				}
			}
			se := p.SequenceFlow
			if se != nil {
				for _, s := range se {
					flowMap[s.Id] = &s
				}
			}
			user := p.UserTask
			if user != nil {
				for _, u := range user {
					flowMap[u.Id] = &u
				}
			}
			end := p.EndEvent
			if end != nil {
				for _, e := range end {
					flowMap[e.Id] = &e
				}
			}
		}
	}
}

func convertXMLToElement(model *Definitions) {
	processes := model.Process
	if processes != nil {
		for _, p := range processes {
			start := p.StartEvent
			if start != nil {
				for _, sta := range start {
					value := flowMap[sta.Id]
					m := []flow{}
					m[0] = value
					sta.setIncoming(m)
				}
			}
			se := p.SequenceFlow
			if se != nil {
				for _, s := range se {
					flowMap[s.Id] = &s
				}
			}
			user := p.UserTask
			if user != nil {
				for _, u := range user {
					flowMap[u.Id] = &u
				}
			}
			end := p.EndEvent
			if end != nil {
				for _, e := range end {
					flowMap[e.Id] = &e
				}
			}
		}
	}
}
