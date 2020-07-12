package variable

type ValueFields interface {
	GetName() string

	GetProcessInstanceId() int64

	GetTaskId() int64

	GetNumberValue() int64

	SetNumberValue(value int64)

	GetTextValue() string

	SetTextValue(value string)
}
