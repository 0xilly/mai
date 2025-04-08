package ast

type AstKind int

const (
	ProgramNode = iota

	BinaryExprNodeKind
	UnaryExprNodeKind
	GroupExprNodeKind

	LiteralNodeKind
	IdentLiteralKind
	ArrayLiteralKind
	CallExprKind
	ArrayIndexExprKind
	FieldAccessExprKind

	BaseTypeKind
	PointerTypeKind
	ArrayTypeKind
	MultiTypeKind
	GenericTypeKind
	NilableTypeKind
)
