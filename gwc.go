package main

import (
	"fmt"
	"os"
)

func Gwc(file *os.File) {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Error: could not get file info for %s: %v\n", file.Name(), err)
		os.Exit(1)
	}
	fmt.Printf("%d %s\n", fileInfo.Size(), fileInfo.Name())
}
