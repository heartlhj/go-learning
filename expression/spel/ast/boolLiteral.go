package ast

import (
	"go-learning/expression/spel"
)

type BoolLiteral struct {
	*Literal
	value spel.TypedValue
}

func (l *BoolLiteral) GetLiteralValue() spel.TypedValue {
	return l.value
}
