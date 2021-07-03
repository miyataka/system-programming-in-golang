package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, "integer: %d\nstring: %s\nfloat: %f", 1, "hoge", 1.0)
}
