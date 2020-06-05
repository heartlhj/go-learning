package ast

import (
	. "go-learning/expression/spel"
)

type LongLiteral struct {
	*Literal
}

func (l LongLiteral) GetValueInternal(expressionState ExpressionState) TypedValue {
	return l.Value
}
