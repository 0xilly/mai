package tir

// Copyright (C) 2025 Anita Anderson

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

type TirKind int

const (
	StartKind TirKind = iota

	Int8
	Int16
	Int32
	Int64

	Uint8
	Uint16
	Uint32
	Uint64

	F32
	F64

	AddKind
	SubKind
	MulKind
	DivKind
	ModKind

	ShiftLeftKind
	ShiftRightKind

	DataKind

	EndKind
)

func (t TirKind) String() string {
	switch t {
	case StartKind:
		return "Start"
	case AddKind:
		return "Add"
	case SubKind:
		return "Sub"
	case MulKind:
		return "Mul"
	case DivKind:
		return "Div"
	case ModKind:
		return "Mod"
	case ShiftLeftKind:
		return "ShiftLeft"
	case ShiftRightKind:
		return "ShiftRight"
	case DataKind:
		return "Data"
	case EndKind:
		return "End"
	default:
		return "Unknown"

	}
}
