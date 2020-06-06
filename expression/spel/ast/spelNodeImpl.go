package ast

import (
	. "go-learning/expression/spel"
)

type SpelNodeImpl struct {
	Children           []SpelNode
	Parent             SpelNode
	Pos                int
	exitTypeDescriptor string
}

func (o *SpelNodeImpl) GetValueInternal(expressionState ExpressionState) TypedValue {
	return o.GetValueInternal(expressionState)
}

func (o *SpelNodeImpl) GetValue(expressionState ExpressionState) interface{} {
	return o.GetValueInternal(expressionState)
}
