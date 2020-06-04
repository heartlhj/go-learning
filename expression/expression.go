package expression

import "go-learning/expression/spel"

type Expression interface {
	GetExpressionString() string

	GetValue() interface{}

	GetValueContext(context spel.EvaluationContext) interface{}
}
