package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt) // Для остановки по нажатию ctrl + c

	var workerNum int
	fmt.Print("Enter number of workers: ")
	_, err := fmt.Scanf("%d", &workerNum)
	if err != nil || workerNum < 1 {
		log.Fatal("Invalid number of workers")
	}

	wg := &sync.WaitGroup{}
	num := make(chan int)

	for i := 1; i < workerNum+1; i++ {
		wg.Add(1)
		go worker(num, wg, i)
	}

	for {
		randNum := rand.Intn(99999999)
		select {
		case <-stopChan:
			fmt.Printf("\nReceived an interrupt, closing channel and stopping workers...\n\n")
			close(num)
			wg.Wait()
			return
		default:
			num <- randNum
			time.Sleep(time.Second)
		}
	}
}

func worker(ch <-chan int, wg *sync.WaitGroup, i int) {
	fmt.Printf("worker %d is started.\n", i)
	defer wg.Done()
	for {
		num, ok := <-ch
		if !ok {
			fmt.Printf("worker %d is shutting down.\n", i)
			break
		}
		fmt.Printf("worker %d: publishing %d\n", i, num)
	}
}
