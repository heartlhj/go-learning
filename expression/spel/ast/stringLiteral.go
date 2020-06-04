package ast

import (
	"go-learning/expression/spel"
)

type StringLiteral struct {
	*Literal
	Value spel.TypedValue
}

func (l *StringLiteral) GetLiteralValue() spel.TypedValue {
	return l.Value
}
