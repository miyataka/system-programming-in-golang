package main

import (
	"fmt"
	"sync"
)

var id int

func generateId(mutex *sync.Mutex) int {
	mutex.Lock()
	defer mutex.Unlock()
	id++
	result := id
	return result
}

func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("id: %d\n", generateId(&mutex))
			wg.Done()
		}()
	}
	wg.Wait()
}
