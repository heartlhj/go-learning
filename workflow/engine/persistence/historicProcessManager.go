package persistence

import (
	"github.com/heartlhj/go-learning/workflow/db"
	. "github.com/heartlhj/go-learning/workflow/model"
	"github.com/prometheus/common/log"
)

type HistoricProcessManager struct {
	HistoricProcess HistoricProcess
}

func (historicProcessManager HistoricProcessManager) Insert() {
	_, err := db.MasterDB.Insert(historicProcessManager.HistoricProcess)
	if err != nil {
		log.Infoln("Create HistoricActinst Err", err)
	}
}
