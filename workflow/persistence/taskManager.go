package persistence

import (
	"github.com/heartlhj/go-learning/workflow/db"
	. "github.com/heartlhj/go-learning/workflow/model"
	"github.com/prometheus/common/log"
)

type TaskManager struct {
	Task Task
}

func (taskManager TaskManager) Insert(execution ExecutionEntity) {
	_, err := db.MasterDB.Insert(taskManager.Task)
	if err != nil {
		log.Infoln("新增数据异常", err)
	}
}
