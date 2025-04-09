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

type Uint16 struct {
	Token token.Token
}

func NewUint16(t token.Token) *Uint16 {
	return &Uint16{Token: t}
}

func (t *Uint16) Kind() tir.TirKind {
	return tir.Uint16
}

func (t *Uint16) String() string {
	return fmt.Sprintf("uint16(%s)", t.Token.Val)
}

func (t *Uint16) Id() int {
	return tir.GetId()
}

func (t *Uint16) constNode() {}
