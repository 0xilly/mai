package token_test

// Copyright (C) 2025 Anita Anderson

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

import (
	"mai/token"
	"testing"
)

func TestLexerFunctionDefinition(t *testing.T) {
	input := `add :: fn(a:i8, b:i8):i8 -> ret a + b`
	expectedTokens := []token.Token{
		{Type: token.T_IDENTIFIER, Val: "add", Line: 1, Column: 1},
		{Type: token.T_COLON_COLON, Val: "::", Line: 1, Column: 5},
		{Type: token.T_FN, Val: "fn", Line: 1, Column: 8},
		{Type: token.T_LPAREN, Val: "(", Line: 1, Column: 10},
		{Type: token.T_IDENTIFIER, Val: "a", Line: 1, Column: 11},
		{Type: token.T_COLON, Val: ":", Line: 1, Column: 12},
		{Type: token.T_I8, Val: "i8", Line: 1, Column: 13},
		{Type: token.T_COMMA, Val: ",", Line: 1, Column: 15},
		{Type: token.T_IDENTIFIER, Val: "b", Line: 1, Column: 17},
		{Type: token.T_COLON, Val: ":", Line: 1, Column: 18},
		{Type: token.T_I8, Val: "i8", Line: 1, Column: 19},
		{Type: token.T_RPAREN, Val: ")", Line: 1, Column: 21},
		{Type: token.T_COLON, Val: ":", Line: 1, Column: 22},
		{Type: token.T_I8, Val: "i8", Line: 1, Column: 23},
		{Type: token.T_ARROW, Val: "->", Line: 1, Column: 26},
		{Type: token.T_RET, Val: "ret", Line: 1, Column: 29},
		{Type: token.T_IDENTIFIER, Val: "a", Line: 1, Column: 33},
		{Type: token.T_PLUS, Val: "+", Line: 1, Column: 35},
		{Type: token.T_IDENTIFIER, Val: "b", Line: 1, Column: 37},
		{Type: token.T_EOF, Val: "", Line: 1, Column: 38},
	}

	lexer := token.NewLexer(input)
	tokens := []token.Token{}

	for {
		tok := lexer.NextToken()
		tokens = append(tokens, tok)
		if tok.Type == token.T_EOF {
			break
		}
	}

	if len(tokens) != len(expectedTokens) {
		t.Fatalf("Token count mismatch. Got %d tokens, expected %d", len(tokens), len(expectedTokens))
	}

	for i, tok := range tokens {
		if tok.Type != expectedTokens[i].Type || tok.Val != expectedTokens[i].Val {
			t.Errorf("Token %d mismatch. Got %+v, expected %+v", i, tok, expectedTokens[i])
		}
	}
}
