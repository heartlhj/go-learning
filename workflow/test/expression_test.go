package test

import (
	peocess "github.com/heartlhj/go-learning/workflow/engine/impl"
	"testing"
)

var key = "task001"

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
	taskService.Complete(13, nil)
}
