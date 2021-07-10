package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("strings reader")
	copyN(os.Stdout, reader, 5)
}

func copyN(dst io.Writer, src io.Reader, length int64) (written int64, err error) {
	written, err = io.Copy(dst, io.LimitReader(src, length))
	if written == length && err != nil {
		return written, nil
	}

	if written < length && err == nil {
		err = io.EOF
	}
	return
}
