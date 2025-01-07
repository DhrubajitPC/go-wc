package main

import (
	"bufio"
	"fmt"
	"os"
)

func CountChar(file *os.File) {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Error: could not get file info for %s: %v\n", file.Name(), err)
		os.Exit(1)
	}
	fmt.Printf("%d %s\n", fileInfo.Size(), fileInfo.Name())
}

func CountLines(file *os.File) {
	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading file: %v\n", err)
	}
	fmt.Printf("%d %s\n", lineCount, file.Name())
}
