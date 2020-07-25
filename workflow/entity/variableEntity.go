package entity

import "time"

type VariableEntity struct {
	Version           int64     `xorm:"version"`
	TaskId            int64     `xorm:"task_id"`
	ProcessInstanceId int64     `xorm:"proc_inst_id"`
	Name              string    `xorm:"name"`
	Type              string    `xorm:"type"`
	Date              time.Time `xorm:"date"`
	Number            int       `xorm:"number"`
	Float             float64   `xorm:"float"`
	Text              string    `xorm:"text"`
	Blob              string    `xorm:"blob"`
}
