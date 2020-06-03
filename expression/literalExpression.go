package expression

type LiteralExpression struct {
	literalValue string
}

func (l *LiteralExpression) getExpressionString() string {
	return l.literalValue
}

func (l *LiteralExpression) getValue() interface{} {
	return l.literalValue
}
