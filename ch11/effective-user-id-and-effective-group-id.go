package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("user ID: %d\n", os.Getuid())
	fmt.Printf("group ID: %d\n", os.Getgid())
	fmt.Printf("effective user ID: %d\n", os.Geteuid())
	fmt.Printf("effective group ID: %d\n", os.Getegid())
}
