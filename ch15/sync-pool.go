package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var count int
	pool := sync.Pool{
		New: func() interface{} {
			count++
			return fmt.Sprintf("created: %d", count)
		},
	}

	pool.Put("removed: 1")
	pool.Put("removed: 2")
	runtime.GC()
	time.Sleep(time.Second * 3)
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
}
