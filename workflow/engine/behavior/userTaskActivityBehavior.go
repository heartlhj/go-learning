package behavior

import (
	"github.com/heartlhj/go-learning/workflow/engine"
	. "github.com/heartlhj/go-learning/workflow/engine/persistence"
	. "github.com/heartlhj/go-learning/workflow/model"
	"time"
)

type UserTaskActivityBehavior struct {
	UserTask engine.UserTask
}

//普通用户节点处理
func (user UserTaskActivityBehavior) Execute(execution engine.ExecutionEntity) {

	task := Task{ProcessInstanceId: execution.GetProcessInstanceId(), Assignee: user.UserTask.Assignee, StartTime: time.Now()}
	task.TaskDefineKey = user.UserTask.Id
	task.TaskDefineName = user.UserTask.Name
	manager := TaskManager{Task: &task}
	manager.Insert()
	handleAssignments(user.UserTask, task.Id)
}

//保存候选用户
func handleAssignments(user engine.UserTask, taskId int64) {
	users := user.CandidateUsers
	if len(users) >= 0 {
		for _, user := range users {
			link := IdentityLink{TaskId: taskId, UserId: user}
			manager := IdentityLinkManager{IdentityLink: link}
			manager.CreateIdentityLink()
		}
	}
}

//普通用户节点处理
func (user UserTaskActivityBehavior) Trigger(execution engine.ExecutionEntity) {
	user.Leave(execution)
}

func (user UserTaskActivityBehavior) Leave(execution engine.ExecutionEntity) {
	element := execution.GetCurrentFlowElement()
	execution.SetCurrentFlowElement(element)
	GetAgenda().PlanTakeOutgoingSequenceFlowsOperation(execution, true)
}
