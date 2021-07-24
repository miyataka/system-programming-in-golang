package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	sid, _ := syscall.Getsid(os.Getpid())
	fmt.Fprintf(os.Stderr, "group ID: %d session ID: %d\n", syscall.Getpgrp(), sid)
}
