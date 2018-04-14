package runtime

import (
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func RuleToInt(rule antlr.ParseTree) (int, error) {
	return strconv.Atoi(rule.GetText())
}

func AnythingToString(val interface{}) (string, error) {

	// Ints
	if num, ok := val.(int); ok {
		return strconv.Itoa(num), nil
	}

	// Floats
	if num, ok := val.(float64); ok {
		return FloatToString(num), nil
	}

	// Strings
	if str, ok := val.(string); ok {
		return str, nil
	}

	return "", fmt.Errorf("Unable to parse value '%s'", val)
}
