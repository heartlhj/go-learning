package model

type Task struct {
	Id                int    `json:"id" xorm:"pk autoincr"`
	taskDefineKey     string `json:"task_define_key"`
	taskDefineName    string `json:"task_define_name"`
	Version           string `json:"version"`
	TenantId          string `json:"tenant_id"`
	Deploy            string `json:"deploy"`
	StartTime         string `json:"start_time"`
	Assignee          string `json:"assignee"`
	ProcessInstanceId string `json:"proc_inst_id"`
}
