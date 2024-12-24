package main

import "fmt"

/*
	Поменять местами два числа без создания временной переменной.

*/

// ClassicSwap - свап через множественное присвоение
func ClassicSwap(a, b int) {
	fmt.Printf("Before classic swap: a = %d ; b = %d \n", a, b)

	a, b = b, a
	fmt.Printf("After classic swap: a = %d ; b = %d\n", a, b)
	fmt.Println()
}

// SwapByMultiplicationAndDivision - свап через умножение и деление
func SwapByMultiplicationAndDivision(a, b int) {
	fmt.Printf("Before multi-division swap: a = %d ; b = %d \n", a, b)
	if a == 0 {
		a = b
		b = 0
	}
	if b == 0 {
		b = a
		a = 0
	}
	a = a * b
	b = a / b
	a = a / b
	fmt.Printf("After multi-division swap: a = %d ; b = %d\n", a, b)
	fmt.Println()
}

// SwapBySumAndSubtraction - свап через сумму и разность
func SwapBySumAndSubtraction(a, b int) {
	fmt.Printf("Before sum+sub swap: a = %d ; b = %d \n", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("After sum+sub swap: a = %d ; b = %d\n", a, b)
	fmt.Println()
}

func main() {

	// Тестирование всех свапов
	num1, num2 := 10, 20
	ClassicSwap(num1, num2)

	num3, num4 := 30, 40
	SwapByMultiplicationAndDivision(num3, num4)

	num5, num6 := 50, 60
	SwapBySumAndSubtraction(num5, num6)

}
