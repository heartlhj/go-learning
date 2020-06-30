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
	manager.Insert()
	handleAssignments(user.UserTask, task.Id)
}

//保存候选用户
func handleAssignments(user UserTask, taskId int) {
	users := user.CandidateUsers
	if len(users) >= 0 {
		for _, user := range users {
			link := IdentityLink{TaskId: taskId, UserId: user}
			manager := IdentityLinkManager{IdentityLink: link}
			manager.CreateIdentityLink()
		}
	}
}
