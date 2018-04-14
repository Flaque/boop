package main

import (
	"fmt"
	"os"

	"github.com/flaque/beep/runtime"
)

func addNewLineToEndOfFile(filename string) {
	f, _ := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	f.WriteString("\n")
	f.Close()
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("File required as input")
		os.Exit(1)
	}

	if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
		fmt.Println("File does not exist")
		os.Exit(1)
	}

	// Fixes ocasional bugs where there's no new line at the end of the file
	// TODO: Definitely not the best way to solve this
	addNewLineToEndOfFile(os.Args[1])

	// Run
	runtime.Run(os.Args[1])
}
