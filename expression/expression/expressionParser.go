package expression

type ExpressionParser interface {
	parseExpression(var1 string) Expression

	doParseExpression(var1 string) Expression
}
