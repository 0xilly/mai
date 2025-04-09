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

type Uint64 struct {
	Token token.Token
}

func NewUint64(t token.Token) *Uint64 {
	return &Uint64{Token: t}
}

func (t *Uint64) Kind() tir.TirKind {
	return tir.Uint64
}

func (t *Uint64) String() string {
	return fmt.Sprintf("uint64(%s)", t.Token.Val)
}

func (t *Uint64) Id() int {
	return tir.GetId()
}

func (t *Uint64) constNode() {}
