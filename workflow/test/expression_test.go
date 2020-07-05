package test

import (
	peocess "github.com/heartlhj/go-learning/workflow/engine/impl"
	"testing"
)

var key = "task001"

//测试数组
func TestStartProcss(t *testing.T) {
	runtime := peocess.RuntimeService{}

	runtime.StartProcessInstanceByKey(key, nil, "", "")
}
