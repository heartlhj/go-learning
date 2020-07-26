package persistence

import (
	"github.com/heartlhj/go-learning/workflow/db"
	. "github.com/heartlhj/go-learning/workflow/model"
	"github.com/prometheus/common/log"
)

type HistoricTaskManager struct {
	HistoricTask HistoricTask
}

func (historicTaskManager HistoricTaskManager) Insert() {
	_, err := db.MasterDB.Insert(historicTaskManager.HistoricTask)
	if err != nil {
		log.Infoln("Create HistoricTask Err", err)
	}
}

func (historicTaskManager HistoricTaskManager) MarkEnded() {
	_, err := db.MasterDB.Where("task_id=?", historicTaskManager.HistoricTask.TaskId).Update(historicTaskManager.HistoricTask)
	if err != nil {
		log.Infoln("Create HistoricTask Err", err)
	}
}
