package model

import "github.com/heartlhj/go-learning/workflow/entity"

//流程实例
type HistoricProcess struct {
	*entity.ProcessInstanceEntity
	Id                int64
	ProcessInstanceId int64 `xorm:"proc_inst_id"`
}

func (HistoricProcess) TableName() string {
	return "hi_process_instance"
}
