package support

import . "go-learning/expression"

type BooleanTypedValue struct {
	*TypedValue
}

func (b *BooleanTypedValue) ForValue(bool2 bool) TypedValue {
	boo := TypedValue{}
	if bool2 {
		boo.Value = true
		return boo
	}

	boo.Value = false
	return boo
}
