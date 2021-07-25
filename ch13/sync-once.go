package main

import (
	"fmt"
	"sync"
)

func initialize() {
	fmt.Println("initialzing...")
}

var once sync.Once

func main() {
	once.Do(initialize)
	once.Do(initialize)
	once.Do(initialize)
}
