package lexer

import (
	"testing"
	"text-adventure/token"
	"text-adventure/vocabulary"
)

func TestNextToken(t *testing.T) {
	tests := []struct {
		inputText     string
		expectedTypes []token.Token
	}{
		{
			"Go North",
			[]token.Token{
				{Type: token.VERB, Literal: "Go"},
				{Type: token.NOUN, Literal: "North"},
				{Type: token.EOF, Literal: string(byte(0))},
			},
		},
		{
			"go north",
			[]token.Token{
				{Type: token.VERB, Literal: "go"},
				{Type: token.NOUN, Literal: "north"},
				{Type: token.EOF, Literal: string(byte(0))},
			},
		},
		{
			"turn right",
			[]token.Token{
				{Type: token.VERB | token.NOUN, Literal: "turn"},
				{Type: token.NOUN, Literal: "right"},
				{Type: token.EOF, Literal: string(byte(0))},
			},
		},
		{
			"pick-up book",
			[]token.Token{
				{Type: token.VERB, Literal: "pick-up"},
				{Type: token.NOUN, Literal: "book"},
				{Type: token.EOF, Literal: string(byte(0))},
			},
		},
	}
	for i, tt := range tests {
		v := vocabulary.DefaultVocabulary()
		l := New(tt.inputText, v)
		for _, exToken := range tt.expectedTypes {
			tok := l.NextToken()
			if tok.Type != exToken.Type {
				t.Fatalf("tests[%d] - tokentype wrong. expected:%q, got:%q", i, exToken.Type.String(), tok.Type.String())
			}
			if tok.Literal != exToken.Literal {
				t.Fatalf("tests[%d] - tokenliteral wrong. expected:%q, got:%q", i, exToken.Literal, tok.Literal)
			}
		}
	}
}
