package test

import (
	"fmt"
	"github.com/heartlhj/go-learning/workflow/engine/behavior"
	peocess "github.com/heartlhj/go-learning/workflow/engine/impl"
	"github.com/heartlhj/go-learning/workflow/event"
	"github.com/heartlhj/go-learning/workflow/runtime"
	"testing"
)

var key = "process_demo"

type ActivitiListener struct {
	name string
}

func (act ActivitiListener) OnEvent(event event.ActivitiEvent) error {
	fmt.Println(event)
	return nil
}

//测试发起流程
func TestStartProcss(t *testing.T) {
	variables := make(map[string]interface{}, 0)
	variables["name"] = "lisi"
	variables["age"] = 18
	variables["isOld"] = true
	runtime := peocess.RuntimeService{}
	runtime.StartProcessInstanceByKey(key, variables, "", "")
}

//测试完成任务
func TestComplete(t *testing.T) {
	taskService := peocess.TaskServiceImpl{}
	variables := make(map[string]interface{}, 0)
	variables["code"] = "0001"
	taskService.Complete(35, variables, true)
}

//测试完成任务
func TestRuntime(t *testing.T) {
	id := runtime.GoroutineId()
	fmt.Println(id)
}

//测试完成任务
func TestListener(t *testing.T) {
	configuration := behavior.GetProcessEngineConfiguration()
	eventListeners := make([]event.ActivitiEventListener, 0)
	eventListeners = append(eventListeners, ActivitiListener{})
	configuration.AddEventListeners(eventListeners)
	taskService := peocess.TaskServiceImpl{}
	taskService.Complete(7, nil, true)
}
