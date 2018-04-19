package runtime

import (
	"fmt"
	"os"
)

func ThrowRuntimeError(logger *Logger, message string) {

	logger.Println("Error:", message)
	os.Exit(1)
}

func ThrowIncorrectFunctionCall(logger *Logger,
	function string, expected int, actual int) {

	ThrowRuntimeError(logger, fmt.Sprintf("Function '%s' expects %d arguments but got %d!",
		function, expected, actual))
}
