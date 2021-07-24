package main

import (
	"fmt"
	"os"
)

func main() {
	path, _ := os.Executable()
	fmt.Printf("exec file name: %s\n", os.Args[0])
	fmt.Printf("exec file path: %s\n", path)
}
