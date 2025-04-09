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

type Uint32 struct {
	Token token.Token
}

func NewUint32(t token.Token) *Uint32 {
	return &Uint32{Token: t}
}

func (t *Uint32) Kind() tir.TirKind {
	return tir.Uint32
}

func (t *Uint32) String() string {
	return fmt.Sprintf("uint32(%s)", t.Token.Val)
}

func (t *Uint32) Id() int {
	return tir.GetId()
}

func (t *Uint32) constNode() {}
