package persistence

import (
	"github.com/heartlhj/go-learning/workflow/db"
	. "github.com/heartlhj/go-learning/workflow/model"
	"github.com/prometheus/common/log"
)

type ProcessInstanceManager struct {
	Instance *ProcessInstance
}

//创建流程实例
func (processInstanceManager ProcessInstanceManager) CreateProcessInstance() {
	_, err := db.MasterDB.Insert(processInstanceManager.Instance)
	if err != nil {
		log.Infoln("新增数据异常", err)
	}
}