package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	reader := io.LimitReader(rand.Reader, 1024)
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(file, reader)
}
