package persistence

import (
	"github.com/heartlhj/go-learning/workflow/db"
	. "github.com/heartlhj/go-learning/workflow/model"
	"github.com/prometheus/common/log"
)

//创建流程实例
func CreateProcessInstance(instance *ProcessInstance) {

	_, err := db.MasterDB.Insert(instance)
	if err != nil {
		log.Infoln("新增数据异常", err)
	}
}
