package expression

type TemplateParserContext struct {
	expressionPrefix string

	expressionSuffix string
}

func (t *TemplateParserContext) getExpressionPrefix() string {
	return t.expressionPrefix
}

func (t *TemplateParserContext) getExpressionSuffix() string {
	return t.expressionSuffix
}
