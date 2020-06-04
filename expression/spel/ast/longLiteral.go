package ast

import (
	"go-learning/expression/spel"
)

type LongLiteral struct {
	*Literal
	Value spel.TypedValue
}

func (l *LongLiteral) GetLiteralValue() spel.TypedValue {
	return l.Value
}
