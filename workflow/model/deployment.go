package model

//部署
type Deployment struct {
	Id         int    `json:"id" xorm:"pk autoincr"`
	Key        string `json:"key"`
	Name       string `json:"name"`
	Version    int    `json:"version"`
	TenantId   string `json:"tenant_id"`
	DeployTime string `json:"deploy_time"`
}
