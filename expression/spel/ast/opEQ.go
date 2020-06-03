package ast

import (
	. "go-learning/expression"
	. "go-learning/expression/spel"
	"go-learning/expression/support"
)

type OpEQ struct {
	*Operator
}

func (o *OpEQ) getValueInternal(expressionState ExpressionState) TypedValue {
	left := o.getLeftOperand().getValueInternal(expressionState).Value
	right := o.getRightOperand().getValueInternal(expressionState).Value
	o.leftActualDescriptor = o.toDescriptorFromObject(left)
	o.rightActualDescriptor = o.toDescriptorFromObject(right)
	check := o.equalityCheck(left, right)
	value := support.BooleanTypedValue{}
	return value.ForValue(check)
}
