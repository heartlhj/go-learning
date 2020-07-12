package variable

type BooleanType struct {
}

func (boolType BooleanType) GetTypeName() string {
	return "bool"
}

func (boolType BooleanType) GetValue(valueFields ValueFields) interface{} {
	return valueFields.GetNumberValue() == 1
}

func (boolType BooleanType) SetValue(value interface{}, valueFields ValueFields) {
	b, ok := value.(bool)
	if ok {
		if b {
			valueFields.SetNumberValue(1)
		} else {
			valueFields.SetNumberValue(0)
		}
	} else {
		valueFields.SetNumberValue(0)
	}
}
