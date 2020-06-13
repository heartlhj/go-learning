package model

//流程源数据
type Bytearry struct {
	Id           int    `json:"id" xorm:"pk autoincr"`
	Key          string `json:"key"`
	Name         string `json:"name"`
	Version      string `json:"version"`
	Bytes        string `json:"bytes"`
	DeploymentId int    `json:"deployment_id"`
}
