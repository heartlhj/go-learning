package ast

import (
	. "go-learning/expression/spel"
)

type BoolLiteral struct {
	*Literal
}

func (l BoolLiteral) GetValueInternal(expressionState ExpressionState) TypedValue {
	return l.Value
}
