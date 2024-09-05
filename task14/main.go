package main

import (
	"fmt"
	"reflect"
)

func main() {
	var (
		a int
		b string
		c bool
	)
	ch := make(chan int)
	fmt.Printf("Using reflect: %s, %s, %s, %s\n", reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c), reflect.TypeOf(ch))
	fmt.Printf("Using switch: %s, %s, %s, %s\n", getType(a), getType(b), getType(c), getType(ch))
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
