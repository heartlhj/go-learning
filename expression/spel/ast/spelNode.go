package ast

import (
	. "github.com/heartlhj/go-learning/expression/spel"
)

//表达式对象
type SpelNode interface {
	GetValue(expressionState ExpressionState) interface{}

	GetValueInternal(expressionState ExpressionState) TypedValue
}
