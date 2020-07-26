package persistence

import (
	"github.com/heartlhj/go-learning/workflow/db"
	"github.com/heartlhj/go-learning/workflow/engine/variable"
	"github.com/prometheus/common/log"
)

type HistoricVariableManager struct {
	HistoricVariable variable.HistoricVariable
}

func (historicVariableManager HistoricVariableManager) Insert() {
	_, err := db.MasterDB.Insert(historicVariableManager.HistoricVariable)
	if err != nil {
		log.Infoln("Create HistoricVariable Err ", err)
	}
}
