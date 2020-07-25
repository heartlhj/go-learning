package persistence

import (
	"github.com/heartlhj/go-learning/workflow/db"
	. "github.com/heartlhj/go-learning/workflow/model"
	"github.com/prometheus/common/log"
)

type ProcessInstanceManager struct {
	Instance *ProcessInstance
}

//创建流程实例
func (processInstanceManager *ProcessInstanceManager) CreateProcessInstance() {
	_, err := db.MasterDB.Insert(processInstanceManager.Instance)
	if err != nil {
		log.Infoln("create processInstance err", err)
	}

	processInstanceManager.createHistoricProcessInstance()
}

//查询流程实例
func (processInstanceManager *ProcessInstanceManager) GetProcessInstance(processInstanceId int64) ProcessInstance {
	instance := make([]ProcessInstance, 0)
	err := db.MasterDB.Id(processInstanceId).Find(&instance)
	if err != nil {
		log.Infoln("create processInstance err", err)
	}
	return instance[0]
}

//删除流程实例
func (processInstanceManager ProcessInstanceManager) DeleteProcessInstance(processInstanceId int64) {
	_, err := db.MasterDB.Id(processInstanceId).Delete(ProcessInstance{})
	if err != nil {
		log.Infoln("delete processInstance err ", err)
	}
}

func (processInstanceManager *ProcessInstanceManager) createHistoricProcessInstance() {
	processInstance := processInstanceManager.Instance
	historicProcess := HistoricProcess{}
	historicProcess.ProcessInstanceEntity = processInstance.ProcessInstanceEntity
	historicProcess.ProcessInstanceId = processInstance.Id
	historicProcessManager := HistoricProcessManager{}
	historicProcessManager.HistoricProcess = historicProcess
	historicProcessManager.Insert()
}
