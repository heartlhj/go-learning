package ast

import (
	. "go-learning/expression/spel"
	. "go-learning/expression/support"
)

type OpAnd struct {
	*Operator
}

func (o *OpAnd) GetValueInternal(expressionState ExpressionState) TypedValue {
	if !getBooleanValue(expressionState, o.getLeftOperand()) {
		value := BooleanTypedValue{}
		return value.ForValue(false)
	}
	booleanValue := getBooleanValue(expressionState, o.getLeftOperand())
	value := BooleanTypedValue{}
	return value.ForValue(booleanValue)
}

func getBooleanValue(state ExpressionState, operand SpelNode) bool {
	value := operand.GetValue(state)
	if value == nil {
		panic("Type conversion problem, cannot convert from [null] to bool")
	}
	return value.(bool)
}
