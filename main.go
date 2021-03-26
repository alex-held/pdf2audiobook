// Package main provides
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world")

	var argLen = len(os.Args)
	if argLen == 0 {
		noArgErr := errors.New("genctl requires at least 1 argument but no argument was provided\n")
		fmt.Fprintf(os.Stderr, "[ERROR]\t %v", noArgErr)
		os.Exit(1)
	}

	pdfPath := os.Args[0]

	msg := fmt.Sprintf("converting %s", pdfPath)


	fmt.Print(msg)
}
