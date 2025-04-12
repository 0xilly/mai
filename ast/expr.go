package ast

import "mai/token"

// Expression
type (
	BinaryExprNode struct {
		Left  Expr
		Op    token.Token
		Right Expr
	}

	UnaryExprNode struct {
		Op    token.Token
		Right Expr
	}

	GroupExprNode struct {
		Open       token.Token
		Expression Expr
		Close      token.Token
	}

	LiteralNode struct {
		Ident token.Token
	}

	ArrayLiteralNode struct {
		Open        token.Token
		Expressions []Expr
		Close       token.Token
	}

	CallNode struct {
		Fn    Expr
		Open  token.Token
		Args  []Expr
		Close token.Token
	}

	ArrayIndexNode struct {
		Open   token.Token
		Target Expr
		Index  Expr
		Close  token.Token
	}

	FieldAccessNode struct {
		Dot    *token.Token
		Parent Expr
		Ident  Expr
	}
)

func (b *BinaryExprNode) Start() int { return b.Left.Start() }
func (b *BinaryExprNode) End() int   { return b.Right.End() }

// func (b *BinaryExprNode) Name() string        { return b.Op.Kind.String() }
func (b *BinaryExprNode) Kind() AstKind   { return BinaryExprNodeKind }
func (b *BinaryExprNode) Visit(v Visitor) { v.VisitBinaryExpr(b) }
func (b *BinaryExprNode) exprNode()       {}

func (u *UnaryExprNode) Start() int { return u.Op.Start }
func (u *UnaryExprNode) End() int   { return u.Right.End() }

// func (u *UnaryExprNode) Name() string        { return u.Op.Kind.String() }
func (u *UnaryExprNode) Kind() AstKind   { return UnaryExprNodeKind }
func (u *UnaryExprNode) Visit(v Visitor) { v.VisitUnaryExpr(u) }
func (u *UnaryExprNode) exprNode()       {}

func (g *GroupExprNode) Start() int { return g.Open.Start }
func (g *GroupExprNode) End() int   { return g.Close.Start }

// func (g *GroupExprNode) Name() string        { return "Group" }
func (g *GroupExprNode) Kind() AstKind   { return GroupExprNodeKind }
func (g *GroupExprNode) Visit(v Visitor) { v.VisitGroupExpr(g) }
func (g *GroupExprNode) exprNode()       {}

func (l *LiteralNode) Start() int      { return l.Ident.Start }
func (l *LiteralNode) End() int        { return l.Ident.Pos }
func (l *LiteralNode) Kind() AstKind   { return LiteralNodeKind }
func (l *LiteralNode) Visit(v Visitor) { v.VisitLiteralExpr(l) }
func (l *LiteralNode) exprNode()       {}

func (a *ArrayLiteralNode) Start() int      { return a.Open.Start }
func (a *ArrayLiteralNode) End() int        { return a.Close.Pos }
func (a *ArrayLiteralNode) Kind() AstKind   { return ArrayLiteralKind }
func (a *ArrayLiteralNode) Visit(v Visitor) { v.VisitArrayLiteralExpr(a) }
func (a *ArrayLiteralNode) exprNode()       {}

func (c *CallNode) Start() int      { return c.Fn.Start() }
func (c *CallNode) End() int        { return c.Close.Pos }
func (c *CallNode) Kind() AstKind   { return CallExprKind }
func (c *CallNode) Visit(v Visitor) { v.VisitCallExpr(c) }
func (c *CallNode) exprNode()       {}

func (a *ArrayIndexNode) Start() int      { return a.Open.Start }
func (a *ArrayIndexNode) End() int        { return a.Close.Pos }
func (a *ArrayIndexNode) Kind() AstKind   { return ArrayIndexExprKind }
func (a *ArrayIndexNode) Visit(v Visitor) { v.VisitArrayIndexExpr(a) }
func (a *ArrayIndexNode) exprNode()       {}

func (f *FieldAccessNode) Start() int {
	if f.Dot == nil {
		return f.Parent.Start()
	}
	return f.Dot.Start
}
func (f *FieldAccessNode) End() int        { return f.Ident.End() }
func (f *FieldAccessNode) Kind() AstKind   { return FieldAccessExprKind }
func (f *FieldAccessNode) Visit(v Visitor) { v.VisitFieldAccessExpr(f) }
func (f *FieldAccessNode) exprNode()       {}
