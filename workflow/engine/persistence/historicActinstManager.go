package persistence

import (
	"github.com/heartlhj/go-learning/workflow/db"
	"github.com/heartlhj/go-learning/workflow/engine"
	. "github.com/heartlhj/go-learning/workflow/model"
	"github.com/prometheus/common/log"
	"reflect"
	"time"
)

type HistoricActinstManager struct {
	HistoricActinst HistoricActinst
}

func (historicActinstManager HistoricActinstManager) Insert() {
	_, err := db.MasterDB.Insert(historicActinstManager.HistoricActinst)
	if err != nil {
		log.Infoln("Create HistoricActinst Err", err)
	}
}

func (historicActinstManager HistoricActinstManager) RecordActivityStart(entity engine.ExecutionEntity) {
	historicActinst := HistoricActinst{}
	historicActinst.ProcessDefineId = entity.GetProcessDefineId()
	historicActinst.ProcessInstanceId = entity.GetProcessInstanceId()
	historicActinst.ActId = entity.GetCurrentActivityId()
	if entity.GetCurrentFlowElement() != nil {
		historicActinst.ActName = entity.GetCurrentFlowElement().GetName()
		historicActinst.ActType = parseActivityType(entity.GetCurrentFlowElement())
	}
	historicActinst.StartTime = time.Now()
	historicActinstManager.HistoricActinst = historicActinst
	historicActinstManager.Insert()
}

func (historicActinstManager HistoricActinstManager) FindUnfinishedHistoricActivityInstancesByExecutionAndActivityId(processInstanceId int64, taskId int64) (HistoricActinst, error) {
	historicActinst := make([]HistoricActinst, 0)
	err := db.MasterDB.Where("task_id = ?", taskId).Where("proc_inst_id = ?", processInstanceId).Find(&historicActinst)
	if err != nil {
		log.Infoln("Select Variable err: ", err)
	}
	if historicActinst == nil || len(historicActinst) <= 0 {
		return historicActinst[0], nil
	}
	return HistoricActinst{}, nil
}

func (historicActinstManager HistoricActinstManager) Update() {
	historicActinst := make([]HistoricActinst, 0)
	_, err := db.MasterDB.Where("task_id = ?", historicActinstManager.HistoricActinst.TaskId).
		Where("proc_inst_id = ?", historicActinstManager.HistoricActinst.ProcessInstanceId).
		Where("end_time = ?", nil).
		Update(&historicActinst)
	if err != nil {
		log.Infoln("Update HistoricActinst err: ", err)
	}
}

func parseActivityType(element engine.FlowElement) string {
	typeOf := reflect.TypeOf(element)
	return typeOf.Name()
}
