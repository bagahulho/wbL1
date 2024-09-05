package main

import (
	"fmt"
	"strings"
)

func reversStringSpace(str string) string {
	words := strings.Split(str, " ")
	length := len(words)

	for i := 0; i < length/2; i++ {
		words[i], words[length-1-i] = words[length-1-i], words[i]
	}

	return strings.Join(words, " ")
}

func main() {
	s := "snow dog sun"
	resS := reversStringSpace(s)
	fmt.Println(resS)
}
