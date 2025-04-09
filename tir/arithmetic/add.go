package arithmetic

// Copyright (C) 2025 Anita Anderson

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

import (
	"fmt"
	"mai/tir"
	"mai/tir/value"
	"strings"
)

type Add struct {
	In1 value.Value
	In2 value.Value
}

func NewAdd(in1, in2 value.Value) *Add {
	return &Add{
		In1: in1,
		In2: in2,
	}
}

func (a *Add) Id() int {
	return tir.GetId()
}

func (a *Add) Kind() tir.TirKind {
	return tir.AddKind
}

func (a *Add) String() string {
	str := a.Kind().String()
	return fmt.Sprintf("%s(%s, %s)", strings.ToLower(str), a.In1.String(), a.In2.String())
}
