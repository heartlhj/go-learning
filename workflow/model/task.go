package model

type Task struct {
	Id                int    `json:"id" xorm:"pk autoincr"`
	taskDefineKey     string `json:"task_define_key"`
	taskDefineName    string `json:"task_define_name"`
	Version           int    `json:"version"`
	TenantId          string `json:"tenant_id"`
	DeploymentId      int    `json:"deployment_id"`
	StartTime         string `json:"start_time"`
	Assignee          string `json:"assignee"`
	ProcessInstanceId int    `json:"proc_inst_id"`
}
