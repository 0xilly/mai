package ast

type (
	Node interface {
		Start() int
		End() int
		Kind() AstKind
		Visit(v Visitor)
	}

	File interface {
		Node
		fileNode()
	}

	Decl interface {
		Node
		declNode()
	}

	Expr interface {
		Node
		exprNode()
	}

	Stmt interface {
		Node
		stmtNode()
	}

	Type interface {
		Node
		typeNode()
	}

	Special interface {
		Node
		specialNode()
	}
)
