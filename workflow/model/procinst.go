package model

import (
	"time"
)

//流程实例
type ProcessInstance struct {
	Id              int64
	Key             string    `xorm:"key"`
	Name            string    `xorm:"name"`
	BusinessKey     string    `xorm:"business_key"`
	TenantId        string    `xorm:"tenant_id"`
	DeploymentId    int64     `xorm:"deployment_id"`
	StartTime       time.Time `xorm:"start_time"`
	StartUserId     string    `xorm:"start_user_id"`
	ProcessDefineId int64     `xorm:"process_define_id"`
}

func (ProcessInstance) TableName() string {
	return "process_instance"
}

func (processInstance ProcessInstance) setBusinessKey(businessKey string) {
	processInstance.BusinessKey = businessKey
}
