package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(m map[int]int, mu *sync.Mutex, num int) {
			defer wg.Done()
			mu.Lock()
			m[num] = num
			mu.Unlock()
		}(m, &mu, i)
	}
	wg.Wait()
	fmt.Println(m)
}
