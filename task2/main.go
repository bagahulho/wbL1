package main

import (
	"fmt"
	"sync"
)

func main() {
	//Создаем массив чисел, у которых надо вывести квадрат
	arr := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	for _, k := range arr {
		wg.Add(1)
		go func(a int) {
			fmt.Println(a * a)
			wg.Done()
		}(k)
		wg.Wait()
	}
}
