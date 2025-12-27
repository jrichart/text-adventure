package parser

import (
	"fmt"
	"text-adventure/ast"
	"text-adventure/lexer"
	"text-adventure/token"
)

type Parser struct {
	lexer     *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:  l,
		errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.HasType(t)
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.HasType(t)
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) Errors() []error {
	return p.Errors()
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t.String(), p.peekToken.Type.String())
	p.errors = append(p.errors, msg)
}

func (p *Parser) ParseCommand() *ast.Command {
	cmd := &ast.Command{}
	for !p.curTokenIs(token.EOF) {
		cmd.Verb = p.parseVerbGroup()
		if cmd.Verb == nil {
			return nil
		}

		if p.curTokenIs(token.NOUN | token.ADJECTIVE | token.ARTICLE) {
			cmd.Object = p.parseNounGroup()
		}
	}
	return cmd
}

func (p *Parser) parseVerbGroup() *ast.VerbGroup {
	vg := &ast.VerbGroup{}

	if !p.curTokenIs(token.VERB) {
		p.addError(fmt.Sprintf("expected type of verb, got: %s for %s", p.curToken.Type.String(), p.curToken.Literal))
		return nil
	}
	vg.Verb = p.curToken
	p.nextToken()

	if p.curTokenIs(token.PARTICLE) {
		particle := p.curToken
		vg.Particle = &particle
		p.nextToken()
	}

	return vg
}

func (p *Parser) parseNounGroup() *ast.NounGroup {
	ng := &ast.NounGroup{}

	if p.curTokenIs(token.ARTICLE) {
		article := &token.Token{Type: token.ARTICLE, Literal: p.curToken.Literal}
		ng.Article = article
		p.nextToken()
	}
	if !p.curTokenIs(token.NOUN) {
		p.addError(fmt.Sprintf("expected type of noun, got: %s for %s", p.curToken.Type.String(), p.curToken.Literal))
		return nil
	}
	ng.Noun = p.curToken
	p.nextToken()

	return ng
}

func (p *Parser) addError(err string) {
	p.errors = append(p.errors, err)
}
