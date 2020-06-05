package expression

import "go-learning/expression/spel"

//获取值
type Expression interface {
	GetExpressionString() string

	GetValue() interface{}

	GetValueContext(context spel.EvaluationContext) interface{}
}
