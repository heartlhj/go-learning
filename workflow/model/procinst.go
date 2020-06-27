package model

//流程实例
type ProcessInstance struct {
	Id          int    `json:"id" xorm:"pk autoincr"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	BusinessKey string `json:"businessKey"`
	TenantId    string `json:"tenant_id"`
	Deploy      string `json:"deploy"`
	StartTime   string `json:"start_time"`
	StartUserId string `json:"start_user_id"`
}

func (processInstance ProcessInstance) setBusinessKey(businessKey string) {
	processInstance.BusinessKey = businessKey
}
