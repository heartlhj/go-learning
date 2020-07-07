package model

//流程源数据
type Bytearry struct {
	Id           int    `xorm:"id" xorm:"pk autoincr"`
	Key          string `xorm:"key"`
	Name         string `xorm:"name"`
	Version      int    `xorm:"version"`
	Bytes        string `xorm:"bytes"`
	DeploymentId int    `xorm:"deployment_id"`
}
