package behavior

import . "github.com/heartlhj/go-learning/workflow/engine"

var (
	flowMap = make(map[string]FlowElement, 0)
)

//将元素存入map
func Converter(d *Definitions) {
	processes := d.Process
	if processes != nil {
		for j, p := range processes {
			start := p.StartEvent
			if start != nil {
				for i, sta := range start {
					flowMap[sta.Id] = start[i]
					processes[j].InitialFlowElement = start[i]
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
					behavior := UserTaskActivityBehavior{UserTask: user[i]}
					user[i].SetBehavior(behavior)
					flowMap[u.Id] = user[i]
				}
			}
			end := p.EndEvent
			if end != nil {
				for i, e := range end {
					flowMap[e.Id] = end[i]
				}
			}
			flows := make([]FlowElement, len(flowMap))
			count := 0
			for _, v := range flowMap {
				flows[count] = v
				count++
			}
			processes[j].Flow = flows
		}
	}
	ConvertXMLToElement(d)
}

//设置元素的出入口
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
						var outgoing = lastFlow.GetOutgoing()
						if outgoing == nil {
							outgoing = make([]*FlowElement, 0)
						}
						newOut := append(outgoing, &flows[i])
						//设置上一个节点出口
						lastFlow.SetOutgoing(newOut)
						//设置当前连线入口
						flows[i].SetSourceFlowElement(&lastFlow)

					}
					//下一个节点
					TargetRef := value.TargetRef
					nextFlow := flowMap[TargetRef]
					if nextFlow != nil {
						incoming := nextFlow.GetIncoming()
						if incoming == nil {
							incoming = make([]*FlowElement, 0)
						}
						newIn := append(incoming, &flows[i])
						m := make([]*FlowElement, 1)
						m[0] = &nextFlow
						//设置当前连线出口
						flows[i].SetTargetFlowElement(&nextFlow)
						//设置写一个节点入口
						nextFlow.SetIncoming(newIn)
					}
				}
			}
		}
	}
}
