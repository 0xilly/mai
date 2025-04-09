package tir

// Copyright (C) 2025 Anita Anderson

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

var id int = 0

type Node interface {
	Id() int
	Kind() TirKind
	String() string
}

func GetId() int {
	localId := id
	id++
	return localId
}

type Add struct {
	id int
}
