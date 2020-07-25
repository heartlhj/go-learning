package persistence

import (
	"github.com/heartlhj/go-learning/workflow/db"
	"github.com/heartlhj/go-learning/workflow/errs"
	. "github.com/heartlhj/go-learning/workflow/model"
	"github.com/prometheus/common/log"
)

type IdentityLinkManager struct {
	IdentityLink IdentityLink
}

//创建流程实例
func (identityLinkManager IdentityLinkManager) CreateIdentityLink() {
	_, err := db.MasterDB.Insert(identityLinkManager.IdentityLink)
	if err != nil {
		log.Infoln("Create IdentityLink Err ", err)
	}
	identityLinkManager.createHistoricIdentityLink()
}

func (identityLinkManager IdentityLinkManager) SelectByProcessInstanceId(processInstanceId int64) ([]IdentityLink, error) {
	variables := make([]IdentityLink, 0)
	err := db.MasterDB.Where("proc_inst_id = ?", processInstanceId).Find(&variables)
	if err != nil {
		log.Infoln("Select Variable err: ", err)
	}
	if variables == nil || len(variables) <= 0 {
		return []IdentityLink{}, errs.ProcessError{}
	}
	return variables, nil
}

func (identityLinkManager IdentityLinkManager) SelectByTaskId(taskId int64) ([]IdentityLink, error) {
	variables := make([]IdentityLink, 0)
	err := db.MasterDB.Where("task_id = ?", taskId).Find(&variables)
	if err != nil {
		log.Infoln("Select Variable err: ", err)
	}
	if variables == nil || len(variables) <= 0 {
		return []IdentityLink{}, errs.ProcessError{}
	}
	return variables, nil
}

func (identityLinkManager IdentityLinkManager) Delete(identityLinkId int64) {
	task := IdentityLink{}
	_, err := db.MasterDB.Id(identityLinkId).Delete(task)
	if err != nil {
		log.Infoln("delete Variable err: ", err)
	}
}

func (identityLinkManager IdentityLinkManager) createHistoricIdentityLink() {
	identityLink := identityLinkManager.IdentityLink
	historicIdentityLink := HistoricIdentityLink{}
	historicIdentityLink.IdentityLinkEntity = identityLink.IdentityLinkEntity

	historicIdentityLinkManager := HistoricIdentityLinkManager{}
	historicIdentityLinkManager.HistoricIdentityLink = historicIdentityLink
	historicIdentityLinkManager.Insert()
}
