package main

import (
	"fmt"
	"math/rand/v2"
)

/*
	Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

// partition - функция для разбиения массива на две части относительно опорного элемента (pivot)
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	// Итерируемся по массиву, перемещая элементы меньше опорного в левую часть
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			// Индекс обновляем
			i += 1
			// Элементы местами меняем
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Меняем местами опорный элемент с элемнтом i+1 для установки pivot на его место
	arr[i+1], arr[high] = arr[high], arr[i+1]

	// Возврат индекса опорного элемента
	return i + 1
}

// quicksort - рекурсивная функция для быстрой сортировки массива
func quicksort(arr []int, low, high int) {

	// условие остановки рекурсии
	if low < high {
		// Разбиваем массив и получаем индекс опорного элемента
		pi := partition(arr, low, high)

		// Рекурсивно сортируем левую часть
		quicksort(arr, low, pi-1)

		// Рекурсивно сортируем правую часть
		quicksort(arr, pi+1, high)
	}
}

func main() {

	// Тестируем на массиве из 10 рандомных чисел
	arr := make([]int, 10)
	for i := range arr {
		arr[i] = rand.IntN(100)
	}

	fmt.Println("source array = ", arr)
	quicksort(arr, 0, len(arr)-1)
	fmt.Println("sorted array = ", arr)
}
