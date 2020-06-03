package ast

import (
	. "go-learning/expression"
	. "go-learning/expression/spel"
)

type SpelNodeImpl struct {
	Children           []SpelNode
	Parent             SpelNode
	Pos                int
	exitTypeDescriptor string
}
type SpelNodeValue interface {
	getValueInternal(expressionState ExpressionState) TypedValue
}

func (o *SpelNodeImpl) getValue(expressionState ExpressionState) interface{} {
	return o.getValue(expressionState)
}
