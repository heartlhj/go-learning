package ast

import (
	. "go-learning/expression/spel"
)

type GetLiteralValue interface {
	GetLiteralValue() TypedValue
}

type Literal struct {
	*SpelNodeImpl
	OriginalValue string
}

func (l *Literal) GetValueInternal(expressionState ExpressionState) TypedValue {
	return l.GetLiteralValue()
}

func (l *Literal) GetLiteralValue() TypedValue {
	return l.GetLiteralValue()
}
