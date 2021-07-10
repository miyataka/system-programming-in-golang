package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {
	var stream io.Reader

	cReader := io.NewSectionReader(computer, 0, 1)
	aReader := io.NewSectionReader(programming, 5, 1)
	sReader := io.NewSectionReader(system, 0, 1)
	iReader := io.NewSectionReader(programming, 8, 1)
	pr, pw := io.Pipe()
	writer := io.MultiWriter(pw, pw)
	go io.CopyN(writer, iReader, 1)
	defer pw.Close()
	stream = io.MultiReader(aReader, sReader, cReader, io.LimitReader(pr, 2))

	io.Copy(os.Stdout, stream)
}
