package persistence

import (
	"github.com/heartlhj/go-learning/workflow/db"
	. "github.com/heartlhj/go-learning/workflow/model"
	"github.com/prometheus/common/log"
)

type TaskManager struct {
	Task Task
}

func (taskManager TaskManager) Insert() {
	_, err := db.MasterDB.Insert(taskManager.Task)
	if err != nil {
		log.Infoln("新增数据异常", err)
	}
}

func (taskManager TaskManager) FindById(taskId int) []Task {
	task := make([]Task, 0)
	err := db.MasterDB.Where("id=?", taskId).Find(&task)
	if err != nil {
		log.Infoln("新增数据异常", err)
	}
	return task

}

func (taskManager TaskManager) DeleteTask(taskId int) {
	task := make([]Task, 0)
	_, err := db.MasterDB.Where("id=?", taskId).Delete(&task)
	if err != nil {
		log.Infoln("新增数据异常", err)
	}
}
