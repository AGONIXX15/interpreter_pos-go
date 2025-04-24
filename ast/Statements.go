package ast

type AssignmentStmt struct {
	Name  string
	ExplicitType Type
	Value Expr
	Inferred bool
}

func (as AssignmentStmt) stmt() {}

type BlockStmt struct {
	Body []Stmt
}

func (b BlockStmt) stmt() {}

type ExpressionStmt struct {
	Expression Expr
}

type Parameter struct {
	Name string
	TypeParameter Type
}

func (es ExpressionStmt) stmt() {}

type FunctionStmt struct {
	Name string
	TypeReturn Type
	Parameters []Parameter
	Body BlockStmt
}

func (fc FunctionStmt) stmt() {}

type ReturnStmt struct {
	Value Expr
}

func (rs ReturnStmt) stmt() {}

type IfCondition struct {
	Condition Expr
	Body BlockStmt
}

type IfStmt struct {
	Branches []IfCondition
	Else BlockStmt
}

func (is IfStmt) stmt(){}



