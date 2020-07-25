package variable

import (
	"github.com/heartlhj/go-learning/workflow/entity"
)

type Variable struct {
	*entity.VariableEntity
	Id int64
}

func (Variable) TableName() string {
	return "variable"
}

func (variable Variable) GetName() string {
	return variable.Name
}

func (variable Variable) GetProcessInstanceId() int64 {
	return variable.ProcessInstanceId
}

func (variable Variable) GetTaskId() int64 {
	return variable.TaskId
}

func (variable Variable) GetNumberValue() int {
	return variable.Number
}

func (variable *Variable) SetNumberValue(value int) {
	variable.Number = value
}

func (variable Variable) GetTextValue() string {
	return variable.Text
}

func (variable *Variable) SetTextValue(value string) {
	variable.Text = value
}

func (variable *Variable) SetValue(value interface{}, variableType VariableType) {
	variableType.SetValue(value, variable)
}

func (variable *Variable) SetBlobValue(value string) {
	variable.Blob = value
}

func (variable Variable) GetBlobValue() string {
	return variable.Blob
}
