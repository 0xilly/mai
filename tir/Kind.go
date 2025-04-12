package tir

import "fmt"

// Copyright (C) 2025 Anita Anderson

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

type TirKind int

const (
	StartKind TirKind = iota

	Int8Kind
	Int16Kind
	Int32Kind
	Int64Kind

	Uint8Kind
	Uint16Kind
	Uint32Kind
	Uint64Kind

	F32Kind
	F64Kind

	AddKind
	SubKind
	MulKind
	DivKind
	ModKind

	EqKind // ==
	NeKind // !=
	LtKind // <
	LeKind // <=
	GtKind // >
	GeKind // >=

	// Bitwise Ops
	BitAndKind // &
	BitOrKind  // |
	BitXorKind // ^

	// Shift Ops
	ShiftLeftKind  // <<
	ShiftRightKind // >>

	// Boolean Logic
	LogicalAndKind // &&
	LogicalOrKind  // ||

	DataKind

	EndKind
)

func (t TirKind) String() string {
	switch t {
	case StartKind:
		return "start"
	case Int8Kind:
		return "int8"
	case Int16Kind:
		return "int16"
	case Int32Kind:
		return "int32"
	case Int64Kind:
		return "int64"
	case Uint8Kind:
		return "uint8"
	case Uint16Kind:
		return "uint16"
	case Uint32Kind:
		return "uint32"
	case Uint64Kind:
		return "uint64"
	case F32Kind:
		return "f32"
	case F64Kind:
		return "f64"
	case AddKind:
		return "add"
	case SubKind:
		return "sub"
	case MulKind:
		return "mul"
	case DivKind:
		return "div"
	case ModKind:
		return "mod"
	case EqKind:
		return "eq"
	case NeKind:
		return "ne"
	case LtKind:
		return "lt"
	case LeKind:
		return "le"
	case GtKind:
		return "gt"
	case GeKind:
		return "ge"
	case BitAndKind:
		return "bit_and"
	case BitOrKind:
		return "bit_or"
	case BitXorKind:
		return "bit_xor"
	case ShiftLeftKind:
		return "shift_left"
	case ShiftRightKind:
		return "shift_right"
	case LogicalAndKind:
		return "and"
	case LogicalOrKind:
		return "or"
	case DataKind:
		return "data"
	case EndKind:
		return "end"
	default:
		return fmt.Sprintf("%d", int(t))
	}
}
