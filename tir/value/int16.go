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

type Int16 struct {
	Token token.Token
}

func NewInt16(t token.Token) *Int16 {
	return &Int16{Token: t}
}

func (t *Int16) Kind() tir.TirKind {
	return tir.Int16
}

func (t *Int16) String() string {
	return fmt.Sprintf("int16(%s)", t.Token.Val)
}

func (t *Int16) Id() int {
	return tir.GetId()
}

func (t *Int16) constNode() {}
