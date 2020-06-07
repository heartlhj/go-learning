package ast

import (
	. "github.com/heartlhj/go-learning/expression/spel"
	. "github.com/heartlhj/go-learning/expression/support"
)

type OpOr struct {
	*Operator
}

func (o *OpOr) GetValueInternal(expressionState ExpressionState) TypedValue {
	if getBooleanValue(expressionState, o.getLeftOperand()) {
		value := BooleanTypedValue{}
		return value.ForValue(true)
	}
	booleanValue := getBooleanValue(expressionState, o.getLeftOperand())
	value := BooleanTypedValue{}
	return value.ForValue(booleanValue)
}
