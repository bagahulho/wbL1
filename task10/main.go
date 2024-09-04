package main

import "fmt"

func main() {
	mas := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	res := makeSets(mas)
	fmt.Println(res)
}

func makeSets(mas []float64) map[int][]float64 {
	res := make(map[int][]float64)
	for _, temp := range mas {
		ind := int(temp) - int(temp)%10
		if !exists(res[ind], temp) {
			res[ind] = append(res[ind], temp)
		}
	}

	return res
}

func exists(list []float64, target float64) bool {
	for _, n := range list {
		if n == target {
			return true
		}
	}

	return false
}
