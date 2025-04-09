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

type F32 struct {
	Token token.Token
}

func NewF32(t token.Token) *F32 {
	return &F32{Token: t}
}

func (t *F32) Kind() tir.TirKind {
	return tir.F32
}

func (t *F32) String() string {
	return fmt.Sprintf("f32(%s)", t.Token.Val)
}

func (t *F32) Id() int {
	return tir.GetId()
}

func (t *F32) constNode() {}
