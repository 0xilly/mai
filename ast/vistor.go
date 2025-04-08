package ast

type Visitor interface {
	VisitLiteralExpr(l *LiteralNode)
	VisitBinaryExpr(b *BinaryExprNode)
	VisitUnaryExpr(u *UnaryExprNode)
	VisitGroupExpr(g *GroupExprNode)
	VisitArrayLiteralExpr(a *ArrayLiteralNode)
	VisitCallExpr(c *CallNode)
	VisitArrayIndexExpr(a *ArrayIndexNode)
	VisitFieldAccessExpr(e *FieldAccessNode)

	VisitBaseType(b *BaseTypeNode)
	VisitPointerType(p *PointerTypeNode)
	VisitArrayType(a *ArrayTypeNode)
	VisitMultiType(m *MultiTypeNode)
	VisitGenericType(g *GenericTypeNode)
	VisitNilableType(n *NilableTypeNode)
}
