package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	charFlag := flag.Bool("c", false, "Count the total number of bytes in the file")
	flag.Parse()

	filePath := flag.Arg(0)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: could not open file %s: %v\n", filePath, err)
		os.Exit(1)
	}
	defer file.Close()

	if *charFlag {
		Gwc(file)
	} else {
		fmt.Printf("Usage: gwc [-c] <file-path>")
		os.Exit(1)
	}
}
