package variable

import "time"

type Variable struct {
	Id                int64
	Version           int64     `xorm:"version"`
	TaskId            int64     `xorm:"task_id"`
	ProcessInstanceId int64     `xorm:"proc_inst_id"`
	Name              string    `xorm:"name"`
	Type              string    `xorm:"type"`
	Date              time.Time `xorm:"date"`
	Number            int64     `xorm:"number"`
	Float             float64   `xorm:"float"`
	Text              string    `xorm:"text"`
	Blob              string    `xorm:"blob"`
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

func (variable Variable) GetNumberValue() int64 {
	return variable.Number
}

func (variable *Variable) SetNumberValue(value int64) {
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
