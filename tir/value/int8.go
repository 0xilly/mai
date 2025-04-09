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

type Int8 struct {
	Token token.Token
}

func NewInt8(t token.Token) *Int8 {
	return &Int8{Token: t}
}

func (t *Int8) Kind() tir.TirKind {
	return tir.Int8
}

func (t *Int8) String() string {
	return fmt.Sprintf("int8(%s)", t.Token.Val)
}

func (t *Int8) Id() int {
	return tir.GetId()
}

func (t *Int8) constNode() {}
