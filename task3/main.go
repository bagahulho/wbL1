package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	sum := 0
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for _, n := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			mu.Lock()
			sum += n * n
			mu.Unlock()
		}(n)
	}

	wg.Wait()
	fmt.Println(sum)
}
