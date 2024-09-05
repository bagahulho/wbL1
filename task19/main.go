package main

import "fmt"

func reversString(str string) string {
	runes := []rune(str)
	length := len(runes)

	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}
	return string(runes)

}

func main() {
	var s string
	fmt.Print("Enter a line: ")
	fmt.Scan(&s)
	resS := reversString(s)
	fmt.Println(resS)
}
