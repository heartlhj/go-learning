package persistence

import (
	"github.com/heartlhj/go-learning/workflow/db"
	. "github.com/heartlhj/go-learning/workflow/engine/variable"
	"github.com/heartlhj/go-learning/workflow/errs"
	"github.com/prometheus/common/log"
)

type VariableManager struct {
	Variable *Variable
}

func (define VariableManager) Create(name string, variableType VariableType, value interface{}) *Variable {
	variable := Variable{}
	variable.Version = 0
	variable.Name = name
	variable.Type = variableType.GetTypeName()
	variable.SetValue(value, variableType)
	return &variable

}

func (defineManager VariableManager) Insert() (err error) {
	err = db.DB().Create(&defineManager.Variable).Error
	if err != nil {
		log.Infoln("Create Variable Error", err)
		return err
	}
	err = defineManager.createHistoricVariable()
	return err
}

func (defineManager VariableManager) createHistoricVariable() (err error) {
	variable := defineManager.Variable
	historicVariable := HistoricVariable{}

	historicVariable.TaskId = variable.TaskId
	historicVariable.ProcessInstanceId = variable.ProcessInstanceId
	historicVariable.Name = variable.Name
	historicVariable.Version = variable.Version
	historicVariable.Type = variable.Type
	historicVariable.Text = variable.Text
	historicVariable.Number = variable.Number
	historicVariable.Date = variable.Date
	historicVariable.Float = variable.Float
	historicVariable.Blob = variable.Blob

	historicVariableManager := HistoricVariableManager{}
	historicVariableManager.HistoricVariable = historicVariable
	return historicVariableManager.Insert()
}

func (defineManager VariableManager) SelectProcessInstanceId(name string, processInstanceId int64) (Variable, error) {
	variables := Variable{}
	err := db.DB().Where("proc_inst_id = ?", processInstanceId).Where("name = ?", name).First(&variables).Error
	if err != nil {
		log.Infoln("Select Variable err: ", err)
		return Variable{}, err
	}
	return variables, nil
}

func (variableManager VariableManager) SelectTaskId(name string, taskId int64) (Variable, error) {
	variables := Variable{}
	err := db.DB().Where("task_id = ?", taskId).Where("name = ?", name).First(&variables).Error
	if err != nil {
		log.Infoln("根据[taskId] 查询流程变量异常", err)
		return Variable{}, err
	}
	return variables, nil
}

func (variableManager VariableManager) SelectByProcessInstanceId(processInstanceId int64) ([]Variable, error) {
	variables := make([]Variable, 0)
	err := db.DB().Where("proc_inst_id = ?", processInstanceId).Find(&variables).Error
	if err != nil {
		log.Infoln("Select Variable err: ", err)
		return variables, err
	}
	if variables != nil && len(variables) > 0 {
		return variables, nil
	}
	return variables, errs.ProcessError{Code: "1001", Msg: "Not Find"}
}

func (variableManager VariableManager) SelectByTaskId(taskId int64) ([]Variable, error) {
	variables := make([]Variable, 0)
	err := db.DB().Where("task_id = ?", taskId).Find(&variables).Error
	if err != nil {
		log.Infoln("Select Variable err: ", err)
		return variables, err
	}
	if variables != nil && len(variables) > 0 {
		return variables, nil
	}
	return variables, errs.ProcessError{Code: "1001", Msg: "Not Find"}
}

func (variableManager VariableManager) Delete(variableId int64) {
	variable := Variable{}
	err := db.DB().Where("id=?", variableId).Delete(variable).Error
	if err != nil {
		log.Infoln("delete Variable err: ", err)
	}
}
