package test

import (
	peocess "github.com/heartlhj/go-learning/workflow/engine/impl"
	"testing"
)

var key = "task001"

//测试发起流程
func TestStartProcss(t *testing.T) {
	runtime := peocess.RuntimeService{}

	runtime.StartProcessInstanceByKey(key, nil, "", "")
}

//测试完成任务
func TestComplete(t *testing.T) {
	taskService := peocess.TaskServiceImpl{}
	taskService.Complete(9, nil)
}
