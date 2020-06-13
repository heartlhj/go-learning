package engine

import . "github.com/heartlhj/go-learning/workflow/model"

type TaskService interface {
	Complete(taskId int, variables map[string]interface{}) Task
}
