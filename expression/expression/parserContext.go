package expression

type ParserContext interface {
	getExpressionPrefix() string

	getExpressionSuffix() string
}
