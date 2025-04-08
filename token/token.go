package token

// Copyright (C) 2025 Anita Anderson

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

//note(anita): based heavily off of go's template lexer https://github.com/golang/go/blob/master/src/text/template/parse/lex.go

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

type TokenType int

const (
	// Special
	T_EOF TokenType = iota
	T_ERROR

	// Comments
	T_COMMENT
	T_DOC_COMMENT

	T_EOL

	// Identifiers & Keywords
	T_IDENTIFIER

	// Some example keywords
	T_SELF
	T_IF
	T_ELSE
	T_ELIF
	T_RET
	T_MUT
	T_FN
	T_DATA
	T_OWN
	T_WITH
	T_PKG
	T_USE
	T_MATCH
	T_BREAK
	T_CONTINUE
	T_ASSERT
	T_EX
	T_LINK
	T_BIND
	T_INTERFACE
	T_ENUM
	T_UNION
	T_FOR
	T_DEFER
	T_ALIAS

	PRIMITIVE_START
	T_U8
	T_U16
	T_U32
	T_U64
	T_USIZE
	T_I8
	T_I16
	T_I32
	T_I64
	T_ISIZE
	T_F32
	T_F64
	T_TRUE
	T_FALSE
	T_STRING
	T_RUNE
	T_BOOL
	T_VOID
	PRIMITIVE_END

	// Literals
	T_LIT_NUMBER
	T_LIT_STRING
	T_LIT_RAW_STRING
	T_LIT_RUNE

	// Operators / punctuation
	T_EQEQ         // ==
	T_NEQ          // !=
	T_LTEQ         // <=
	T_GTEQ         // >=
	T_ANDAND       // &&
	T_OROR         // ||
	T_COLON_COLON  // ::
	T_ARROW        // ->
	T_DOTDOT       // ..
	T_ELLIPSIS     // ...
	T_L_ARROW      // <-
	T_COLON_ASSIGN // :=
	T_PLUS_ASSIGN  // +=
	T_MINUS_ASSIGN // -=
	T_MUL_ASSIGN   // *=
	T_DIV_ASSIGN   // /=
	T_MOD_ASSIGN   // %=

	// single-char
	T_ASSIGN       // =
	T_LT           // <
	T_GT           // >
	T_PLUS         // +
	T_MINUS        // -
	T_STAR         // *
	T_SLASH        // /
	T_PERCENT      // %
	T_BANG         // !
	T_COLON        // :
	T_SEMI         // ;
	T_DOT          // .
	T_COMMA        // ,
	T_LPAREN       // (
	T_RPAREN       // )
	T_LBRACK       // [
	T_RBRACK       // ]
	T_LBRACE       // {
	T_RBRACE       // }
	T_CARET        // ^
	T_AT           // @
	T_DOLLARSIGN   // $
	T_QUESTIONMARK // ?
	T_BIT_AND      // &
	T_BIT_OR       // |
	T_BIT_NOT      // ~
	T_SHIFT_LEFT   // <<
	T_SHIFT_RIGHT  // >>
)

