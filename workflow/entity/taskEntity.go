package entity

import "time"

type TaskEntity struct {
	TaskDefineKey     string    `xorm:"task_define_key"`
	TaskDefineName    string    `xorm:"task_define_name"`
	Version           int       `xorm:"version"`
	TenantId          string    `xorm:"tenant_id"`
	DeploymentId      int       `xorm:"deployment_id"`
	StartTime         time.Time `xorm:"start_time"`
	Assignee          string    `xorm:"assignee"`
	ProcessInstanceId int64     `xorm:"proc_inst_id"`
}
