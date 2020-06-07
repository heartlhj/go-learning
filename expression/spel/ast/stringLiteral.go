package ast

import (
	. "github.com/heartlhj/go-learning/expression/spel"
)

/**

 */
type StringLiteral struct {
	*Literal
}

func (l StringLiteral) GetValueInternal(expressionState ExpressionState) TypedValue {
	return l.Value
}
