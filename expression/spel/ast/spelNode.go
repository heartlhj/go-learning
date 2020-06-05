package ast

import (
	. "go-learning/expression/spel"
)

type SpelNode interface {
	GetValue(expressionState ExpressionState) interface{}
	GetValueInternal(expressionState ExpressionState) TypedValue
	SetPos(pos int)
}
