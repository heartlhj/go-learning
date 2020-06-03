package standard

import (
	"fmt"
	. "go-learning/expression/utils"
	"strings"
)

const (
	IS_DIGIT   = 0x01
	S_HEXDIGIT = 0x02
	IS_ALPHA   = 0x04
)

var FLADS = [256]rune{}

type Tokenizer struct {
	ExpressionString string
	charsToProcess   []rune
	pos              int
	max              int
	tokens           []Token
}

func init() {
	var ch rune = 0
	for ; ch <= 9; ch++ {

	}
}

func (t *Tokenizer) Process() []Token {
	if t.pos < t.max {
		ch := t.charsToProcess[t.pos]
		if isAlphabetic(ch) {

		} else {
			st := string(ch)
			switch st {

			case "+":
				if t.isTwoCharToken(INC) {
					t.pushPairToken(TokenKind{tokenKindType: INC})
				} else {
					t.pushCharToken(TokenKind{tokenKindType: INC})
				}
				break
			case "-":
				if t.isTwoCharToken(TokenKind{tokenKindType: DEC}) {
					t.pushPairToken(TokenKind{tokenKindType: DEC})
				} else {
					t.pushCharToken(TokenKind{tokenKindType: MINUS})
				}
				break
			case "(":
				t.pushCharToken(TokenKind{tokenKindType: LPAREN})
				break
			case ")":
				t.pushCharToken(TokenKind{tokenKindType: RPAREN})
				break
			case "[":
				t.pushCharToken(TokenKind{tokenKindType: LSQUARE})
				break
			case "#":
				t.pushCharToken(TokenKind{tokenKindType: HASH})
				break
			case "]":
				t.pushCharToken(TokenKind{tokenKindType: RSQUARE})
				break
			case "{":
				t.pushCharToken(TokenKind{tokenKindType: LCURLY})
				break
			case "}":
				t.pushCharToken(TokenKind{tokenKindType: RCURLY})
				break
			case "@":
				t.pushCharToken(TokenKind{tokenKindType: BEAN_REF})
				break
			case "!":
				if t.isTwoCharToken(TokenKind{tokenKindType: NE}) {
					t.pushPairToken(TokenKind{tokenKindType: NE})
				} else if t.isTwoCharToken(TokenKind{tokenKindType: PROJECT}) {
					t.pushPairToken(TokenKind{tokenKindType: PROJECT})
				} else {
					t.pushCharToken(TokenKind{tokenKindType: NOT})
				}
				break
			case "=":
				if t.isTwoCharToken(TokenKind{tokenKindType: EQ}) {
					t.pushPairToken(TokenKind{tokenKindType: EQ})
				} else {
					t.pushCharToken(TokenKind{tokenKindType: ASSIGN})
				}
				break
			case "&":
				if t.isTwoCharToken(TokenKind{tokenKindType: SYMBOLIC_AND}) {
					t.pushPairToken(TokenKind{tokenKindType: SYMBOLIC_AND})
				} else {
					t.pushCharToken(TokenKind{tokenKindType: FACTORY_BEAN_REF})
				}
				break
			case "|":
				if !t.isTwoCharToken(TokenKind{tokenKindType: SYMBOLIC_OR}) {
					fmt.Errorf("")
				}
				t.pushPairToken(TokenKind{tokenKindType: SYMBOLIC_OR})
				break
			case "$":
				if t.isTwoCharToken(TokenKind{tokenKindType: SELECT_LAST}) {
					t.pushPairToken(TokenKind{tokenKindType: SELECT_LAST})
				} else {
					t.lexIdentifier()
				}
				break
			case ">":
				if t.isTwoCharToken(TokenKind{tokenKindType: GE}) {
					t.pushPairToken(TokenKind{tokenKindType: GE})
				} else {
					t.pushCharToken(TokenKind{tokenKindType: GE})
				}
				break
			case "<":
				if t.isTwoCharToken(TokenKind{tokenKindType: LE}) {
					t.pushPairToken(TokenKind{tokenKindType: LE})
				} else {
					t.pushCharToken(TokenKind{tokenKindType: LE})
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
			}

		}
	}
	return t.tokens
}

func (t *Tokenizer) isTwoCharToken(kind TokenKind) bool {
	return len(kind.tokenChars) == 2 && t.charsToProcess[t.pos] == kind.tokenChars[0] &&
		t.charsToProcess[t.pos+1] == kind.tokenChars[1]
}

func (t *Tokenizer) pushPairToken(kind TokenKind) {
	t.tokens = append(t.tokens, Token{kind: kind, startPos: t.pos, endPos: t.pos + 1})
	t.pos++
}

func (t *Tokenizer) pushCharToken(kind TokenKind) {
	t.tokens = append(t.tokens, Token{kind: kind, startPos: t.pos, endPos: t.pos + 2})
	t.pos += 2
}

func (t *Tokenizer) pushOneCharOrTwoCharToken(kind TokenKind, pos int, data []rune) {
	t.tokens = append(t.tokens, Token{kind: kind, startPos: pos, data: string(data), endPos: t.pos + 2})
}

func isAlphabetic(ch rune) bool {
	char := int(ch)
	if char > 255 {
		return false
	}
	return true
}

func (t *Tokenizer) lexIdentifier() {
	start := t.pos
	for isAlphabetic(t.charsToProcess[t.pos]) {
		t.pos++
	}
	runes := t.charsToProcess[start:t.pos]
	alternative_operator_names := []string{"DIV", "EQ", "GE", "GT", "LE", "LT", "MOD", "NE", "NOT"}
	if (t.pos-start) == 2 || (t.pos-start) == 3 {
		asString := strings.ToUpper(string(runes))
		idx := BinarySearch(alternative_operator_names, asString)
		if idx >= 0 {
			//t.pushOneCharOrTwoCharToken(TokenKind.valueOf(asString), start, runes)
			return
		}
	}
	t.tokens = append(t.tokens, Token{TokenKind{tokenKindType: IDENTIFIER}, string(runes), start, t.pos})
}
