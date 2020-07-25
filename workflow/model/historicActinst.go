package model

import (
	"time"
)

//流程实例
type HistoricActinst struct {
	Id                int64
	ProcessDefineId   int64     `xorm:"process_define_id"`
	ProcessInstanceId int64     `xorm:"proc_inst_id"`
	TaskId            int64     `xorm:"task_id"`
	ActId             string    `xorm:"act_id"`
	ActName           string    `xorm:"act_name"`
	ActType           string    `xorm:"act_type"`
	TenantId          string    `xorm:"tenant_id"`
	StartTime         time.Time `xorm:"start_time"`
	EndTime           time.Time `xorm:"end_time"`
	StartUserId       string    `xorm:"start_user_id"`
	Assignee          string    `xorm:"assignee"`
}

func (HistoricActinst) TableName() string {
	return "hi_actinst"
}
