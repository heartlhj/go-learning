package ast

import . "go-learning/expression/spel"

type LiteralValue struct {
	*SpelNodeImpl
}

func (l *LiteralValue) GetLiteralValue() TypedValue {
	return TypedValue{}
}
