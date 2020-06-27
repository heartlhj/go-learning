package behavior

import (
	. "github.com/heartlhj/go-learning/workflow/model"
	. "github.com/heartlhj/go-learning/workflow/persistence"
)

type UserTaskActivityBehavior struct {
	UserTask UserTask
}

//普通用户节点处理
func (user UserTaskActivityBehavior) Execute(execution ExecutionEntity) {

	task := Task{Assignee: user.UserTask.Assignee}
	manager := TaskManager{Task: task}
	manager.Insert(execution)
}
