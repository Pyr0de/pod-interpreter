package stmt

import (
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

type StmtScope struct {
	Open bool
}

type StmtIf struct {
	Expression group.Group
	Block []Stmt
	Else *StmtIf
}

