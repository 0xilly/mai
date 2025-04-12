package tir

// Copyright (C) 2025 Anita Anderson

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

import (
	"fmt"
	"mai/token"
)

type (
	Int8 struct {
		ID    int
		Token token.Token
	}

	Uint8 struct {
		ID    int
		Token token.Token
	}

	Int16 struct {
		ID    int
		Token token.Token
	}

	Uint16 struct {
		ID    int
		Token token.Token
	}

	Int32 struct {
		ID    int
		Token token.Token
	}

	Uint32 struct {
		ID    int
		Token token.Token
	}

	Int64 struct {
		ID    int
		Token token.Token
	}

	Uint64 struct {
		ID    int
		Token token.Token
	}

	F32 struct {
		ID    int
		Token token.Token
	}

	F64 struct {
		ID    int
		Token token.Token
	}
)

func NewInt8(t token.Token) *Int8 {
	return &Int8{
		ID:    getId(),
		Token: t,
	}
}

func (t *Int8) Kind() TirKind  { return Int8Kind }
func (t *Int8) String() string { return fmt.Sprintf("int8(%s)", t.Token.Val) }
func (t *Int8) Id() int        { return t.ID }
func (t *Int8) valueNode()     {}

func NewUint8(t token.Token) *Uint8 {
	return &Uint8{
		ID:    getId(),
		Token: t,
	}
}

func (t *Uint8) Kind() TirKind  { return Uint8Kind }
func (t *Uint8) String() string { return fmt.Sprintf("uint8(%s)", t.Token.Val) }
func (t *Uint8) Id() int        { return t.ID }
func (t *Uint8) valueNode()     {}

func NewInt16(t token.Token) *Int16 {
	return &Int16{
		ID:    getId(),
		Token: t,
	}
}

func (t *Int16) Kind() TirKind  { return Int16Kind }
func (t *Int16) String() string { return fmt.Sprintf("int16(%s)", t.Token.Val) }
func (t *Int16) Id() int        { return t.ID }
func (t *Int16) valueNode()     {}

func NewUint16(t token.Token) *Uint16 {
	return &Uint16{
		ID:    getId(),
		Token: t,
	}
}

func (t *Uint16) Kind() TirKind  { return Uint16Kind }
func (t *Uint16) String() string { return fmt.Sprintf("uint16(%s)", t.Token.Val) }
func (t *Uint16) Id() int        { return t.ID }
func (t *Uint16) valueNode()     {}

func NewInt32(t token.Token) *Int32 {
	return &Int32{
		ID:    getId(),
		Token: t,
	}
}

func (t *Int32) Kind() TirKind  { return Int32Kind }
func (t *Int32) String() string { return fmt.Sprintf("int32(%s)", t.Token.Val) }
func (t *Int32) Id() int        { return t.ID }
func (t *Int32) valueNode()     {}

func NewUint32(t token.Token) *Uint32 {
	return &Uint32{
		ID:    getId(),
		Token: t,
	}
}

func (t *Uint32) Kind() TirKind  { return Uint32Kind }
func (t *Uint32) String() string { return fmt.Sprintf("uint32(%s)", t.Token.Val) }
func (t *Uint32) Id() int        { return t.ID }
func (t *Uint32) valueNode()     {}

func NewInt64(t token.Token) *Int64 {
	return &Int64{
		ID:    getId(),
		Token: t,
	}
}

func (t *Int64) Kind() TirKind  { return Int64Kind }
func (t *Int64) String() string { return fmt.Sprintf("int64(%s)", t.Token.Val) }
func (t *Int64) Id() int        { return t.ID }
func (t *Int64) valueNode()     {}

func NewUint64(t token.Token) *Uint64 {
	return &Uint64{
		ID:    getId(),
		Token: t,
	}
}

func (t *Uint64) Kind() TirKind  { return Uint64Kind }
func (t *Uint64) String() string { return fmt.Sprintf("uint64(%s)", t.Token.Val) }
func (t *Uint64) Id() int        { return t.ID }
func (t *Uint64) valueNode()     {}

func NewF32(t token.Token) *F32 {
	return &F32{
		ID:    id,
		Token: t,
	}
}

func (t *F32) Kind() TirKind  { return F32Kind }
func (t *F32) String() string { return fmt.Sprintf("f32(%s)", t.Token.Val) }
func (t *F32) Id() int        { return t.ID }
func (t *F32) valueNode()     {}

func NewF64(t token.Token) *F64 {
	return &F64{
		ID:    getId(),
		Token: t}
}
func (t *F64) Kind() TirKind  { return F64Kind }
func (t *F64) String() string { return fmt.Sprintf("f64(%s)", t.Token.Val) }
func (t *F64) Id() int        { return t.ID }
func (t *F64) valueNode()     {}
