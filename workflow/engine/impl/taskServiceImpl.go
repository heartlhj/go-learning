package engine

import . "github.com/heartlhj/go-learning/workflow/model"

type TaskServiceImpl struct {
}

//流程审批完成
func (task TaskServiceImpl) Complete(taskId int, variables map[string]interface{}) Task {
	return Task{}
}
