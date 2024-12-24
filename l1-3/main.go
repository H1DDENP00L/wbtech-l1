package main

import (
	"fmt"
	"sync"
)

/*
	Дана последовательность чисел: 2,4,6,8,10.
	Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.
*/

func main() {

	nums := []int{2, 4, 6, 8, 10}
	sum := 0
	// Мьютекс для синхронизации доступа к переменной sum
	var mu sync.Mutex

	// WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Итерируемся по массиву чисел
	for _, num := range nums {
		// Увеличиваем счетчик waitGroup
		wg.Add(1)

		// Запускаем горутину для вычисления и добавления квадрата числа к сумме
		go func(num int, wg *sync.WaitGroup) {
			defer wg.Done()  // Уменьшаем счетчик waitGroup при завершении горутины
			mu.Lock()        // Блокируем доступ к sum
			sum += num * num // Добавляем квадрат числа к сумме
			mu.Unlock()      // Разблокируем доступ к sum

		}(num, &wg)
	}
	wg.Wait() // Ожидаем завершения всех горутин
	fmt.Println("Result = ", sum)
}
