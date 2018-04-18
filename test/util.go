package test

import (
	"bytes"

	"github.com/Flaque/boop/runtime"
)

func run(code string) string {
	var buf bytes.Buffer
	runtime.Run(code, &buf)
	return buf.String()
}
