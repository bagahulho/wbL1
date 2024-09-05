package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := binarySearch(s, 5)
	fmt.Println(index)
}

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	if right < left {
		return -1
	}

	for left <= right {
		mid := (left + right) / 2 // Находим индекс центра
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target { // Если искомый элемент  больше центрального элемента
			left = mid + 1 // Значит он находится справа и двигаемся в право, отсекая левую часть
		} else {
			right = mid - 1
		}
	}
	return -1 // Элемент не найден
}
