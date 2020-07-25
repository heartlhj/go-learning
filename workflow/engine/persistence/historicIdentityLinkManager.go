package persistence

import (
	"github.com/heartlhj/go-learning/workflow/db"
	"github.com/heartlhj/go-learning/workflow/model"
	"github.com/prometheus/common/log"
)

type HistoricIdentityLinkManager struct {
	HistoricIdentityLink model.HistoricIdentityLink
}

func (historicIdentityLink HistoricIdentityLinkManager) Insert() {
	_, err := db.MasterDB.Insert(historicIdentityLink.HistoricIdentityLink)
	if err != nil {
		log.Infoln("Create HistoricIdentityLink Err", err)
	}
}
