package standard

import (
	"fmt"
	. "go-learning/expression/utils"
	"strings"
)

const (
	IS_DIGIT    = 0x01
	IS_HEXDIGIT = 0x02
	IS_ALPHA    = 0x04
)

var FLAGS = [256]rune{}

type Tokenizer struct {
	ExpressionString string
	charsToProcess   []rune
	pos              int
	max              int
	tokens           []Token
}

func (t *Tokenizer) InitTokenizer() {
	t.initFlags()
	expressionString := t.ExpressionString
	runes := []rune(expressionString)
	t.charsToProcess = runes
	t.max = len(t.charsToProcess)
	t.pos = 0
}

func (t *Tokenizer) initFlags() {

	for ch := '0'; ch <= '9'; ch++ {
		FLAGS[ch] |= IS_DIGIT | IS_HEXDIGIT
	}
	for ch := 'A'; ch <= 'F'; ch++ {
		FLAGS[ch] |= IS_DIGIT | IS_HEXDIGIT
	}
	for ch := 'a'; ch <= 'f'; ch++ {
		FLAGS[ch] |= IS_HEXDIGIT
	}
	for ch := 'A'; ch <= 'Z'; ch++ {
		FLAGS[ch] |= IS_ALPHA
	}
	for ch := 'a'; ch <= 'z'; ch++ {
		FLAGS[ch] |= IS_ALPHA
	}

}

