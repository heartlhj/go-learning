package ast

import (
	. "go-learning/expression"
	. "go-learning/expression/spel"
)

const (
	THIS = "this"
	ROOT = "root"
)

type VariableReference struct {
	*SpelNodeImpl
	Name string
}

func (v VariableReference) getValueInternal(state ExpressionState) TypedValue {
	return TypedValue{}
}
