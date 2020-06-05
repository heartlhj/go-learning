package ast

import (
	. "go-learning/expression/spel"
	"go-learning/expression/support"
)

//处理小于
type OpLT struct {
	*Operator
}

func (o *OpLT) GetValueInternal(expressionState ExpressionState) TypedValue {
	left := o.getLeftOperand().GetValueInternal(expressionState).Value
	right := o.getRightOperand().GetValueInternal(expressionState).Value
	o.leftActualDescriptor = o.toDescriptorFromObject(left)
	o.rightActualDescriptor = o.toDescriptorFromObject(right)

	leftV := left.(int)
	rightV := right.(int)
	check := leftV < rightV
	value := support.BooleanTypedValue{}
	return value.ForValue(check)
}
