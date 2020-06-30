package model

type IdentityLink struct {
	Id                int    `json:"id" xorm:"pk autoincr"`
	Type              string `json:"type"`
	TaskId            int    `json:"task_id"`
	ProcessInstanceId int    `json:"processInstance_id"`
	GroupId           string `json:"group_id"`
	UserId            string `json:"user_id"`
}
