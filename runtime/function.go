package runtime

import "github.com/Flaque/boop/parser"

type Function struct {
	name  string
	args  []string
	block parser.IBlockContext
}

func NewFunc(name string, args []string, block parser.IBlockContext) Function {
	return Function{name, args, block}
}
