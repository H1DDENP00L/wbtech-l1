package main

import (
	"fmt"
	"math/big"
)

/*
	Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a и b,
	значение которых > 2^20.


*/

// AddBigInt - сложение
func AddBigInt(a *big.Int, b *big.Int) *big.Int {
	return big.NewInt(0).Add(a, b)
}

// SubtractBigInt - вычитание
func SubtractBigInt(a *big.Int, b *big.Int) *big.Int {
	return big.NewInt(0).Sub(a, b)
}

// MultiplyBigInt - умножение
func MultiplyBigInt(a *big.Int, b *big.Int) *big.Int {
	return big.NewInt(0).Mul(a, b)
}

// DivideBigInt - деление
func DivideBigInt(a *big.Int, b *big.Int) *big.Int {
	return big.NewInt(0).Div(a, b)
}

func main() {
	a, _ := new(big.Int).SetString("7000000000000000000000000000000000000", 10)
	b, _ := new(big.Int).SetString("2000000000000000000000000000000000000", 10)

	fmt.Println("Результат сложения:", AddBigInt(a, b))
	fmt.Println("Результат вычитания:", SubtractBigInt(a, b))
	fmt.Println("Результат умножения:", MultiplyBigInt(a, b))
	fmt.Println("Результат деления:", DivideBigInt(a, b))
}