func (t *Tokenizer) Process() []Token {
	for t.pos < t.max {
		ch := t.charsToProcess[t.pos]
		if isAlphabetic(ch) {
			t.lexIdentifier()
		} else {
			st := string(ch)
			switch st {

			case "+":
				if t.isTwoCharToken(TokenKind{TokenKindType: INC, TokenChars: []rune(INC), HasPayload: len([]rune(INC)) == 0}) {
					t.pushPairToken(TokenKind{TokenKindType: INC, TokenChars: []rune(INC), HasPayload: len([]rune(INC)) == 0})
				} else {
					t.pushCharToken(TokenKind{TokenKindType: INC, TokenChars: []rune(INC), HasPayload: len([]rune(INC)) == 0})
				}
				break
			case "-":
				if t.isTwoCharToken(TokenKind{TokenKindType: DEC, TokenChars: []rune(DEC), HasPayload: len([]rune(DEC)) == 0}) {
					t.pushPairToken(TokenKind{TokenKindType: DEC, TokenChars: []rune(DEC), HasPayload: len([]rune(DEC)) == 0})
				} else {
					t.pushCharToken(TokenKind{TokenKindType: MINUS, TokenChars: []rune(MINUS), HasPayload: len([]rune(MINUS)) == 0})
				}
				break
			case "(":
				t.pushCharToken(TokenKind{TokenKindType: LPAREN, TokenChars: []rune(LPAREN), HasPayload: len([]rune(LPAREN)) == 0})
				break
			case ")":
				t.pushCharToken(TokenKind{TokenKindType: RPAREN, TokenChars: []rune(RPAREN), HasPayload: len([]rune(RPAREN)) == 0})
				break
			case "[":
				t.pushCharToken(TokenKind{TokenKindType: LSQUARE, TokenChars: []rune(LSQUARE), HasPayload: len([]rune(LSQUARE)) == 0})
				break
			case "#":
				t.pushCharToken(TokenKind{TokenKindType: HASH, TokenChars: []rune(HASH), HasPayload: len([]rune(HASH)) == 0})
				break
			case "]":
				t.pushCharToken(TokenKind{TokenKindType: RSQUARE, TokenChars: []rune(RSQUARE), HasPayload: len([]rune(RSQUARE)) == 0})
				break
			case "{":
				t.pushCharToken(TokenKind{TokenKindType: LCURLY, TokenChars: []rune(LCURLY), HasPayload: len([]rune(LCURLY)) == 0})
				break
			case "}":
				t.pushCharToken(TokenKind{TokenKindType: RCURLY, TokenChars: []rune(RCURLY), HasPayload: len([]rune(RCURLY)) == 0})
				break
			case "@":
				t.pushCharToken(TokenKind{TokenKindType: BEAN_REF, TokenChars: []rune(BEAN_REF), HasPayload: len([]rune(BEAN_REF)) == 0})
				break
			case "!":
				if t.isTwoCharToken(TokenKind{TokenKindType: NE, TokenChars: []rune(NE), HasPayload: len([]rune(NE)) == 0}) {
					t.pushPairToken(TokenKind{TokenKindType: NE, TokenChars: []rune(NE), HasPayload: len([]rune(NE)) == 0})
				} else if t.isTwoCharToken(TokenKind{TokenKindType: NE, TokenChars: []rune(NE), HasPayload: len([]rune(NE)) == 0}) {
					t.pushPairToken(TokenKind{TokenKindType: PROJECT, TokenChars: []rune(NE), HasPayload: len([]rune(NE)) == 0})
				} else {
					t.pushCharToken(TokenKind{TokenKindType: NOT, TokenChars: []rune(NOT), HasPayload: len([]rune(NOT)) == 0})
				}
				break
			case "=":
				if t.isTwoCharToken(TokenKind{TokenKindType: EQ, TokenChars: []rune(EQ), HasPayload: len([]rune(EQ)) == 0}) {
					t.pushPairToken(TokenKind{TokenKindType: EQ, TokenChars: []rune(EQ), HasPayload: len([]rune(EQ)) == 0})
				} else {
					t.pushCharToken(TokenKind{TokenKindType: ASSIGN, TokenChars: []rune(ASSIGN), HasPayload: len([]rune(ASSIGN)) == 0})
				}
				break
			case "&":
				if t.isTwoCharToken(TokenKind{TokenKindType: SYMBOLIC_AND, TokenChars: []rune(SYMBOLIC_AND), HasPayload: len([]rune(SYMBOLIC_AND)) == 0}) {
					t.pushPairToken(TokenKind{TokenKindType: SYMBOLIC_AND, TokenChars: []rune(SYMBOLIC_AND), HasPayload: len([]rune(SYMBOLIC_AND)) == 0})
				} else {
					t.pushCharToken(TokenKind{TokenKindType: FACTORY_BEAN_REF, TokenChars: []rune(FACTORY_BEAN_REF), HasPayload: len([]rune(FACTORY_BEAN_REF)) == 0})
				}
				break
			case "|":
				if !t.isTwoCharToken(TokenKind{TokenKindType: SYMBOLIC_OR, TokenChars: []rune(SYMBOLIC_OR), HasPayload: len([]rune(SYMBOLIC_OR)) == 0}) {
					fmt.Errorf("")
				}
				t.pushPairToken(TokenKind{TokenKindType: SYMBOLIC_OR, TokenChars: []rune(SYMBOLIC_OR), HasPayload: len([]rune(SYMBOLIC_OR)) == 0})
				break
			case "$":
				if t.isTwoCharToken(TokenKind{TokenKindType: SELECT_LAST, TokenChars: []rune(SELECT_LAST), HasPayload: len([]rune(SELECT_LAST)) == 0}) {
					t.pushPairToken(TokenKind{TokenKindType: SELECT_LAST, TokenChars: []rune(SELECT_LAST), HasPayload: len([]rune(SELECT_LAST)) == 0})
				} else {
					t.lexIdentifier()
				}
				break
			case ">":
				if t.isTwoCharToken(TokenKind{TokenKindType: GE, TokenChars: []rune(GE), HasPayload: len([]rune(GE)) == 0}) {
					t.pushPairToken(TokenKind{TokenKindType: GE, TokenChars: []rune(GE), HasPayload: len([]rune(GE)) == 0})
				} else {
					t.pushCharToken(TokenKind{TokenKindType: GE, TokenChars: []rune(GE), HasPayload: len([]rune(GE)) == 0})
				}
				break
			case "<":
				if t.isTwoCharToken(TokenKind{TokenKindType: LE, TokenChars: []rune(LE), HasPayload: len([]rune(LE)) == 0}) {
					t.pushPairToken(TokenKind{TokenKindType: LE, TokenChars: []rune(LE), HasPayload: len([]rune(LE)) == 0})
				} else {
					t.pushCharToken(TokenKind{TokenKindType: LE, TokenChars: []rune(LE), HasPayload: len([]rune(LE)) == 0})
				}
				break
			case "0":
			case "1":
			case "2":
			case "3":
			case "4":
			case "5":
			case "6":
			case "7":
			case "8":
			case "9":
				//
				break
			case " ":
			case "\t":
			case "\r":
			case "\n":
				t.pos++
				break
			case "'":
				t.lexQuotedStringLiteral()
				break
			case "\"":
				t.lexDoubleQuotedStringLiteral()
				break
			case string(0):
				t.pos++
				break
			}

		}
	}
	return t.tokens
}

func (t *Tokenizer) isTwoCharToken(kind TokenKind) bool {
	return len(kind.TokenChars) == 2 && t.charsToProcess[t.pos] == kind.TokenChars[0] &&
		t.charsToProcess[t.pos+1] == kind.TokenChars[1]
}

func (t *Tokenizer) pushPairToken(kind TokenKind) {
	t.tokens = append(t.tokens, Token{Kind: kind, StartPos: t.pos, EndPos: t.pos + 2})
	t.pos += 2
}

