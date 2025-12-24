package token

type TokenType string

const (
	WORD = "WORD"
	VERB = "VERB"
	NOUN = "NOUN"

	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)

type Token struct {
	Type    TokenType
	Literal string
}
