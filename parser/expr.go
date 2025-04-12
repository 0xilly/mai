package parser

// Copyright (C) 2025 Anita Anderson

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

import (
	"mai/ast"
	"mai/token"
)

// note(anta): expression = logical_or_expr
func (p *Parser) expr() ast.Expr {
	return p.logical_or()
}

func (p *Parser) logical_or() ast.Expr {
	expr := p.logical_and()

	for p.match(token.T_OROR) {
		op, _ := p.consume(token.T_OROR)
		right := p.logical_and()
		expr = &ast.BinaryExprNode{
			Left:  expr,
			Op:    op,
			Right: right,
		}
	}

	return expr
}

// note(anta): logical_and parses logical AND expressions ("&&").
func (p *Parser) logical_and() ast.Expr {
	expr := p.bitwise_or()

	for p.match(token.T_ANDAND) {
		op, _ := p.consume(token.T_ANDAND)
		right := p.bitwise_or()
		expr = &ast.BinaryExprNode{
			Left:  expr,
			Op:    op,
			Right: right,
		}
	}

	return expr
}

// note(anta): bitwise_or parses bitwise OR expressions ("|").
func (p *Parser) bitwise_or() ast.Expr {
	expr := p.bitwise_xor()

	for p.match(token.T_BIT_OR) {
		op, _ := p.consume(token.T_BIT_OR)
		right := p.bitwise_xor()
		expr = &ast.BinaryExprNode{
			Left:  expr,
			Op:    op,
			Right: right,
		}
	}

	return expr
}

// note(anta): bitwise_xor parses bitwise XOR expressions ("^").
func (p *Parser) bitwise_xor() ast.Expr {
	expr := p.bitwise_and()

	for p.match(token.T_CARET) {
		op, _ := p.consume(token.T_CARET)
		right := p.bitwise_and()
		expr = &ast.BinaryExprNode{
			Left:  expr,
			Op:    op,
			Right: right,
		}
	}

	return expr
}

// note(anta): bitwise_and parses bitwise AND expressions ("&").
func (p *Parser) bitwise_and() ast.Expr {
	expr := p.shift()

	for p.match(token.T_BIT_AND) {
		op, _ := p.consume(token.T_BIT_AND)
		right := p.shift()
		expr = &ast.BinaryExprNode{
			Left:  expr,
			Op:    op,
			Right: right,
		}
	}

	return expr
}

// note(anta): shift parses shift expressions ("<<" and ">>").
func (p *Parser) shift() ast.Expr {
	expr := p.additive()

	for p.match(token.T_SHIFT_LEFT, token.T_SHIFT_RIGHT) {
		op, _ := p.consume(token.T_SHIFT_LEFT, token.T_SHIFT_RIGHT)
		right := p.additive()
		expr = &ast.BinaryExprNode{
			Left:  expr,
			Op:    op,
			Right: right,
		}
	}

	return expr
}

// note(anta): additive parses addition and subtraction expressions ("+", "-").
func (p *Parser) additive() ast.Expr {
	expr := p.multiplicative()

	for p.match(token.T_PLUS, token.T_MINUS) {
		op, _ := p.consume(token.T_PLUS, token.T_MINUS)
		right := p.multiplicative()
		expr = &ast.BinaryExprNode{
			Left:  expr,
			Op:    op,
			Right: right,
		}
	}

	return expr
}

// note(anta): multiplicative parses multiplication, division, and modulo expressions ("*", "/", "%").
func (p *Parser) multiplicative() ast.Expr {
	expr := p.unary()

	for p.match(token.T_STAR, token.T_SLASH, token.T_PERCENT) {
		op, _ := p.consume(token.T_STAR, token.T_SLASH, token.T_PERCENT)
		right := p.unary()
		expr = &ast.BinaryExprNode{
			Left:  expr,
			Op:    op,
			Right: right,
		}
	}

	return expr
}

// note(anta): unary parses unary expressions ("-", "!", "~").
func (p *Parser) unary() ast.Expr {
	if p.match(token.T_MINUS, token.T_BANG, token.T_BIT_NOT) {
		op, _ := p.consume(token.T_MINUS, token.T_BANG, token.T_BIT_NOT)
		right := p.unary()
		return &ast.UnaryExprNode{
			Op:    op,
			Right: right,
		}
	}

	return p.primary()
}

