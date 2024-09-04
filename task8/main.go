package main

import (
	"fmt"
	"log"
)

func main() {
	var n, i, bit int64

	fmt.Print("In which number do you want to change a bit: ")
	fmt.Scan(&n)

	fmt.Print("Which bit to change: ")
	fmt.Scan(&i)
	if i < 0 || i > 63 {
		log.Fatal("Invalid bit index")
	}

	fmt.Print("Which bit to put in? (0 or 1): ")
	fmt.Scan(&bit)

	fmt.Printf("Part 1: binary: %05b, dec: %d\n", n, n)

	if bit == 0 {
		n &^= 1 << i
	} else if bit == 1 {
		n |= 1 << i
	} else {
		log.Fatalf("%d is an invalid number\n", bit)
	}

	fmt.Printf("Part 2: binary: %05b, dec: %d\n", n, n)
}
