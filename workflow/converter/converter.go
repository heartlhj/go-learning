package converter

import "go-learning/workflow/db"

func convertXMLToElement(model *db.Definitions) *db.Definitions {
	return nil

}

func converter(d *db.Definitions) {
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
