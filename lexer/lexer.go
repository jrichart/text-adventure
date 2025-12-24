package lexer

import (
	"strings"
	"text-adventure/token"
	"text-adventure/vocabulary"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
	vocab        *vocabulary.Vocabulary
}

func New(input string, vocab *vocabulary.Vocabulary) *Lexer {
	l := &Lexer{
		input: input,
		vocab: vocab,
	}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()
	if l.ch == 0 {
		tok = newToken(token.EOF, l.ch)
		return tok
	}
	if isLetter(l.ch) {
		tok.Literal = l.readWord()
		tok.Type = l.classifyWord(tok.Literal)
		return tok
	} else {
		tok = newToken(token.ILLEGAL, l.ch)
	}
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readWord() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) classifyWord(word string) token.TokenType {
	lower := strings.ToLower(word)
	if l.vocab.IsVerb(lower) {
		return token.VERB
	}
	if l.vocab.IsNoun(lower) {
		return token.NOUN
	}
	return token.WORD
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '-'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
