package model

type IdentityLink struct {
	Id                int    `xorm:"id" xorm:"pk autoincr"`
	Type              string `xorm:"type"`
	TaskId            int    `xorm:"task_id"`
	ProcessInstanceId int    `xorm:"proc_inst_id"`
	GroupId           string `xorm:"group_id"`
	UserId            string `xorm:"user_id"`
}
