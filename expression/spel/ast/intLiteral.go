package ast

import (
	"go-learning/expression/spel"
)

type IntLiteral struct {
	*Literal
	Value spel.TypedValue
}

func (l *IntLiteral) GetLiteralValue() spel.TypedValue {
	return l.Value
}
