package persistence

import (
	"github.com/heartlhj/go-learning/workflow/db"
	"github.com/heartlhj/go-learning/workflow/model"
	"github.com/prometheus/common/log"
)

type IdentityLinkManager struct {
	IdentityLink model.IdentityLink
}

//创建流程实例
func (identityLinkManager IdentityLinkManager) CreateIdentityLink() {
	_, err := db.MasterDB.Insert(identityLinkManager.IdentityLink)
	if err != nil {
		log.Infoln("新增数据异常", err)
	}
}
