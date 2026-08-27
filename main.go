package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if !strings.HasSuffix(strings.ToLower(inputFile), ".txt") {
		fmt.Println("Error: input file must be a .txt file")
		return
	}

	if !strings.HasSuffix(strings.ToLower(outputFile), ".txt") {
		fmt.Println("Error: output file must be a .txt file")
		return
	}
}
