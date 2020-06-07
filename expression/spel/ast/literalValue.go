package ast

import . "github.com/heartlhj/go-learning/expression/spel"

type LiteralValue struct {
	*SpelNodeImpl
}

func (l *LiteralValue) GetLiteralValue() TypedValue {
	return TypedValue{}
}
