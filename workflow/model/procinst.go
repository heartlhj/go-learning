package model

import (
	"time"
)

//流程实例
type ProcessInstance struct {
	Id           int       `json:"id" xorm:"pk autoincr"`
	Key          string    `json:"key"`
	Name         string    `json:"name"`
	Version      int       `json:"version"`
	BusinessKey  string    `json:"business_key"`
	TenantId     string    `json:"tenant_id"`
	DeploymentId int       `json:"deployment_id"`
	StartTime    time.Time `json:"start_time"`
	StartUserId  string    `json:"start_user_id"`
}

func (processInstance ProcessInstance) setBusinessKey(businessKey string) {
	processInstance.BusinessKey = businessKey
}
