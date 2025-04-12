package tir

import "fmt"

var id int = 0

func getId() int {
	localId := id
	id++
	return localId
}

type (
	Node interface {
		Id() int
		Kind() TirKind
		String() string
	}

	Value interface {
		Node
		valueNode()
	}
)

type BiNode struct {
	Node
	Id    int
	Left  Value
	kind  TirKind
	Right Value
}

func NewBiNode(left Value, kind TirKind, right Value) *BiNode {
	return &BiNode{
		Id:    getId(),
		Left:  left,
		kind:  kind,
		Right: right,
	}
}

func (b *BiNode) String() string {
	return fmt.Sprintf("%s(%s, %s)", b.Kind().String(), b.Left.String(), b.Right.String())
}
func (b *BiNode) Kind() TirKind { return b.kind }
