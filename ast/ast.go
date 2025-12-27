package ast

import (
	"fmt"
	"text-adventure/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Command struct {
	Verb           *VerbGroup
	Object         *NounGroup
	IndirectObject *NounGroup
}

type NounGroup struct {
	Article    *token.Token
	Adjectives []token.Token
	Noun       token.Token
}

type VerbGroup struct {
	Verb     token.Token
	Particle *token.Token
}

func (c *Command) TokenLiteral() string {
	if c.Verb != nil {
		return c.Verb.TokenLiteral()
	}
	return ""
}

func (c *Command) String() string {
	result := c.Verb.String()
	if c.Object != nil {
		result += " " + c.Object.String()
	}
	if c.IndirectObject != nil {
		result += " " + c.IndirectObject.String()
	}
	return result
}

func (ng *NounGroup) TokenLiteral() string { return ng.Noun.Literal }
func (ng *NounGroup) String() string {
	result := ""
	if ng.Article != nil {
		result += ng.Article.Literal + " "
	}

	for _, adj := range ng.Adjectives {
		result += adj.Literal + " "
	}

	result += ng.Noun.Literal
	return result
}

func (vb *VerbGroup) TokenLiteral() string { return vb.Verb.Literal }
func (vb *VerbGroup) String() string {
	if vb.Particle != nil {
		return vb.Verb.Literal + " " + vb.Particle.Literal
	}
	return vb.Verb.Literal
}

type ParseError struct {
	Message string
	Token   string
}

func (pe *ParseError) Error() string {
	return fmt.Sprintf("parse error: %s (at token: %q)", pe.Message, pe.Token)
}
