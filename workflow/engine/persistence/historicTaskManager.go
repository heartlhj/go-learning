package persistence

import (
	"github.com/heartlhj/go-learning/workflow/db"
	. "github.com/heartlhj/go-learning/workflow/model"
	"github.com/prometheus/common/log"
)

type HistoricTaskManager struct {
	HistoricTask HistoricTask
}

func (historicTaskManager HistoricTaskManager) Insert() (err error) {
	err = db.TXDB.Create(&historicTaskManager.HistoricTask).Error
	if err != nil {
		log.Infoln("Create HistoricTask Err", err)
	}
	return err
}

func (historicTaskManager HistoricTaskManager) MarkEnded() {
	err := db.TXDB.Model(&HistoricTask{}).Where("task_id=?", historicTaskManager.HistoricTask.TaskId).Update(&historicTaskManager.HistoricTask).Error
	if err != nil {
		log.Infoln("Create HistoricTask Err", err)
	}
}
