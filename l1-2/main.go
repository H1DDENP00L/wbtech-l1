package main

import (
	"fmt"
	"sync"
)

/*
	Дана последовательность чисел: 2,4,6,8,10.
	Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

func main() {

	// WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup

	// Исходный массив чисел
	nums := []int{2, 4, 6, 8, 10}

	// Итерируемся по массиву чисел
	for _, num := range nums {
		wg.Add(1) // Увеличиваем счетчик waitGroup на 1

		// Запускаем горутину для вывода квадрата числа
		go func(num int, wg *sync.WaitGroup) {
			defer wg.Done()         // Уменьшаем счетчик waitGroup при завершении горутины
			fmt.Print(num*num, " ") // Выводим квадрат числа
		}(num, &wg)
	}
	wg.Wait()

}
