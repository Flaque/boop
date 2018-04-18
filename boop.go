package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Flaque/boop/runtime"
)

func killf(str string, args ...interface{}) {
	fmt.Printf(str+"\n", args...)
	os.Exit(1)
}

func read(filename string) string {
	dat, err := ioutil.ReadFile(filename)

	if err != nil {
		killf("Error reading file %s. Messagife: %s", filename, err)
	}

	return string(dat)
}

func addNewLineToEndOfFile(filename string) {
	f, _ := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	f.WriteString("\n")
	f.Close()
}

func main() {
	if len(os.Args) <= 1 {
		killf("File required as input!")
	}

	if val, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
		killf("File '%s' does not exist", val)
	}

	code := read(os.Args[1])

	valFlag := len(os.Args) > 2 && os.Args[2] == "-c"
	if valFlag {
		runtime.Validate(code)
		return
	}

	// Main run command
	runtime.Run(code, os.Stdout)
}
