package persistence

import (
	"github.com/heartlhj/go-learning/workflow/db"
	"github.com/heartlhj/go-learning/workflow/model"
	"github.com/prometheus/common/log"
)

type VariableManager struct {
	Variable model.Variable
}

func (define VariableManager) Insert(key string) {
	_, err := db.MasterDB.Insert(define.Variable)
	if err != nil {
		log.Infoln("新增数据异常", err)
	}
}