var tokenNames = map[TokenType]string{
	T_EOF:         "EOF",
	T_ERROR:       "ERROR",
	T_COMMENT:     "COMMENT",
	T_DOC_COMMENT: "DOC_COMMENT",

	T_IDENTIFIER: "IDENTIFIER",
	T_SELF:       "SELF",
	T_IF:         "IF",
	T_ELSE:       "ELSE",
	T_ELIF:       "ELIF",
	T_RET:        "RET",
	T_MUT:        "MUT",
	T_FN:         "FN",
	T_DATA:       "DATA",
	T_OWN:        "OWN",
	T_WITH:       "WITH",
	T_PKG:        "PKG",
	T_USE:        "USE",
	T_MATCH:      "MATCH",
	T_BREAK:      "BREAK",
	T_CONTINUE:   "CONTINUE",

	T_LIT_NUMBER:     "NUMBER",
	T_LIT_STRING:     "STRING",
	T_LIT_RAW_STRING: "RAW_STRING",
	T_LIT_RUNE:       "RUNE",

	T_ASSERT:    "ASSERT",
	T_EX:        "EX",
	T_LINK:      "LINK",
	T_BIND:      "BIND",
	T_INTERFACE: "INTERFACE",
	T_ENUM:      "ENUM",
	T_UNION:     "UNION",
	T_FOR:       "FOR",
	T_ALIAS:     "ALIAS",

	T_U8:     "U8",
	T_U16:    "U16",
	T_U32:    "U32",
	T_U64:    "U64",
	T_USIZE:  "USIZE",
	T_I8:     "I8",
	T_I16:    "I16",
	T_I32:    "I32",
	T_I64:    "I64",
	T_ISIZE:  "ISIZE",
	T_F32:    "F32",
	T_F64:    "F64",
	T_TRUE:   "TRUE",
	T_FALSE:  "FALSE",
	T_STRING: "STRING",
	T_RUNE:   "RUNE",
	T_BOOL:   "BOOL",
	T_VOID:   "VOID",

	T_EQEQ:         "==",
	T_NEQ:          "!=",
	T_LTEQ:         "<=",
	T_GTEQ:         ">=",
	T_ANDAND:       "&&",
	T_OROR:         "||",
	T_COLON_COLON:  "::",
	T_ARROW:        "->",
	T_DOTDOT:       "..",
	T_ELLIPSIS:     "...",
	T_L_ARROW:      "<-",
	T_COLON_ASSIGN: ":=",
	T_PLUS_ASSIGN:  "+=",
	T_MINUS_ASSIGN: "-=",
	T_MUL_ASSIGN:   "*=",
	T_DIV_ASSIGN:   "/=",
	T_MOD_ASSIGN:   "%=",

	T_ASSIGN:       "=",
	T_LT:           "<",
	T_GT:           ">",
	T_PLUS:         "+",
	T_MINUS:        "-",
	T_STAR:         "*",
	T_SLASH:        "/",
	T_PERCENT:      "%",
	T_BANG:         "!",
	T_COLON:        ":",
	T_SEMI:         ";",
	T_DOT:          ".",
	T_COMMA:        ",",
	T_LPAREN:       "(",
	T_RPAREN:       ")",
	T_LBRACK:       "[",
	T_RBRACK:       "]",
	T_LBRACE:       "{",
	T_RBRACE:       "}",
	T_CARET:        "^",
	T_AT:           "@",
	T_DOLLARSIGN:   "$",
	T_QUESTIONMARK: "?",
	T_BIT_AND:      "&",
	T_BIT_OR:       "|",
	T_BIT_NOT:      "!",
	T_SHIFT_LEFT:   "<<",
	T_SHIFT_RIGHT:  ">>",
}

func (tn TokenType) String() string {
	if s, ok := tokenNames[tn]; ok {
		return s
	}
	return fmt.Sprintf("%s", tn)
}

func (tn TokenType) IsPrimitive() bool { return tn >= PRIMITIVE_START && tn <= PRIMITIVE_END }

type Token struct {
	Type   TokenType
	Val    string
	Line   int
	Column int
	Start  int
	Pos    int
	Width  int
}

func (t Token) String() string {
	return fmt.Sprintf("[%d:%d] %s(%q)", t.Line, t.Column, t.Type, t.Val)
}

var keywords = map[string]TokenType{
	"self":      T_SELF,
	"if":        T_IF,
	"else":      T_ELSE,
	"elif":      T_ELIF,
	"ret":       T_RET,
	"mut":       T_MUT,
	"fn":        T_FN,
	"data":      T_DATA,
	"own":       T_OWN,
	"with":      T_WITH,
	"pkg":       T_PKG,
	"use":       T_USE,
	"match":     T_MATCH,
	"break":     T_BREAK,
	"continue":  T_CONTINUE,
	"assert":    T_ASSERT,
	"ex":        T_EX,
	"link":      T_LINK,
	"bind":      T_BIND,
	"interface": T_INTERFACE,
	"enum":      T_ENUM,
	"union":     T_UNION,
	"for":       T_FOR,

	"u8":     T_U8,
	"u16":    T_U16,
	"u32":    T_U32,
	"u64":    T_U64,
	"usize":  T_USIZE,
	"i8":     T_I8,
	"i16":    T_I16,
	"i32":    T_I32,
	"i64":    T_I64,
	"isize":  T_ISIZE,
	"f32":    T_F32,
	"f64":    T_F64,
	"true":   T_TRUE,
	"false":  T_FALSE,
	"string": T_STRING,
	"rune":   T_RUNE,
	"bool":   T_BOOL,
	"void":   T_VOID,
	"alias":  T_ALIAS,
}

// Lexer
type Lexer struct {
	input  string
	start  int
	pos    int
	width  int
	tokens chan Token

	line   int
	column int
}

const eof = -1

type stateFn func(*Lexer) stateFn

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input:  input,
		tokens: make(chan Token),
		line:   1,
		column: 0,
	}
	go l.run()
	return l
}

func (l *Lexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	close(l.tokens)
}

func (l *Lexer) NextToken() Token {
	return <-l.tokens
}

