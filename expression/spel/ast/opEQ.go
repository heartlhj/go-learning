package ast

import (
	. "github.com/heartlhj/go-learning/expression/spel"
	"github.com/heartlhj/go-learning/expression/support"
)

type OpEQ struct {
	*Operator
}

func (o *OpEQ) GetValueInternal(expressionState ExpressionState) TypedValue {
	left := o.getLeftOperand().GetValueInternal(expressionState).Value
	right := o.getRightOperand().GetValueInternal(expressionState).Value
	o.leftActualDescriptor = o.toDescriptorFromObject(left)
	o.rightActualDescriptor = o.toDescriptorFromObject(right)
	check := o.equalityCheck(left, right)
	value := support.BooleanTypedValue{}
	return value.ForValue(check)
}
