package main

// Не понятно зачем делать переменную justString глобальной,
// если она используется только в одной функции.
// Это увеличивает шанс случайно записать в неё что-то
// в будущем из новой функции и увеличивает длину программы.
// var justString string
//
//	func someFunc() {
//		// Длинный и не особо читаемый способ передать int в функцию
//	    v := createHugeString(1 << 10)
//	    // Тут запись int'а уже другая и нет уверенности,
//	    // что заданный индекс находится в пределах длины строки
//	    // Если index out of range, то случится паника
//	    justString = v[:100]
//	}
//
// Данный код приводит к утечке памяти, так как мы создаем строку размером 1 << 10,
// что равно 1024, а потом используем только 100 первых символов, а остаток остается висеть в памяти
// и не освободится пока программа работает. Если мы будем вызывать много раз нашу функцию someFunc()
// память неизбежно будет покидать нас)
//
// Не понятно зачем выносить функционал программы в отдельную функцию,
// если это единственная функция в программе, код читается хуже.
//
//	func main() {
//		 someFunc()
//	}

import (
	"fmt"
)

func main() {
	v := createHugeString(1024)

	// Желаемая длина обычной строки
	var length int
	fmt.Print("Enter length: ")
	fmt.Scan(&length)

	var justString string

	switch {
	case len(v) >= length:
		justString = v[:length]
	case len(v) < length:
		fmt.Printf("Строка длиной %d не может передать %d символов\n", len(v), length)
		fmt.Println("Возвращаем строку максимально возможной длины")
		justString = v
	}

	fmt.Println(justString)
}

func createHugeString(length int) string {
	var hugeString string
	for i := 0; i < length; i++ {
		hugeString += "1"
	}

	return hugeString
}