func (l *Lexer) emit(tn TokenType) {
	val := l.input[l.start:l.pos]
	l.tokens <- Token{
		Type:   tn,
		Val:    val,
		Line:   l.line,
		Column: l.column - (l.pos - l.start) + 1,
		Start:  l.start,
		Pos:    l.pos,
		Width:  l.width,
	}
	l.start = l.pos
}

func (l *Lexer) ignore() {
	l.start = l.pos
}

func (l *Lexer) backup() {
	l.pos -= l.width
}

func (l *Lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

func (l *Lexer) errorf(format string, args ...interface{}) stateFn {
	msg := fmt.Sprintf(format, args...)
	l.tokens <- Token{
		Type:   T_ERROR,
		Val:    msg,
		Line:   l.line,
		Column: l.column,
	}
	return nil
}

func (l *Lexer) next() rune {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += w
	l.width = w
	if r == '\n' {
		l.line++
		l.column = 0
	} else {
		l.column++
	}
	return r
}

func lexText(l *Lexer) stateFn {
	for {
		r := l.next()

		switch r {
		case eof:
			l.emit(T_EOF)
			return nil

		// Whitespace
		case ' ', '\t', '\r':
			l.ignore()
		case '\n':
			l.emit(T_EOL)

		case '/':
			if l.peek() == '=' {
				l.next()
				l.emit(T_DIV_ASSIGN)
			} else if l.peek() == '/' {
				l.next()
				if l.peek() == '/' {
					l.next()
					return lexDocLineComment
				} else {
					return lexLineComment
				}
			} else if l.peek() == '*' {
				l.next()
				if l.peek() == '*' {
					l.next()
					return lexDocMultiLineComment
				} else {
					return lexMultiLineComment
				}
			} else {
				l.emit(T_SLASH)
			}

		case '=':
			if l.peek() == '=' {
				l.next()
				l.emit(T_EQEQ)
			} else {
				l.emit(T_ASSIGN)
			}

		case '!':
			if l.peek() == '=' {
				l.next()
				l.emit(T_NEQ)
			} else {
				l.emit(T_BANG)
			}

		case '<':
			switch l.peek() {
			case '=':
				l.next()
				l.emit(T_LTEQ)
			case '-':
				l.next()
				l.emit(T_L_ARROW)
			case '<':
				l.next()
				l.emit(T_SHIFT_LEFT)
			default:
				l.emit(T_LT)
			}

		case '>':
			switch {
			case l.peek() == '=':
				l.next()
				l.emit(T_GTEQ)
			case l.peek() == '>':
				l.next()
				l.emit(T_SHIFT_RIGHT)
			default:
				l.emit(T_GT)
			}

		case '&':
			switch {
			case l.peek() == '&':
				l.next()
				l.emit(T_ANDAND)
			default:
				l.emit(T_BIT_AND)
			}

		case '|':
			switch {
			case l.peek() == '|':
				l.next()
				l.emit(T_OROR)
			default:
				l.emit(T_BIT_OR)
			}

		case ':':
			if l.peek() == '=' {
				l.next()
				l.emit(T_COLON_ASSIGN)
			} else if l.peek() == ':' {
				l.next()
				l.emit(T_COLON_COLON)
			} else {
				l.emit(T_COLON)
			}

		case '-':
			switch l.peek() {
			case '=':
				l.next()
				l.emit(T_MINUS_ASSIGN)
			case '>':
				l.next()
				l.emit(T_ARROW)
			default:
				l.emit(T_MINUS)
			}

		case '+':
			if l.peek() == '=' {
				l.next()
				l.emit(T_PLUS_ASSIGN)
			} else {
				l.emit(T_PLUS)
			}

		case '*':
			if l.peek() == '=' {
				l.next()
				l.emit(T_MUL_ASSIGN)
			} else {
				l.emit(T_STAR)
			}

		case '%':
			if l.peek() == '=' {
				l.next()
				l.emit(T_MOD_ASSIGN)
			} else {
				l.emit(T_PERCENT)
			}

		case '.':
			if l.peek() == '.' {
				l.next()
				if l.peek() == '.' {
					l.next()
					l.emit(T_ELLIPSIS)
				} else {
					l.emit(T_DOTDOT)
				}
			} else {
				l.emit(T_DOT)
			}

		case ';':
			l.emit(T_SEMI)
		case ',':
			l.emit(T_COMMA)
		case '(':
			l.emit(T_LPAREN)
		case ')':
			l.emit(T_RPAREN)
		case '[':
			l.emit(T_LBRACK)
		case ']':
			l.emit(T_RBRACK)
		case '{':
			l.emit(T_LBRACE)
		case '}':
			l.emit(T_RBRACE)
		case '^':
			l.emit(T_CARET)

		case '"':
			return lexString

		case 'r':
			if l.peek() == '"' {
				l.next()
				return lexRawString
			}
			l.backup()
			return lexIdentifier

		case '\'':
			return lexRune
		case '@':
			l.emit(T_AT)
		case '$':
			l.emit(T_DOLLARSIGN)
		case '?':
			l.emit(T_QUESTIONMARK)
		case '~':
			l.emit(T_BIT_NOT)

		default:
			if isAlpha(r) {
				l.backup()
				return lexIdentifier
			} else if isDigit(r) {
				l.backup()
				return lexNumber
			} else {
				return l.errorf("unrecognized character: %q", r)
			}
		}
	}
}

// note(anita): Doc comment lexers
func lexDocLineComment(l *Lexer) stateFn {
	for {
		r := l.next()
		if r == '\n' || r == eof {
			l.backup()
			l.emit(T_DOC_COMMENT)
			return lexText
		}
	}
}

func lexDocMultiLineComment(l *Lexer) stateFn {
	for {
		r := l.next()
		if r == eof {
			return l.errorf("unterminated doc multi-line comment")
		}
		if r == '*' && l.peek() == '/' {
			l.next()
			l.emit(T_DOC_COMMENT)
			return lexText
		}
	}
}

func lexLineComment(l *Lexer) stateFn {
	for {
		r := l.next()
		if r == '\n' || r == eof {
			l.backup()
			l.emit(T_COMMENT)
			return lexText
		}
	}
}

func lexMultiLineComment(l *Lexer) stateFn {
	for {
		r := l.next()
		if r == eof {
			return l.errorf("unterminated multi-line comment")
		}
		if r == '*' && l.peek() == '/' {
			l.next()
			l.emit(T_COMMENT)
			return lexText
		}
	}
}

// note(anita): Identifiers & Keywords
func lexIdentifier(l *Lexer) stateFn {
	for {
		r := l.next()
		if !isAlphaNumeric(r) {
			l.backup()
			break
		}
	}
	val := l.input[l.start:l.pos]
	if kw, ok := keywords[val]; ok {
		l.emit(kw)
	} else {
		l.emit(T_IDENTIFIER)
	}
	return lexText
}

func lexNumber(l *Lexer) stateFn {
	r := l.next()
	if r == '0' {
		switch l.peek() {
		case 'x', 'X':
			l.next()
			for isHexDigit(l.peek()) {
				l.next()
			}
		case 'b', 'B':
			l.next()
			for isBinDigit(l.peek()) {
				l.next()
			}
		case 'o', 'O':
			l.next()
			for isOctDigit(l.peek()) {
				l.next()
			}
		default:
			for isDecimalOrUnderscore(l.peek()) {
				l.next()
			}
		}
	} else {
		for isDecimalOrUnderscore(l.peek()) {
			l.next()
		}
	}
	l.emit(T_LIT_NUMBER)
	return lexText
}

func lexString(l *Lexer) stateFn {
	for {
		r := l.next()
		switch r {
		case eof:
			return l.errorf("unterminated string literal")
		case '\\':
			l.next()
		case '"':
			l.emit(T_LIT_STRING)
			return lexText
		}
	}
}

func lexRawString(l *Lexer) stateFn {
	for {
		r := l.next()
		if r == eof {
			return l.errorf("unterminated raw string literal")
		}
		if r == '"' {
			l.emit(T_LIT_RAW_STRING)
			return lexText
		}
	}
}

func lexRune(l *Lexer) stateFn {
	for {
		r := l.next()
		switch r {
		case eof:
			return l.errorf("unterminated rune literal")
		case '\\':
			l.next()
		case '\'':
			l.emit(T_LIT_RUNE)
			return lexText
		}
	}
}

func isAlpha(r rune) bool {
	return r == '_' || unicode.IsLetter(r)
}

func isAlphaNumeric(r rune) bool {
	return isAlpha(r) || unicode.IsDigit(r)
}

func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func isHexDigit(r rune) bool {
	return (r >= '0' && r <= '9') ||
		(r >= 'a' && r <= 'f') ||
		(r >= 'A' && r <= 'F')
}

func isOctDigit(r rune) bool {
	return r >= '0' && r <= '7'
}

func isBinDigit(r rune) bool {
	return r == '0' || r == '1'
}

func isDecimalOrUnderscore(r rune) bool {
	return unicode.IsDigit(r) || r == '_'
}
