package ast

import (
	. "go-learning/expression/spel"
)

// Float64 类型
type FloatLiteral struct {
	*Literal
}

func (l FloatLiteral) GetValueInternal(expressionState ExpressionState) TypedValue {
	return l.Value
}
