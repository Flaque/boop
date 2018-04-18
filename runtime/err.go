package runtime

import (
	"os"
)

func ThrowRuntimeError(logger *Logger, message string) {

	logger.Println("Error:", message)
	os.Exit(1)
}
