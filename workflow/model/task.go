package model

import "github.com/heartlhj/go-learning/workflow/entity"

type Task struct {
	*entity.TaskEntity
	Id int64
}

func (Task) TableName() string {
	return "task"
}
