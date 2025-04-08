package parser

import (
	"fmt"
	"mai/token"
	"slices"
)

type (
	ParseError struct {
		File    string
		Line    int
		Column  int
		Message string
	}

	Parser struct {
		file    string
		lexer   *token.Lexer
		current token.Token
		next    token.Token

		Errors []ParseError
	}
)

func NewParser(file string) *Parser {
	p := &Parser{
		lexer: token.NewLexer(file),
	}
	//note(anita): set tokens
	p.advance()
	p.advance()
	return p
}

func (p *Parser) advance() {
	p.current = p.next
	p.next = p.lexer.NextToken()
}

func (p *Parser) match(types ...token.TokenType) bool {
	return slices.Contains(types, p.current.Type)
}

func (p *Parser) eatEol() {
	if p.match(token.T_EOL) {
		p.advance()
	}
}

func (p *Parser) consume(types ...token.TokenType) (token.Token, bool) {
	if p.match(types...) {
		cur := p.current
		p.advance()
		return cur, true
	}
	return p.current, false
}

func (p *Parser) report_bad_token(expect token.TokenType) {
	msg := fmt.Sprintf("Bad token [%s] found expected [%s].", p.current.Type.String(), expect.String())

	err := ParseError{
		File:    p.file,
		Line:    p.current.Line,
		Column:  p.current.Column,
		Message: msg,
	}

	p.Errors = append(p.Errors, err)

}

func (p *Parser) sync() {
	for p.current.Type != token.T_EOF {
		switch p.current.Type {
		case token.T_RPAREN, token.T_RBRACK, token.T_RBRACE, token.T_EOL, token.T_COMMA:
			p.advance()
			return
		}
		p.advance()
	}
}
