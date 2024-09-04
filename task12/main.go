package main

import "fmt"

func main() {
	mas := []string{"cat", "cat", "dog", "cat", "tree"}
	res := makeSet(mas)
	fmt.Println(res)
}

func makeSet(mas []string) []string {
	result := make([]string, 0)
	for _, word := range mas {
		if !exists(result, word) {
			result = append(result, word)
		}
	}
	return result
}

func exists(list []string, target string) bool {
	for _, n := range list {
		if n == target {
			return true
		}
	}

	return false
}
