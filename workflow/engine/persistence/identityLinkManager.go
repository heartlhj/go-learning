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
func (identityLinkManager IdentityLinkManager) CreateIdentityLink() (err error) {
	err = db.DB().Create(&identityLinkManager.IdentityLink).Error
	if err != nil {
		log.Infoln("Create IdentityLink Err ", err)
		return err
	}
	err = identityLinkManager.createHistoricIdentityLink()
	return err
}

func (identityLinkManager IdentityLinkManager) SelectByProcessInstanceId(processInstanceId int64) ([]IdentityLink, error) {
	identityLink := make([]IdentityLink, 0)
	err := db.DB().Where("proc_inst_id = ?", processInstanceId).Find(&identityLink).Error
	if err != nil {
		log.Infoln("Select Variable err: ", err)
	}
	if identityLink == nil || len(identityLink) <= 0 {
		return []IdentityLink{}, errs.ProcessError{Code: "1001", Msg: "Not find"}
	}
	return identityLink, nil
}

func (identityLinkManager IdentityLinkManager) SelectByTaskId(taskId int64) ([]IdentityLink, error) {
	identityLink := make([]IdentityLink, 0)
	err := db.DB().Where("task_id = ?", taskId).Find(&identityLink).Error
	if err != nil {
		log.Infoln("Select Variable err: ", err)
	}
	if identityLink != nil && len(identityLink) > 0 {
		return identityLink, nil
	}
	return identityLink, errs.ProcessError{Code: "1001", Msg: "Not Find"}
}

func (identityLinkManager IdentityLinkManager) Delete(identityLinkId int64) {
	identityLink := IdentityLink{}
	err := db.DB().Where("id=?", identityLinkId).Delete(&identityLink).Error
	if err != nil {
		log.Infoln("delete Variable err: ", err)
	}
}

func (identityLinkManager IdentityLinkManager) createHistoricIdentityLink() (err error) {
	identityLink := identityLinkManager.IdentityLink
	historicIdentityLink := HistoricIdentityLink{}
	historicIdentityLink.UserId = identityLink.UserId
	historicIdentityLinkManager := HistoricIdentityLinkManager{}
	historicIdentityLinkManager.HistoricIdentityLink = historicIdentityLink
	err = historicIdentityLinkManager.Insert()
	return err
}
