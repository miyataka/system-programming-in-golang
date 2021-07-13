package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start 5 seconds timer")
	timer := time.After(5 * time.Second)

	<-timer
	fmt.Println("finished 5 seconds")
}
