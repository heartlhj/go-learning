package variable

import (
	"github.com/heartlhj/go-learning/workflow/entity"
)

type HistoricVariable struct {
	*entity.VariableEntity
	Id int64
}

func (HistoricVariable) TableName() string {
	return "hi_variable"
}

func (variable HistoricVariable) GetName() string {
	return variable.Name
}

func (variable HistoricVariable) GetProcessInstanceId() int64 {
	return variable.ProcessInstanceId
}

func (variable HistoricVariable) GetTaskId() int64 {
	return variable.TaskId
}

func (variable HistoricVariable) GetNumberValue() int {
	return variable.Number
}

func (variable *HistoricVariable) SetNumberValue(value int) {
	variable.Number = value
}

func (variable HistoricVariable) GetTextValue() string {
	return variable.Text
}

func (variable *HistoricVariable) SetTextValue(value string) {
	variable.Text = value
}

func (variable *HistoricVariable) SetValue(value interface{}, variableType VariableType) {
	variableType.SetValue(value, variable)
}

func (variable *HistoricVariable) SetBlobValue(value string) {
	variable.Blob = value
}

func (variable HistoricVariable) GetBlobValue() string {
	return variable.Blob
}
