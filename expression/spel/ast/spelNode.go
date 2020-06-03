package ast

import "go-learning/expression/spel"

type SpelNode interface {
	getValue(expressionState spel.ExpressionState) interface{}
}
