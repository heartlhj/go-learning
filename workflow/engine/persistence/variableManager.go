package persistence

import (
	"github.com/heartlhj/go-learning/workflow/db"
	. "github.com/heartlhj/go-learning/workflow/engine/variable"
	"github.com/heartlhj/go-learning/workflow/errs"
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

func (defineManager VariableManager) SelectProcessInstanceId(name string, processInstanceId int64) (Variable, error) {
	variables := make([]*Variable, 0)
	err := db.MasterDB.Where("proc_inst_id = ?", processInstanceId).Where("name = ?", name).Limit(1, 0).Find(&variables)
	if err != nil {
		log.Infoln("新增数据异常", err)
	}
	if variables == nil || len(variables) <= 0 {
		return Variable{}, errs.ProcessError{}
	}
	return *variables[0], nil
}

func (defineManager VariableManager) SelectTakId(name string, taskId int64) (Variable, error) {
	variables := make([]*Variable, 0)
	err := db.MasterDB.Where("task_id = ?", taskId).Where("name = ?", name).Limit(1, 0).Find(&variables)
	if err != nil {
		log.Infoln("新增数据异常", err)
	}
	if variables == nil || len(variables) <= 0 {
		return Variable{}, errs.ProcessError{}
	}
	return *variables[0], nil
}

func (defineManager VariableManager) SelectByProcessInstanceId(processInstanceId int64) ([]Variable, error) {
	variables := make([]Variable, 0)
	err := db.MasterDB.Where("proc_inst_id = ?", processInstanceId).Find(&variables)
	if err != nil {
		log.Infoln("新增数据异常", err)
	}
	if variables == nil || len(variables) <= 0 {
		return []Variable{}, errs.ProcessError{}
	}
	return variables, nil
}

func (defineManager VariableManager) SelectByTaskId(taskId int64) ([]Variable, error) {
	variables := make([]Variable, 0)
	err := db.MasterDB.Where("task_id = ?", taskId).Find(&variables)
	if err != nil {
		log.Infoln("新增数据异常", err)
	}
	if variables == nil || len(variables) <= 0 {
		return []Variable{}, errs.ProcessError{}
	}
	return variables, nil
}
