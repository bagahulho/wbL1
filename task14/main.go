package main

import (
	"fmt"
)

func main() {
	var (
		a int
		b string
		c bool
	)
	ch := make(chan int)
	fmt.Println(getType(a))
	fmt.Println(getType(b))
	fmt.Println(getType(c))
	fmt.Println(getType(ch))
}

func getType(n interface{}) string {
	switch n.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan"
	default:
		return "unknown"
	}
}
