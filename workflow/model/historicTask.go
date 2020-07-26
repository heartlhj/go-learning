package model

import (
	"time"
)

type HistoricTask struct {
	Id                int64
	TaskId            int64     `xorm:"task_id"`
	TaskDefineKey     string    `xorm:"task_define_key"`
	TaskDefineName    string    `xorm:"task_define_name"`
	TenantId          string    `xorm:"tenant_id"`
	DeploymentId      int       `xorm:"deployment_id"`
	StartTime         time.Time `xorm:"start_time"`
	EndTime           time.Time `xorm:"end_time"`
	Assignee          string    `xorm:"assignee"`
	ProcessInstanceId int64     `xorm:"proc_inst_id"`
}

func (HistoricTask) TableName() string {
	return "hi_task"
}
