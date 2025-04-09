package value

import (
	"fmt"
	"mai/tir"
	"mai/token"
)

type Uint8 struct {
	Token token.Token
}

func NewUint8(t token.Token) *Uint8 {
	return &Uint8{Token: t}
}

func (t *Uint8) Kind() tir.TirKind {
	return tir.Uint8
}

func (t *Uint8) String() string {
	return fmt.Sprintf("uint8(%s)", t.Token.Val)
}

func (t *Uint8) Id() int {
	return tir.GetId()
}

func (t *Uint8) constNode() {}
