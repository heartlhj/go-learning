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

func (historicProcessManager HistoricProcessManager) MarkEnded() {
	historicProcess := historicProcessManager.HistoricProcess
	_, err := db.MasterDB.Where("proc_inst_id=?", historicProcess.ProcessInstanceId).Update(historicProcess)
	if err != nil {
		log.Infoln("delete HistoricProcess Err", err)
	}
	historicActinst := HistoricActinst{}
	historicActinst.EndTime = historicProcess.EndTime
	historicProcess.ProcessInstanceId = historicProcess.Id
	historicActinstManager := HistoricActinstManager{}
	historicActinstManager.HistoricActinst = historicActinst
	historicActinstManager.UpdateProcessInstanceId()
}
