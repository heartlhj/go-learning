package persistence

import (
	"github.com/heartlhj/go-learning/workflow/db"
	. "github.com/heartlhj/go-learning/workflow/model"
	"github.com/prometheus/common/log"
)

type HistoricProcessManager struct {
	HistoricProcess HistoricProcess
}

func (historicProcessManager HistoricProcessManager) Insert() (err error) {
	err = db.TXDB.Create(&historicProcessManager.HistoricProcess).Error
	if err != nil {
		log.Infoln("Create HistoricActinst Err", err)
	}
	return err
}

func (historicProcessManager HistoricProcessManager) MarkEnded() (err error) {
	historicProcess := historicProcessManager.HistoricProcess
	err = db.TXDB.Where("proc_inst_id=?", historicProcess.ProcessInstanceId).Update(&historicProcess).Error
	if err != nil {
		log.Infoln("delete HistoricProcess Err", err)
		return err
	}
	historicActinst := HistoricActinst{}
	historicActinst.EndTime = historicProcess.EndTime
	historicProcess.ProcessInstanceId = historicProcess.Id
	historicActinstManager := HistoricActinstManager{}
	historicActinstManager.HistoricActinst = historicActinst
	err = historicActinstManager.UpdateProcessInstanceId()
	return err
}
