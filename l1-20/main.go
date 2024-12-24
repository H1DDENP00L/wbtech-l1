package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
	Разработать программу, которая переворачивает слова в строке.
	Пример: «snow dog sun — sun dog snow».
*/

// reverseWords - функция для переворачивания порядка слов в строке
func reverseWords(str string) (string, error) {

	// Проверяем не пришла ли нам на вход пустая строка
	if len(str) == 0 {
		return str, errors.New("empty string")
	} else if len(str) == 1 { // Обработка строки на 1 слово
		return str, nil
	}

	// Разбиваем строку на слова (используя пробелы и другие разделители)
	words := strings.Fields(str)

	// Переворачиваем слайс слов (меняем местами слова с начала и с конца)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Собираем слова в строку с разделителем (пробел)
	return strings.Join(words, " "), nil
}

func main() {
	displayReversed("snow sun dog")
	displayReversed("hÃllo Göödnight 😊")
	displayReversed("100 200 300")
}

func displayReversed(str string) {
	reversed, err := reverseWords(str)
	if err != nil {
		fmt.Println("Некорректная строка")
	}
	fmt.Printf("Было <%s> ; Стало <%s>\n\n", str, reversed)
}
