package main

import (
	"fmt"
	"math/big" // Пакет "math/big" позволяет работать с огромными числами
)

func main() {
	// Инициализируем наши большие переменные
	a := big.NewInt(8388608)  // 2^22
	b := big.NewInt(16777216) // 2^23

	// Выполняем арифмитические операции

	mulResult := new(big.Int).Mul(a, b)
	fmt.Println(mulResult)

	divResult := new(big.Float).Quo(new(big.Float).SetInt(a), new(big.Float).SetInt(b))
	fmt.Println(divResult)

	sumResult := new(big.Int).Add(a, b)
	fmt.Println(sumResult)

	subResult := new(big.Int).Sub(a, b)
	fmt.Println(subResult)
}
