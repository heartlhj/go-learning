package ast

import (
	. "go-learning/expression/spel"
)

/**

 */
type StringLiteral struct {
	*Literal
}

func (l StringLiteral) GetValueInternal(expressionState ExpressionState) TypedValue {
	return l.Value
}
