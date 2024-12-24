package main

import "fmt"

/*
	Реализовать пересечение двух неупорядоченных множеств
*/

// Intersection - функция для нахождения пересечения двух множеств
func Intersection(set1, set2 []int) []int {

	// Создаем мапу (множество) из элементов set1
	values := make(map[int]struct{})
	for _, v := range set1 {
		values[v] = struct{}{}
	}

	// Создаем слайс для результата
	result := make([]int, 0)

	// Итерируемся по set2 и проверяем наличие элементов в values
	for _, v := range set2 {
		if _, ok := values[v]; ok {
			result = append(result, v)
		}
		// Удаляем элемент из множества values, чтобы избежать дубликатов в результате
		delete(values, v)
	}
	// Возвращаем слайс с пересечением множеств
	return result
}

func main() {
	// Создаем множества
	set1 := []int{1, 2, 8, 10, 20, 5, 6}
	set2 := []int{1, 9, 4, 5, 0, -20, 8}

	// Выводим пересечения
	fmt.Println(Intersection(set1, set2)) // [1 5 8]
}
