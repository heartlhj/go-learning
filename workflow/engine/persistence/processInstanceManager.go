package persistence

import (
	"github.com/heartlhj/go-learning/workflow/db"
	. "github.com/heartlhj/go-learning/workflow/model"
	"github.com/prometheus/common/log"
	"time"
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
	processInstanceManager.recordActivityEnd(processInstanceId)
}

func (processInstanceManager ProcessInstanceManager) recordActivityEnd(processInstanceId int64) {
	historicProcessManager := HistoricProcessManager{}
	historicProcess := HistoricProcess{}
	historicProcess.ProcessInstanceId = processInstanceId
	historicProcess.EndTime = time.Now()
	historicProcessManager.HistoricProcess = historicProcess
	historicProcessManager.MarkEnded()
}

func (processInstanceManager *ProcessInstanceManager) createHistoricProcessInstance() {
	processInstance := processInstanceManager.Instance
	historicProcess := HistoricProcess{}
	//historicProcess.ProcessInstanceEntity = processInstance.ProcessInstanceEntity
	historicProcess.ProcessInstanceId = processInstance.Id
	historicProcess.DeploymentId = processInstance.DeploymentId
	historicProcess.TenantId = processInstance.TenantId
	historicProcess.StartTime = processInstance.StartTime
	historicProcess.Name = processInstance.Name
	historicProcess.BusinessKey = processInstance.BusinessKey
	historicProcess.StartUserId = processInstance.StartUserId
	historicProcess.Key = processInstance.Key

	historicProcessManager := HistoricProcessManager{}
	historicProcessManager.HistoricProcess = historicProcess
	historicProcessManager.Insert()
}
