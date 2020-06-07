package expression

import "github.com/heartlhj/go-learning/expression/spel"

type LiteralExpression struct {
	*spel.ExpressionImpl
	literalValue string
}

func (l *LiteralExpression) GetExpressionString() string {
	return l.literalValue
}

func (l *LiteralExpression) GetValue() interface{} {
	return l.literalValue
}
