package model

import "time"

//部署
type Deployment struct {
	Id         int       `xorm:"id" xorm:"pk autoincr"`
	Key        string    `xorm:"key"`
	Name       string    `xorm:"name"`
	Version    int       `xorm:"version"`
	TenantId   string    `xorm:"tenant_id"`
	DeployTime time.Time `xorm:"deploy_time"`
}
