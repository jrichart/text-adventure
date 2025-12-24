package parser

import (
	"text-adventure/ast"
	"text-adventure/lexer"
	"text-adventure/token"
)

type Parser struct {
	lexer     *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []*ast.ParseError
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:  l,
		errors: []*ast.ParseError{},
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) Errors() []error {
	return p.Errors()
}

func (p *Parser) ParseCommand() *ast.Command {
	cmd := &ast.Command{}

	if p.curToken.Type != token.VERB {
		p.errors = append(p.errors, &ast.ParseError{
			Message: "expected verb",
			Token:   p.curToken.Literal,
		})
	}
	cmd.VerbGroup = ast.VerbGroup{Verb: ast.Verb{Value: p.curToken.Literal}}
	p.nextToken()

	if p.curToken.Type == token.NOUN {
		noungroup := p.parseNounGroup()
		cmd.NounGroup = noungroup
	}

	if p.curToken.Type != token.EOF {
		p.errors = append(p.errors, &ast.ParseError{
			Message: "unexpected token after noun group",
			Token:   p.curToken.Literal,
		})
	}
	return cmd
}

func (p *Parser) parseNounGroup() ast.NounGroup {
	ng := ast.NounGroup{
		Adjectives: []*ast.Adjective{},
	}

	if p.curToken.Type != token.NOUN {
		p.errors = append(p.errors, &ast.ParseError{
			Message: "expected noun",
			Token:   p.curToken.Literal,
		})
	}
	ng.Noun = ast.Noun{Value: p.curToken.Literal}
	p.nextToken()

	return ng
}
