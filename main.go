package main

import (
	"bufio"
	"fmt"
	"strings"
)

var source = `first line
second line
third line
`

func main() {
	secanner := bufio.NewScanner(strings.NewReader(source))
	for secanner.Scan() {
		fmt.Printf("%v\n", secanner.Text())
	}
}
