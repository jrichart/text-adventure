package token

type TokenType int64

const (
	ILLEGAL TokenType = 0
	VERB    TokenType = 1 << iota
	NOUN
	ARTICLE
	ADJECTIVE
	PREPOSITION
	PARTICLE
	ADVERB
	EOF
)

type Token struct {
	Type    TokenType
	Literal string
}

func (t *Token) HasType(tokenType TokenType) bool {
	return t.Type&tokenType != 0
}

func (tt *TokenType) String() string {
	switch *tt {
	case ILLEGAL:
		return "ILLEGAL"
	case VERB:
		return "VERB"
	case NOUN:
		return "NOUN"
	case ARTICLE:
		return "ARTICLE"
	case ADJECTIVE:
		return "ADJECTIVE"
	case PREPOSITION:
		return "PREPOSITION"
	case PARTICLE:
		return "PARTICLE"
	case ADVERB:
		return "ADVERB"
	case VERB | NOUN:
		return "VERB or NOUN"
	case EOF:
		return "EOF"
	default:
		return "UNKNOWN"
	}
}
