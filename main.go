package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var source = `first line
second line
third line
`

func main() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%v\n", line)
		if err == io.EOF {
			break
		}
	}

	fmt.Println("vim-go")
}
