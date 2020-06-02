package expression

import (
	. "go-learning/expression/expression/common"
	"strings"
)

type TemplateAwareExpressionParser struct {
	Bracket
}

type Bracket struct {
	bracket string

	pos int
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
		suffixIndex := skipToCorrectEndSuffix(suffix, expressionString, afterPrefixIndex)
		if suffixIndex == -1 {

		}

		if suffixIndex == afterPrefixIndex {

		}
	}
	return nil
}

func skipToCorrectEndSuffix(suffix string, expressionString string, afterPrefixIndex int) int {
	pos := afterPrefixIndex
	maxlen := len(expressionString)
	nextSuffix := strings.Index(expressionString, suffix)
	if nextSuffix == -1 {
		return -1
	} else {
		stack := Node{}
		for pos < maxlen && (!isSuffixHere(expressionString, pos, suffix) || stack.Value != nil) {
			pos++
			ch := string(expressionString[pos])
			switch ch {
			case "'":
			case "\\":
				{
					endLiteral := indexOf(expressionString, ch, pos+1)
					if endLiteral == -1 {
						//err
					}
					pos = endLiteral
					break
				}
			case "(":
			case "[":
			case "{":
				{
					var bracket = Bracket{ch, pos}

					node := Node{Value: bracket}
					Push(node)
					break
				}
			case ")":
			case "]":
			case "}":
				{
					if stack.Value == nil {

					}
					pop, b := Pop(&stack)
					if b {
						bracket := pop.(Bracket)
						closeBracket := bracket.compatibleWithCloseBracket(ch)
						if !closeBracket {

						}
					}
				}
			}

		}
		//stack := Node{}
		if stack.Value != nil {
			//err
		} else {
			if !isSuffixHere(expressionString, pos, suffix) {
				return -1
			}
			return pos
		}
	}
	return 0
}

func isSuffixHere(expressionString string, pos int, suffix string) bool {
	suffixPosition := 0
	for i := 0; i < len(suffix) && pos < len(expressionString); i++ {
		s := string(expressionString[pos])
		suffixPosition++
		s2 := string(suffix[suffixPosition])
		if s != s2 {
			return false
		}
	}
	return suffixPosition == len(suffix)
}

func indexOf(expressionString string, ch string, fromIndex int) int {
	max := len(expressionString)
	if fromIndex < 0 {
		fromIndex = 0
	}
	if fromIndex >= max {
		return -1
	}
	for i := fromIndex; i < max; i++ {
		if string(expressionString[i]) == ch {
			return i
		}
	}
	return -1
}

func (b *Bracket) compatibleWithCloseBracket(ch string) bool {
	if b.bracket == "{" {
		return ch == "}"
	}
	if b.bracket == "[" {
		return ch == "]"
	}
	if b.bracket == "{" {
		return ch == "}"
	} else {
		return ch == ")"
	}
}
