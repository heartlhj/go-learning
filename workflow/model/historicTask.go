package model

import "github.com/heartlhj/go-learning/workflow/entity"

type HistoricTask struct {
	*entity.TaskEntity
	Id     int64
	TaskId int64 `xorm:"task_id"`
}

func (HistoricTask) TableName() string {
	return "hi_task"
}
