package main

import (
	"fmt"
	"time"
)

func main() {
	stopChan := make(chan struct{})

	go func() {
		for {
			select {
			case <-stopChan:
				fmt.Println("stop channel")
				return
			default:
				fmt.Println("working")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	close(stopChan)
}
