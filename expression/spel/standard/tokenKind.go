package standard

type TokenKindType string

const (
	LPAREN TokenKindType = "("

	RPAREN TokenKindType = ")"

	COMMA TokenKindType = ","

	COLON TokenKindType = ":"

	HASH TokenKindType = "#"

	RSQUARE TokenKindType = "]"

	LSQUARE TokenKindType = "["

	LCURLY TokenKindType = "{"

	RCURLY TokenKindType = "}"

	DOT TokenKindType = "."

	PLUS TokenKindType = "+"

	STAR TokenKindType = "*"

	GE      TokenKindType = ">="
	GT      TokenKindType = ">"
	LE      TokenKindType = "<="
	LT      TokenKindType = "<"
	EQ      TokenKindType = "=="
	NE      TokenKindType = "!="
	PROJECT TokenKindType = "!["
	MOD     TokenKindType = "%"
	NOT     TokenKindType = "!"
	ASSIGN  TokenKindType = "="
	INC     TokenKindType = "++"
	DEC     TokenKindType = "--"

	MINUS  TokenKindType = "-"
	SELECT TokenKindType = "?["
	POWER  TokenKindType = "^"
	ELVIS  TokenKindType = "?:"

	SAFE_NAVI        TokenKindType = "?."
	BEAN_REF         TokenKindType = "@"
	FACTORY_BEAN_REF TokenKindType = "&"
	SYMBOLIC_OR      TokenKindType = "||"

	SYMBOLIC_AND TokenKindType = "&&"
	BETWEEN      TokenKindType = "between"

	SELECT_LAST TokenKindType = "$["

	IDENTIFIER

	LITERAL_INT

	LITERAL_LONG

	LITERAL_HEXINT

	LITERAL_HEXLONG

	LITERAL_STRING

	LITERAL_REAL

	LITERAL_REAL_FLOAT
)

type TokenKind struct {
	TokenChars    []rune
	HasPayload    bool
	TokenKindType TokenKindType
}

func (t *TokenKind) InitTokenKind() *TokenKind {
	t.TokenChars = []rune(t.TokenKindType)
	t.HasPayload = len(t.TokenChars) == 0
	return t
}
