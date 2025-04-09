package value

// Copyright (C) 2025 Anita Anderson

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

import (
	"fmt"
	"mai/tir"
	"mai/token"
)

type Int32 struct {
	Token token.Token
}

func NewInt32(t token.Token) *Int32 {
	return &Int32{Token: t}
}

func (t *Int32) Kind() tir.TirKind {
	return tir.Int32
}

func (t *Int32) String() string {
	return fmt.Sprintf("int32(%s)", t.Token.Val)
}

func (t *Int32) Id() int {
	return tir.GetId()
}

func (t *Int32) constNode() {}
