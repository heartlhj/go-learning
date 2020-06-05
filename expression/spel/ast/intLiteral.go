package ast

import (
	. "go-learning/expression/spel"
)

// int 类型
type IntLiteral struct {
	*Literal
}

func (l IntLiteral) GetValueInternal(expressionState ExpressionState) TypedValue {
	return l.Value
}
