package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("user ID: %d\n", os.Getuid())
	fmt.Printf("group ID: %d\n", os.Getgid())
	groups, _ := os.Getgroups()
	fmt.Printf("sub group IDs: %v\n", groups)
}
