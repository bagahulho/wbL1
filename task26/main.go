package main

import (
	"fmt"
	"strings"
)

func main() {
	var line string
	fmt.Print("Enter your line: ")
	fmt.Scanln(&line)
	fmt.Println(isUnique(line))
}

func isUnique(s string) bool {
	// меняем регистр на строчные символы
	s = strings.ToLower(s)

	// создаем мапу для хранения символов
	charMap := make(map[rune]bool)

	// проходим по строке, заполняем мапу символами
	for _, c := range s {
		// Если такой символ уже есть, возвращаем false
		if charMap[c] {
			return false
		}
		charMap[c] = true
	}

	return true
}
