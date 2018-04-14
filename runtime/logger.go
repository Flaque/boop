package runtime

import (
	"fmt"
	"io"
)

// Logger lets you print to a specific file instead of just stdout
// which can be useful for testing
type Logger struct {
	out io.Writer
}

func NewLogger(to io.Writer) Logger {
	return Logger{to}
}

func (l *Logger) Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(l.out, a...)
}

func (l *Logger) Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(l.out, a...)
}
