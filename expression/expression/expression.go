package expression

type Expression interface {
	getExpressionString() string

	getValue() interface{}
}
