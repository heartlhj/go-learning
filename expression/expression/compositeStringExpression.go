package expression

type CompositeStringExpression struct {
	expressionString string
	expressions      []Expression
}

func (c *CompositeStringExpression) getExpressionString() string {
	return c.expressionString
}

func (c *CompositeStringExpression) getValue() interface{} {
	//s := ""

	return "c.literalValue"
}
