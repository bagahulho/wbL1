package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func findDistance(p1, p2 *Point) float64 {
	res := math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
	return res
}

func main() {
	p1 := NewPoint(10, 20)
	p2 := NewPoint(30, 60)
	fmt.Println(findDistance(p1, p2))
}
