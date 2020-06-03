package expression

import (
	"container/list"
	"fmt"
	. "go-learning/expression/err"
	. "go-learning/expression/spel"
	. "go-learning/expression/spel/ast"
	. "go-learning/expression/spel/standard"
)

type InternalSpelExpressionParser struct {
	*TemplateAwareExpressionParser

	expressionString string

	tokenStreamLength int

	tokenStreamPointer int

	configuration SpelParserConfiguration

	tokenStream []Token

	constructedNodes list.List
}

func (i *InternalSpelExpressionParser) doParseExpression(expressionString string) Expression {

	i.expressionString = expressionString
	tokenizer := Tokenizer{ExpressionString: expressionString}
	i.tokenStream = tokenizer.Process()
	i.tokenStreamLength = len(i.tokenStream)
	i.tokenStreamPointer = 0

	return nil
}

func (i *InternalSpelExpressionParser) takeToken() Token {
	if i.tokenStreamPointer >= i.tokenStreamLength {
		fmt.Errorf("No token")
	}
	token := i.tokenStream[i.tokenStreamPointer]
	i.tokenStreamPointer++
	return token
}

func (i *InternalSpelExpressionParser) nextToken() (Token, error) {
	if i.tokenStreamPointer >= i.tokenStreamLength {
		return Token{}, ExpressionErr{Code: "123", Msg: "321"}
	}
	token := i.tokenStream[i.tokenStreamPointer]
	i.tokenStreamPointer++
	return token, nil
}

func (i *InternalSpelExpressionParser) peekToken() (Token, error) {
	if i.tokenStreamPointer >= i.tokenStreamLength {
		return Token{}, ExpressionErr{Code: "123", Msg: "321"}
	}
	token := i.tokenStream[i.tokenStreamPointer]
	return token, nil
}

func eatExpression() SpelNodeImpl {
	return SpelNodeImpl{}
}

func (i *InternalSpelExpressionParser) maybeEatRelationalOperator() (Token, error) {
	token, err := i.peekToken()
	if err != nil {
		if token.IsNumericRelationalOperator() {
			return token, nil
		}
	}
	return token, err
}

func (i *InternalSpelExpressionParser) eatRelationalExpression() (SpelNode, error) {
	expr, err := i.eatSumExpression()
	relationalOperatorToken, err1 := i.maybeEatRelationalOperator()
	if err1 != nil {
		t := i.takeToken()
		rhExpr, err3 := i.eatSumExpression()
		if err == nil {
			panic("Problem parsing left operand")
		}
		if err3 == nil {
			panic("Problem parsing right operand")
		}
		kindType := relationalOperatorToken.Kind.TokenKindType

		if relationalOperatorToken.IsNumericRelationalOperator() {
			pos := toPos(t.StartPos, t.EndPos)
			if kindType == EQ {
				eq := OpEQ{}
				eq.Pos = pos
				var nodes []SpelNode
				nodes[0] = expr
				nodes[1] = rhExpr
				eq.Children = nodes
				eq.Parent = eq
				return eq, nil
			}
		}
	}
	return expr, nil
}

func (i *InternalSpelExpressionParser) eatSumExpression() (SpelNode, error) {
	expr, err := i.eatProductExpression()
	if err != nil && i.peekTokenTwo(INC, DEC) {

	}
	return expr, nil
}

func (i *InternalSpelExpressionParser) eatProductExpression() (SpelNode, error) {
	expr, err := i.eatPowerIncDecExpression()
	if err != nil && i.peekTokenTwo(INC, DEC) {

	}
	return expr, nil
}

func (i *InternalSpelExpressionParser) eatPowerIncDecExpression() (SpelNode, error) {
	expr, err := i.eatUnaryExpression()
	if err != nil && i.peekTokenTwo(INC, DEC) {

	}
	return expr, nil
}

func (i *InternalSpelExpressionParser) eatUnaryExpression() (SpelNode, error) {
	if i.peekTokens(PLUS, MINUS, NOT) {
		t := i.takeToken()
		_, err := i.eatUnaryExpression()
		if err != nil {
			panic("No node")
		}
		if t.Kind.TokenKindType == NOT {
		}
	}
	return i.eatPrimaryExpression()
}

func (i *InternalSpelExpressionParser) eatPrimaryExpression() (SpelNode, error) {
	node, err := i.eatStartNode()
	if err != nil {
		return node, err
	}
	return &SpelNodeImpl{}, nil
}

func (i *InternalSpelExpressionParser) eatStartNode() (SpelNode, error) {
	if i.maybeEatFunctionOrVar() {
		return i.pop(), nil
	}
	return &SpelNodeImpl{}, nil
}

func (i *InternalSpelExpressionParser) maybeEatFunctionOrVar() bool {
	if !i.peekTokenOnly(HASH) {
		return true
	}
	token := i.takeToken()

	functionOrVariableName := i.eatToken(IDENTIFIER)
	args := i.maybeEatMethodArgs()
	if args == nil {
		reference := VariableReference{Name: functionOrVariableName.StringValue()}
		reference.Pos = toPos(token.StartPos, functionOrVariableName.EndPos)
		i.push(reference)
		return true
	}
	return true
}

func toPos(start int, end int) int {
	return (start << 16) + end
}

func (i *InternalSpelExpressionParser) push(newNode SpelNode) {
	i.constructedNodes.PushBack(newNode)
}

func (i *InternalSpelExpressionParser) pop() SpelNode {
	return i.constructedNodes.Back().Value.(SpelNode)
}

func (i *InternalSpelExpressionParser) maybeEatMethodArgs() []SpelNodeImpl {
	if !i.peekTokenOnly(LPAREN) {
		return nil
	}
	//args := make([]SpelNodeImpl,0)
	return nil
}

func (i *InternalSpelExpressionParser) consumeArguments(accumulatedArguments []SpelNodeImpl) {
	token, err := i.peekToken()
	if err != nil {
		panic("Expected token")
	}
	//var next Token
	i.nextToken()
	token, err = i.peekToken()
	if err == nil {
		panic("Unexpectedly ran out of arguments")
	}
	if token.Kind.TokenKindType != RPAREN {
		accumulatedArguments = append(accumulatedArguments)
	}

}
func (i *InternalSpelExpressionParser) eatExpression() SpelNodeImpl {
	return SpelNodeImpl{}
}

func (i *InternalSpelExpressionParser) eatLogicalOrExpression() SpelNodeImpl {
	return SpelNodeImpl{}
}

func (i *InternalSpelExpressionParser) eatToken(expectedKind TokenKindType) Token {
	token, err := i.nextToken()
	if err != nil {
		panic("Unexpectedly ran out of input")
	}
	if token.Kind.TokenKindType != expectedKind {
		panic("Unexpected token.")
	}
	return token
}
func (i *InternalSpelExpressionParser) peekTokenOnly(possible1 TokenKindType) bool {
	token, err := i.peekToken()
	if err != nil {
		return false
	}
	return token.Kind.TokenKindType == possible1
}

func (i *InternalSpelExpressionParser) peekTokenTwo(possible1 TokenKindType, possible2 TokenKindType) bool {
	token, err := i.peekToken()
	if err != nil {
		return false
	}
	return (token.Kind.TokenKindType == possible1) || token.Kind.TokenKindType == possible2
}
func (i *InternalSpelExpressionParser) peekTokens(possible1 TokenKindType, possible2 TokenKindType, possible3 TokenKindType) bool {
	token, err := i.peekToken()
	if err != nil {
		return false
	}
	return (token.Kind.TokenKindType == possible1) || token.Kind.TokenKindType == possible2 || token.Kind.TokenKindType == possible3
}
