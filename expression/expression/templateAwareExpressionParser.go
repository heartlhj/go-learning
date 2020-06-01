package expression

import "strings"

type TemplateAwareExpressionParser struct {
}

func (t *TemplateAwareExpressionParser) parseExpression(expressionString string) Expression {
	return nil
}

func parseTemplate(expressionString string, context ParserContext) Expression {
	if expressionString == "" {
		return &LiteralExpression{""}
	}
	expressions := parseExpressions(expressionString, context)
	if len(expressions) == 1 {
		return expressions[0]
	}
	return &CompositeStringExpression{expressionString, expressions}
}

func parseExpressions(expressionString string, context ParserContext) []Expression {
	expressions := make([]Expression, 0)
	prefix := context.getExpressionPrefix()
	suffix := context.getExpressionSuffix()
	startIdx := 0
	if startIdx < len(expressionString) {
		prefixIndex := strings.Index(expressionString, prefix)
		if prefixIndex >= startIdx {
			runes := []rune(expressionString)
			expressions = append(expressions, &LiteralExpression{string(runes[startIdx:prefixIndex])})
		}
		afterPrefixIndex := prefixIndex + len(prefix)

	}
}

func skipToCorrectEndSuffix(suffix string, expressionString string, afterPrefixIndex int) int {
	pos := afterPrefixIndex
	maxlen := len(expressionString)
	nextSuffix := strings.Index(expressionString, suffix)
	if nextSuffix == -1 {
		return -1
	}

}
