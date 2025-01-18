package stmt

import (
	"fmt"

	"github.com/Pyr0de/pod-interpreter/cmd/eval"
)

func (s StmtPrint)Run() bool {
	v, err := eval.Evaluate(&s.Expression)
	if err {
		return true
	}
	fmt.Println(v)
	return false
}
