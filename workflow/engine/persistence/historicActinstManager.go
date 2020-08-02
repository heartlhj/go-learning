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
	err := db.TXDB.Create(&historicActinstManager.HistoricActinst).Error
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
		historicActinst.ActType = historicActinstManager.parseActivityType(entity.GetCurrentFlowElement())
	}
	historicActinst.StartTime = time.Now()
	historicActinstManager.HistoricActinst = historicActinst
	historicActinstManager.Insert()
}

func (historicActinstManager HistoricActinstManager) FindUnfinishedHistoricActivityInstancesByExecutionAndActivityId(processInstanceId int64, actId string) (HistoricActinst, error) {
	historicActinst := HistoricActinst{}
	err := db.TXDB.Where("act_id = ?", actId).Where("proc_inst_id = ?", processInstanceId).First(&historicActinst).Error
	if err != nil {
		log.Infoln("Select HistoricActinst err: ", err)
		return HistoricActinst{}, err
	}
	return historicActinst, nil
}

func (historicActinstManager HistoricActinstManager) Update() (err error) {
	err = db.TXDB.Model(&HistoricActinst{}).Where("act_id = ?", historicActinstManager.HistoricActinst.ActId).
		Where("proc_inst_id = ?", historicActinstManager.HistoricActinst.ProcessInstanceId).
		Where("end_time IS NULL").
		Update(&historicActinstManager.HistoricActinst).Error
	if err != nil {
		log.Infoln("Update HistoricActinst err: ", err)
	}
	return err
}
func (historicActinstManager HistoricActinstManager) UpdateProcessInstanceId() (err error) {
	err = db.TXDB.Model(&HistoricActinst{}).Where("proc_inst_id = ?", historicActinstManager.HistoricActinst.ProcessInstanceId).
		Update(&historicActinstManager.HistoricActinst).Error
	if err != nil {
		log.Infoln("Update HistoricActinst err: ", err)
	}
	return err
}

func (historicActinstManager HistoricActinstManager) UpdateTaskId() (err error) {
	err = db.TXDB.Model(&HistoricActinst{}).Where("task_id = ?", historicActinstManager.HistoricActinst.TaskId).
		Update(&historicActinstManager.HistoricActinst).Error
	if err != nil {
		log.Infoln("Update HistoricActinst err: ", err)
	}
	return err
}

func (historicActinstManager HistoricActinstManager) RecordTaskCreated(element engine.FlowElement, entity engine.ExecutionEntity) (err error) {
	var actinst = HistoricActinst{}
	actinst, err = historicActinstManager.FindUnfinishedHistoricActivityInstancesByExecutionAndActivityId(entity.GetProcessInstanceId(), element.GetId())
	if err == nil {
		actinst.EndTime = time.Now()
		historicActinstManager.HistoricActinst = actinst
		err = historicActinstManager.Update()
	}
	return err
}

func (historicActinstManager HistoricActinstManager) parseActivityType(element engine.FlowElement) string {
	typeOf := reflect.TypeOf(element)
	return typeOf.Name()
}
