package main

import "fmt"

/*
	Удалить i-ый элемент из слайса.
*/

func deleteElementAtIndex(sl []int, index int) []int {

	// Проверка индекса на корректность
	if index < 0 || index >= len(sl) {
		return sl
	}
	// Используем copy для сдвига элементов в левую сторону
	copy(sl[index:], sl[index+1:])
	return sl[:len(sl)-1]
}

func main() {
	fmt.Println(deleteElementAtIndex([]int{1, 2, 3, 4, 5, 6}, 2))
}
