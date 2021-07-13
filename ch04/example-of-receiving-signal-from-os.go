package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// make channel which buffer size is 1
	signals := make(chan os.Signal, 1)
	// receive SIGINT (Ctrl + C)
	signal.Notify(signals, syscall.SIGINT)

	// wait signal
	fmt.Println("Waiting SIGINT (Ctrl + C)")
	<-signals
	fmt.Println("SIGINT arrived")
}
