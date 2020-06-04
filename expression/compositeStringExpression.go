package expression

import "go-learning/expression/spel"

type CompositeStringExpression struct {
	*spel.ExpressionImpl
	ExpressionString string
	Expressions      []Expression
}

func (c *CompositeStringExpression) GetExpressionString() string {
	return c.ExpressionString
}

func (c *CompositeStringExpression) GetValue() interface{} {
	//s := ""

	return "c.literalValue"
}
