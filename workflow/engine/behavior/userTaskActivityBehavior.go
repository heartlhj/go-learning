package behavior

import (
	"github.com/heartlhj/go-learning/workflow/engine"
	. "github.com/heartlhj/go-learning/workflow/engine/persistence"
	"github.com/heartlhj/go-learning/workflow/event"
	. "github.com/heartlhj/go-learning/workflow/event/impl"
	. "github.com/heartlhj/go-learning/workflow/model"
	"time"
)

type UserTaskActivityBehavior struct {
	UserTask engine.UserTask
}

//普通用户节点处理
func (user UserTaskActivityBehavior) Execute(execution engine.ExecutionEntity) (err error) {
	task := Task{}
	task.ProcessInstanceId = execution.GetProcessInstanceId()
	task.Assignee = user.UserTask.Assignee
	task.StartTime = time.Now()
	task.TaskDefineKey = user.UserTask.Id
	task.TaskDefineName = user.UserTask.Name
	manager := TaskManager{Task: &task}
	err = manager.Insert(execution)
	if err != nil {
		return err
	}
	err = handleAssignments(user.UserTask, task.Id)

	// All properties set, now firing 'create' events
	if GetProcessEngineConfiguration().EventDispatcher.IsEnabled() {
		activitiEntityEvent, err := CreateEntityEvent(event.TASK_CREATED, task)
		if err != nil {
			return err
		}
		GetProcessEngineConfiguration().EventDispatcher.DispatchEvent(activitiEntityEvent)
	}
	return err
}

//保存候选用户
func handleAssignments(user engine.UserTask, taskId int64) (err error) {
	users := user.CandidateUsers
	if len(users) >= 0 {
		for _, user := range users {
			link := IdentityLink{}
			link.TaskId = taskId
			link.UserId = user
			identityLinkManager := GetIdentityLinkManager()
			identityLinkManager.IdentityLink = link
			err = identityLinkManager.CreateIdentityLink()
			if err != nil {
				return err
			}
		}
	}
	return err
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
