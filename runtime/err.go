package runtime

import (
	"fmt"
	"os"
)

func ThrowRuntimeError(message string) {

	fmt.Println(message)
	os.Exit(1)
}
