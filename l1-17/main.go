package main

import "fmt"

/*
	Реализовать бинарный поиск встроенными методами языка.

*/

func main() {
	fmt.Println(binarySearch([]int{1, 3, 57, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59}, 11)) // 3 индекс
}

// binarySearch - функция для нахождения индекса элемента за O(log(n))
func binarySearch(nums []int, target int) any {

	// Проверка на НЕпустоту массива и на то, что искомое значение не выходит за границы массива
	// Проверку можно реализовать так, потому что массив для бинарного поиска должен быть остортированным по возрастанию
	if (nums == nil || len(nums) == 0) || (target < nums[0] || target > nums[len(nums)-1]) {
		return -1
	}
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2

		// Искомый элемент найден, возвращаем индекс
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target { // Искомый элемент больше, сдвигаем левую границу
			low = mid + 1
		} else { // Искомый элемент меньше, сдвигаем правую границу
			high = mid - 1
		}
	}
	// Искомый элемент не найден, возвращаем -1
	return -1
}
