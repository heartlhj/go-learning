package model

import (
	"time"
)

//流程实例
type HistoricProcess struct {
	Id                int64
	ProcessInstanceId int64     `xorm:"proc_inst_id"`
	Key               string    `xorm:"key"`
	Name              string    `xorm:"name"`
	BusinessKey       string    `xorm:"business_key"`
	TenantId          string    `xorm:"tenant_id"`
	DeploymentId      int64     `xorm:"deployment_id"`
	StartTime         time.Time `xorm:"start_time"`
	EndTime           time.Time `xorm:"end_time"`
	StartUserId       string    `xorm:"start_user_id"`
	ProcessDefineId   int64     `xorm:"process_define_id"`
}

func (HistoricProcess) TableName() string {
	return "hi_process_instance"
}
