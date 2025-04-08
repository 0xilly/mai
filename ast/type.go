package ast

import "mai/token"

// types
type (
	BaseTypeNode struct {
		Typ token.Token
	}

	PointerTypeNode struct {
		Pointer token.Token
		Typ     Type
	}

	ArrayTypeNode struct {
		Open  token.Token
		Size  Expr
		Close token.Token
		Typ   Type
	}

	MultiTypeNode struct {
		Open  token.Token
		Types []Type
		Close token.Token
	}

	NilableTypeNode struct {
		QuestionMark token.Token
		Type         Type
	}

	GenericTypeNode struct {
		Dollar token.Token
		Type   Type
	}
)

func (b *BaseTypeNode) Start() int      { return b.Typ.Start }
func (b *BaseTypeNode) End() int        { return b.Typ.Pos }
func (b *BaseTypeNode) Kind() AstKind   { return BaseTypeKind }
func (b *BaseTypeNode) Visit(v Visitor) { v.VisitBaseType(b) }
func (b *BaseTypeNode) typeNode()       {}

func (p *PointerTypeNode) Start() int      { return p.Pointer.Start }
func (p *PointerTypeNode) End() int        { return p.End() }
func (p *PointerTypeNode) Kind() AstKind   { return PointerTypeKind }
func (p *PointerTypeNode) Visit(v Visitor) { v.VisitPointerType(p) }
func (p *PointerTypeNode) typeNode()       {}

func (a *ArrayTypeNode) Start() int      { return a.Open.Start }
func (a *ArrayTypeNode) End() int        { return a.End() }
func (a *ArrayTypeNode) Kind() AstKind   { return ArrayTypeKind }
func (a *ArrayTypeNode) Visit(v Visitor) { v.VisitArrayType(a) }
func (a *ArrayTypeNode) typeNode()       {}

func (m *MultiTypeNode) Start() int      { return m.Open.Start }
func (m *MultiTypeNode) End() int        { return m.Close.Pos }
func (m *MultiTypeNode) Kind() AstKind   { return MultiTypeKind }
func (m *MultiTypeNode) Visit(v Visitor) { v.VisitMultiType(m) }
func (m *MultiTypeNode) typeNode()       {}

func (n *NilableTypeNode) Start() int      { return n.QuestionMark.Start }
func (n *NilableTypeNode) End() int        { return n.Type.End() }
func (n *NilableTypeNode) Kind() AstKind   { return NilableTypeKind }
func (n *NilableTypeNode) Visit(v Visitor) { v.VisitNilableType(n) }
func (n *NilableTypeNode) typeNode()       {}

func (g *GenericTypeNode) Start() int      { return g.Dollar.Start }
func (g *GenericTypeNode) End() int        { return g.Dollar.Pos }
func (g *GenericTypeNode) Kind() AstKind   { return GenericTypeKind }
func (g *GenericTypeNode) Visit(v Visitor) { v.VisitGenericType(g) }
func (g *GenericTypeNode) typeNode()       {}
