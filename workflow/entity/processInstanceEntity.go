package entity

import "time"

type ProcessInstanceEntity struct {
	Key             string    `xorm:"key"`
	Name            string    `xorm:"name"`
	Version         int       `xorm:"version"`
	BusinessKey     string    `xorm:"business_key"`
	TenantId        string    `xorm:"tenant_id"`
	DeploymentId    int64     `xorm:"deployment_id"`
	StartTime       time.Time `xorm:"start_time"`
	StartUserId     string    `xorm:"start_user_id"`
	ProcessDefineId int64     `xorm:"process_define_id"`
}
