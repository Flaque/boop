package runtime

import (
	"github.com/Flaque/boop/parser"
)

type Function struct {
	name string
	args []string
	guts parser.IFuncgutsContext
}

func NewFunc(name string, args []string, guts parser.IFuncgutsContext) Function {
	return Function{name, args, guts}
}
