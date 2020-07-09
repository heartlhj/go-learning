package model

import "time"

type Variable struct {
	Id                int64
	TaskId            int64     `xorm:"task_id"`
	ProcessInstanceId int64     `xorm:"proc_inst_id"`
	Type              string    `xorm:"type"`
	Date              time.Time `xorm:"date"`
	Number            int64     `xorm:"number"`
	Blob              string    `xorm:"blob"`
}
