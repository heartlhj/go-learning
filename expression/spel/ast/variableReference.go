package ast

import (
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

func (v VariableReference) GetValueInternal(state ExpressionState) TypedValue {
	return state.LookupVariable(v.Name)
}
