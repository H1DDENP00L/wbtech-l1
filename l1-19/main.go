package main

import "fmt"

/*
	Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.

*/

// reverseString - функция для переворачивания строки (учитывает unicode)
func reverseString(str string) string {

	// Обработка пустой строки и строки из 1 символа
	if len(str) == 0 || len(str) == 1 {
		return str
	}

	// Преобразуем строку в слайс рун (чтобы корректно работало с Unicode)
	symbols := []rune(str)

	// Итерируемся по рунам и меняем местами (с начала и с конца)
	for i, j := 0, len(symbols)-1; i < j; i, j = i+1, j-1 {
		symbols[i], symbols[j] = symbols[j], symbols[i]
	}
	// Создаем строку из слайса рун
	return string(symbols)
}

func main() {

	fmt.Println(reverseString("главрыба")) // абырвалг

	fmt.Println(reverseString("再见zàijiàn")) // nàijiàz见再

}
