package ast

import (
	. "go-learning/expression/spel"
	"go-learning/expression/support"
)

//处理大于等于
type OpGE struct {
	*Operator
}

func (o *OpGE) GetValueInternal(expressionState ExpressionState) TypedValue {
	left := o.getLeftOperand().GetValueInternal(expressionState).Value
	right := o.getRightOperand().GetValueInternal(expressionState).Value
	o.leftActualDescriptor = o.toDescriptorFromObject(left)
	o.rightActualDescriptor = o.toDescriptorFromObject(right)
	leftV := left.(int)
	rightV := right.(int)
	check := leftV >= rightV
	value := support.BooleanTypedValue{}
	return value.ForValue(check)
}
