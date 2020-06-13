package model

//流程实例
type Procinst struct {
	Id          int    `json:"id" xorm:"pk autoincr"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	TenantId    string `json:"tenant_id"`
	Deploy      string `json:"deploy"`
	StartTime   string `json:"start_time"`
	StartUserId string `json:"start_user_id"`
}
