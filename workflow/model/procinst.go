package model

import (
	"github.com/heartlhj/go-learning/workflow/entity"
)

//流程实例
type ProcessInstance struct {
	*entity.ProcessInstanceEntity
	Id int64
}

func (ProcessInstance) TableName() string {
	return "process_instance"
}

func (processInstance ProcessInstance) setBusinessKey(businessKey string) {
	processInstance.BusinessKey = businessKey
}
