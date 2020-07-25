package entity

type IdentityLinkEntity struct {
	Type              string `xorm:"type"`
	TaskId            int64  `xorm:"task_id"`
	ProcessInstanceId int    `xorm:"proc_inst_id"`
	GroupId           string `xorm:"group_id"`
	UserId            string `xorm:"user_id"`
}
