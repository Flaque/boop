package runtime

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func RuleToInt(rule antlr.ParseTree) (int, error) {
	return strconv.Atoi(rule.GetText())
}
