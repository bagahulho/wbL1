package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	mas := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(ch1 chan int, mas []int) {
		for _, num := range mas {
			ch1 <- num
		}
		close(ch1)
		wg.Done()
	}(ch1, mas)

	wg.Add(1)
	go func(ch1 chan int, ch2 chan int) {
		for num := range ch1 {
			ch2 <- num * num
		}
		close(ch2)
		wg.Done()
	}(ch1, ch2)

	wg.Add(1)
	go func(ch2 chan int) {
		for num := range ch2 {
			fmt.Println(num)
		}
		wg.Done()
	}(ch2)

	wg.Wait()
}
