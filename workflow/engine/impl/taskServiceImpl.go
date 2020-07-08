package engine

import (
	"github.com/heartlhj/go-learning/workflow/engine/behavior"
	"github.com/heartlhj/go-learning/workflow/engine/impl/cmd"
	. "github.com/heartlhj/go-learning/workflow/model"
)

type TaskServiceImpl struct {
}

//流程审批完成
func (task TaskServiceImpl) Complete(taskId int, variables map[string]interface{}) Task {
	exe := behavior.GetServiceImpl().CommandExecutor.Exe(cmd.CompleteCmd{TaskId: taskId, Variables: variables})
	return exe.(Task)
}
