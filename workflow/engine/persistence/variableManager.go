package persistence

import (
	"github.com/heartlhj/go-learning/workflow/db"
	. "github.com/heartlhj/go-learning/workflow/engine/variable"
	"github.com/prometheus/common/log"
)

type VariableManager struct {
	Variable Variable
}

func (define VariableManager) Create(name string, variableType VariableType, value interface{}) *Variable {
	variable := Variable{}
	variable.Version = 0
	variable.Name = name
	variable.Type = variableType.GetTypeName()
	variable.SetValue(value, variableType)
	return &variable

}

func (defineManager VariableManager) Insert(define *Variable) {
	_, err := db.MasterDB.Insert(define)
	if err != nil {
		log.Infoln("新增数据异常", err)
	}
}

func (defineManager VariableManager) SelectProcessInstanceId(name string, processInstanceId int64) Variable {
	variable := Variable{}
	err := db.MasterDB.Where("proc_inst_id = ?", processInstanceId).Where("name = ?", name).Limit(1, 0).Find(&variable)
	if err != nil {
		log.Infoln("新增数据异常", err)
	}
	return variable
}
