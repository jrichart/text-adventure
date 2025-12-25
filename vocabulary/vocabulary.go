package vocabulary

import (
	"strings"
	"text-adventure/token"
)

type Vocabulary struct {
	words map[string]token.TokenType
}

func New() *Vocabulary {
	return &Vocabulary{
		words: make(map[string]token.TokenType),
	}
}

func (v *Vocabulary) LookupWord(word string) token.Token {
	lower := strings.ToLower(word)
	tokenType, exists := v.words[lower]
	if !exists {
		tokenType = token.ILLEGAL
	}
	return token.Token{
		Literal: word,
		Type:    tokenType,
	}
}

// The Vocabulary that is available at the start of the game.
// More could be added based on interactions during the game
func DefaultVocabulary() *Vocabulary {
	v := New()

	var vocabulary = map[string]token.TokenType{
		// Directions
		"north": token.NOUN,
		"south": token.NOUN,
		"east":  token.NOUN,
		"west":  token.NOUN,
		"up":    token.NOUN | token.PARTICLE,
		"down":  token.NOUN,
		"left":  token.NOUN,
		"right": token.NOUN,

		// items
		"book":     token.NOUN,
		"shelf":    token.NOUN,
		"room":     token.NOUN,
		"desk":     token.NOUN,
		"skeleton": token.NOUN,
		"paper":    token.NOUN,
		"note":     token.NOUN,
		"door":     token.NOUN,
		"lamp":     token.NOUN,
		"rock":     token.NOUN,
		"sword":    token.NOUN,

		// actions
		"get":   token.VERB,
		"go":    token.VERB,
		"drop":  token.VERB,
		"look":  token.VERB,
		"take":  token.VERB,
		"turn":  token.VERB | token.NOUN,
		"light": token.VERB | token.NOUN,
		"walk":  token.VERB | token.NOUN,

		// acticles
		"the": token.ARTICLE,

		// prepositions
		"in":   token.PREPOSITION,
		"with": token.PREPOSITION,

		// TODO: turn into verb and particle
		"pick-up": token.VERB,
	}
	v.words = vocabulary
	return v
}
