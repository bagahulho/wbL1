package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	num := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func(ctx context.Context, ch <-chan int) {
		val := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Reader stopped")
				return
			case val = <-ch:
				fmt.Println(val)
			}
		}
	}(ctx, num)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Writer stopped")
			return
		default:
			num <- rand.Intn(10000000)
			time.Sleep(500 * time.Millisecond)
		}
	}

}