func (t *Tokenizer) pushCharToken(kind TokenKind) {
	t.tokens = append(t.tokens, Token{Kind: kind, StartPos: t.pos, EndPos: t.pos + 1})
	t.pos++
}

func (t *Tokenizer) pushOneCharOrTwoCharToken(kind TokenKind, pos int, data []rune) {
	t.tokens = append(t.tokens, Token{Kind: kind, StartPos: pos, Data: string(data), EndPos: pos + len(kind.TokenKindType)})
}

func (t *Tokenizer) pushIntToken(data []rune, start int, end int) {
	kind := TokenKind{TokenKindType: LITERAL_INT, TokenChars: []rune(LITERAL_INT), HasPayload: len([]rune(LITERAL_INT)) == 0}
	t.tokens = append(t.tokens, Token{Kind: kind, StartPos: start, Data: string(data), EndPos: end})
}

func isAlphabetic(ch rune) bool {
	char := int(ch)
	if char > 255 {
		return false
	}
	return (FLAGS[ch] & IS_ALPHA) != 0
}

func isDigit(ch rune) bool {
	char := int(ch)
	if char > 255 {
		return false
	}
	return (FLAGS[ch] & IS_DIGIT) != 0
}

func isIdentifier(ch rune) bool {
	return isAlphabetic(ch) || isDigit(ch) || ch == '_' || ch == '$'
}

func isHexadecimalDigit(ch rune) bool {
	char := int(ch)
	if char > 255 {
		return false
	}
	return (FLAGS[ch] & IS_HEXDIGIT) != 0
}

func (t *Tokenizer) lexIdentifier() {
	start := t.pos
	t.pos++
	for isIdentifier(t.charsToProcess[t.pos]) {
		t.pos++
	}
	runes := t.charsToProcess[start:t.pos]
	alternativeOperatorNames := []string{"DIV", "EQ", "GE", "GT", "LE", "LT", "MOD", "NE", "NOT"}
	if (t.pos-start) == 2 || (t.pos-start) == 3 {
		asString := strings.ToUpper(string(runes))
		idx := BinarySearch(alternativeOperatorNames, asString)
		if idx >= 0 {
			//t.pushOneCharOrTwoCharToken(TokenKind.valueOf(asString), start, runes)
			return
		}
	}
	t.tokens = append(t.tokens, Token{TokenKind{TokenKindType: IDENTIFIER, TokenChars: []rune(IDENTIFIER), HasPayload: len([]rune(IDENTIFIER)) == 0}, string(runes), start, t.pos})
}

func (t *Tokenizer) lexNumericLiteral(ch rune) {
	//start := t.pos
	//i,_ := strconv.Atoi(string(r))
	//t.pos++
	//endOfNumber := t.pos
	//r := t.charsToProcess[t.pos]
	//

}

func (t *Tokenizer) lexQuotedStringLiteral() {
	start := t.pos
	terminated := false
	for !terminated {
		t.pos++
		ch := t.charsToProcess[t.pos]
		if string(ch) == "'" {
			if t.pos < t.max-1 {
				if string(t.charsToProcess[t.pos+1]) == "'" {
					t.pos++
				} else {
					terminated = true
				}
			}
			terminated = true
			if t.pos == t.max {
				panic("Cannot find terminating '' for string")
			}
		}
	}
	t.pos++
	process := t.charsToProcess[start:t.pos]
	t.tokens = append(t.tokens, Token{TokenKind{TokenKindType: LITERAL_STRING, TokenChars: []rune(LITERAL_STRING), HasPayload: len([]rune(LITERAL_STRING)) == 0}, string(process), start, t.pos})
}

func (t *Tokenizer) lexDoubleQuotedStringLiteral() {
	start := t.pos
	terminated := false
	for !terminated {
		t.pos++
		ch := t.charsToProcess[t.pos]
		if string(ch) == "\"" {
			if string(t.charsToProcess[t.pos+1]) == "\"" {
				t.pos++
			} else {
				terminated = true
			}

			if t.pos == t.max-1 {
				panic("Cannot find terminating \" for string")
			}
		}
	}
	t.pos++
	process := t.charsToProcess[start:t.pos]
	t.tokens = append(t.tokens, Token{TokenKind{TokenKindType: LITERAL_STRING, TokenChars: []rune(LITERAL_STRING), HasPayload: len([]rune(LITERAL_STRING)) == 0}, string(process), start, t.pos})
}
func (t *Tokenizer) isChar(a rune, b rune) bool {
	r := t.charsToProcess[t.pos]
	return r == a || r == b

}
