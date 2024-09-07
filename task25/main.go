package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")
	Sleep(3)
	fmt.Println("End after 3 seconds")
}

func Sleep(sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
}
