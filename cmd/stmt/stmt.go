package stmt

import (
	"github.com/Pyr0de/pod-interpreter/cmd/env"
	"github.com/Pyr0de/pod-interpreter/cmd/group"
	"github.com/Pyr0de/pod-interpreter/cmd/token"
)

type Stmt struct {
	Stype token.TokenType
	Statement StmtRun
}

type StmtRun interface {
	Run() bool
}

type StmtPrint struct {
	Expression group.Group
}

type StmtAssign struct {
	Expression group.Group
	Init bool
}

type StmtExpression struct {
	Expression group.Group
}

type StmtBlock struct {
	Block []Stmt
}

type StmtIf struct {
	Expression group.Group
	Block StmtBlock
	Else *StmtIf
}

type StmtWhile struct {
	Expression group.Group
	Block StmtBlock
}

type StmtFor struct {
	Initialization Stmt
	Condition group.Group
	Step Stmt
	Block StmtBlock
}

type StmtFunc struct {
	Name token.Token
	Parameters []token.Token
	Block StmtBlock
	Env *env.Environment
}

type StmtFuncCall struct {
	Name token.Token
	Parameters []group.Group
}

type StmtEmpty struct {}
