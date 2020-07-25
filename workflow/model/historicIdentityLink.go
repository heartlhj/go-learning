package model

import "github.com/heartlhj/go-learning/workflow/entity"

type HistoricIdentityLink struct {
	*entity.IdentityLinkEntity
	Id int64
}

func (HistoricIdentityLink) TableName() string {
	return "hi_identity_link"
}
