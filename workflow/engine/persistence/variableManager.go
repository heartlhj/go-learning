package persistence

import (
	"github.com/heartlhj/go-learning/workflow/db"
	. "github.com/heartlhj/go-learning/workflow/engine/variable"
	"github.com/heartlhj/go-learning/workflow/entity"
	"github.com/heartlhj/go-learning/workflow/errs"
	"github.com/prometheus/common/log"
)

type VariableManager struct {
	Variable *Variable
}

func (define VariableManager) Create(name string, variableType VariableType, value interface{}) *Variable {
	variable := Variable{VariableEntity: &entity.VariableEntity{}}
	variable.Version = 0
	variable.Name = name
	variable.Type = variableType.GetTypeName()
	variable.SetValue(value, variableType)
	return &variable

}

func (defineManager VariableManager) Insert() {
	_, err := db.MasterDB.Insert(defineManager.Variable)
	if err != nil {
		log.Infoln("Create Variable Error", err)
	}
	defineManager.createHistoricVariable()
}

func (defineManager VariableManager) createHistoricVariable() {
	variable := defineManager.Variable
	historicVariable := HistoricVariable{}
	historicVariable.VariableEntity = variable.VariableEntity

	historicVariableManager := HistoricVariableManager{}
	historicVariableManager.HistoricVariable = historicVariable
	historicVariableManager.Insert()
}

func (defineManager VariableManager) SelectProcessInstanceId(name string, processInstanceId int64) (Variable, error) {
	variables := make([]*Variable, 0)
	err := db.MasterDB.Where("proc_inst_id = ?", processInstanceId).Where("name = ?", name).Limit(1, 0).Find(&variables)
	if err != nil {
		log.Infoln("Select Variable err: ", err)
	}
	if variables != nil || len(variables) >= 0 {
		return Variable{}, errs.ProcessError{}
	}
	return *variables[0], nil
}

func (variableManager VariableManager) SelectTaskId(name string, taskId int64) (Variable, error) {
	variables := make([]*Variable, 0)
	err := db.MasterDB.Where("task_id = ?", taskId).Where("name = ?", name).Limit(1, 0).Find(&variables)
	if err != nil {
		log.Infoln("根据[taskId] 查询流程变量异常", err)
	}
	if variables != nil || len(variables) >= 0 {
		return Variable{}, errs.ProcessError{}
	}
	return *variables[0], nil
}

func (variableManager VariableManager) SelectByProcessInstanceId(processInstanceId int64) ([]Variable, error) {
	variables := make([]Variable, 0)
	err := db.MasterDB.Where("proc_inst_id = ?", processInstanceId).Find(&variables)
	if err != nil {
		log.Infoln("Select Variable err: ", err)
	}
	if variables == nil || len(variables) <= 0 {
		return []Variable{}, errs.ProcessError{}
	}
	return variables, nil
}

func (variableManager VariableManager) SelectByTaskId(taskId int64) ([]Variable, error) {
	variables := make([]Variable, 0)
	err := db.MasterDB.Where("task_id = ?", taskId).Find(&variables)
	if err != nil {
		log.Infoln("Select Variable err: ", err)
		return []Variable{}, nil
	}
	if variables == nil || len(variables) <= 0 {
		return []Variable{}, errs.ProcessError{}
	}
	return variables, nil
}

func (variableManager VariableManager) Delete(variableId int64) {
	task := Variable{}
	_, err := db.MasterDB.Id(variableId).Delete(task)
	if err != nil {
		log.Infoln("delete Variable err: ", err)
	}
}
