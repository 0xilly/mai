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

type Int64 struct {
	Token token.Token
}

func NewInt64(t token.Token) *Int64 {
	return &Int64{Token: t}
}

func (t *Int64) Kind() tir.TirKind {
	return tir.Int64
}

func (t *Int64) String() string {
	return fmt.Sprintf("int64(%s)", t.Token.Val)
}

func (t *Int64) Id() int {
	return tir.GetId()
}

func (t *Int64) constNode() {}