// note(anta): primary parses primary expressions and applies postfix operations.
func (p *Parser) primary() ast.Expr {
	var expr ast.Expr

	switch {
	case p.match(token.T_LIT_NUMBER,
		token.T_LIT_RAW_STRING,
		token.T_LIT_STRING,
		token.T_LIT_RUNE,
		token.T_TRUE,
		token.T_FALSE):
		expr = p.literal()

	case p.match(token.T_IDENTIFIER):
		expr = p.identifier_expr()

	case p.match(token.T_LPAREN):
		expr = p.group()

	default:
		p.report_generic_error("Bad primary")
		p.sync()
		return nil
	}

	// Apply any postfix operations to the primary expression
	return p.postfix(expr)
}

// note(anta): postfix parses zero or more postfix operations on the given expression.
func (p *Parser) postfix(expr ast.Expr) ast.Expr {
	for {
		switch {
		case p.match(token.T_DOT):
			expr = p.fieldAccess(expr)
		case p.match(token.T_LPAREN):
			expr = p.call(expr)
		case p.match(token.T_LBRACK):
			expr = p.index(expr)
		default:
			return expr
		}
	}
}

// note(anta): fieldAccess handles expressions like expr.field
func (p *Parser) fieldAccess(parent ast.Expr) ast.Expr {
	dot, ok := p.consume(token.T_DOT)
	if !ok {
		p.report_bad_token(token.T_DOT)
		p.sync()
		return parent
	}

	fieldIdent, ok := p.consume(token.T_IDENTIFIER)
	if !ok {
		p.report_bad_token(token.T_IDENTIFIER)
		p.sync()
		return parent
	}

	return &ast.FieldAccessNode{
		Dot:    &dot,
		Parent: parent,
		Ident:  &ast.LiteralNode{Ident: fieldIdent},
	}
}

func (p *Parser) call(callee ast.Expr) ast.Expr {
	openParen, ok := p.consume(token.T_LPAREN)
	if !ok {
		p.report_bad_token(token.T_LPAREN)
		p.sync()
		return callee
	}

	var args []ast.Expr
	if !p.match(token.T_RPAREN) {
		args = append(args, p.expr())
		for p.match(token.T_COMMA) {
			_, ok := p.consume(token.T_COMMA)
			if !ok {
				p.report_bad_token(token.T_COMMA)
				p.sync()
				break
			}
			args = append(args, p.expr())
		}
	}

	closeParen, ok := p.consume(token.T_RPAREN)
	if !ok {
		p.report_bad_token(token.T_RPAREN)
		p.sync()
		return callee
	}

	return &ast.CallNode{
		Fn:    callee,
		Open:  openParen,
		Args:  args,
		Close: closeParen,
	}
}

// note(anta): index handles indexing expressions like expr[index]
func (p *Parser) index(target ast.Expr) ast.Expr {
	openBracket, ok := p.consume(token.T_LBRACK)
	if !ok {
		p.report_bad_token(token.T_LBRACK)
		p.sync()
		return target
	}

	indexExpr := p.expr()

	closeBracket, ok := p.consume(token.T_RBRACK)
	if !ok {
		p.report_bad_token(token.T_RBRACK)
		p.sync()
	}

	return &ast.ArrayIndexNode{
		Open:   openBracket,
		Target: target,
		Index:  indexExpr,
		Close:  closeBracket,
	}
}

// note(anta): literal parses a literal expression.
func (p *Parser) literal() ast.Expr {
	lit, ok := p.consume(
		token.T_LIT_NUMBER,
		token.T_LIT_RAW_STRING,
		token.T_LIT_STRING,
		token.T_LIT_RUNE,
		token.T_TRUE,
		token.T_FALSE,
	)

	if !ok {
		p.report_bad_literal()
		p.sync()
		return nil
	}
	return &ast.LiteralNode{Ident: lit}
}

// note(anta): identifier_expr parses an identifier expression.
func (p *Parser) identifier_expr() ast.Expr {
	ident, ok := p.consume(token.T_IDENTIFIER)
	if !ok {
		p.report_bad_token(token.T_IDENTIFIER)
		p.sync()
		return nil
	}

	return &ast.LiteralNode{Ident: ident}
}

// note(anta): group parses a grouped expression "(expr)".
func (p *Parser) group() ast.Expr {
	openParen, _ := p.consume(token.T_LPAREN)

	expr := p.expr()

	closeParen, ok := p.consume(token.T_RPAREN)
	if !ok {
		p.report_bad_token(token.T_RPAREN)
		p.sync()
	}

	return &ast.GroupExprNode{
		Open:       openParen,
		Expression: expr,
		Close:      closeParen,
	}
}
