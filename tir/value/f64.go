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

type F64 struct {
	Token token.Token
}

func NewF64(t token.Token) *F64 {
	return &F64{Token: t}
}

func (t *F64) Kind() tir.TirKind {
	return tir.F64
}

func (t *F64) String() string {
	return fmt.Sprintf("f64(%s)", t.Token.Val)
}

func (t *F64) Id() int {
	return tir.GetId()
}

func (t *F64) constNode() {}
