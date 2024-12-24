package main

import (
	"fmt"
	"strings"
)

/*
	Разработать программу, которая проверяет,
	что все символы в строке уникальные (true — если уникальные, false etc).
	Функция проверки должна быть регистронезависимой.
*/

// areAllCharsUnique - функция проверки на уникальность символов в строке
func areAllCharsUnique(str string) bool {
	//  Делаем регистронезависимым
	str = strings.ToLower(str)
	chars := make(map[rune]bool)
	for _, char := range str {
		// Если символ уже встречался в карте - возвращаем false
		if _, exists := chars[char]; exists {
			return false
		}
		//  В противном случае добавляем его в мапу
		chars[char] = true
	}
	return true
}

func main() {
	// Тестируем разные строки
	testStrings := []string{"abcd", "abCdefAaf", "aabcd", "", "A", "aA", "ab c"}
	for _, str := range testStrings {
		fmt.Printf("<%s>: %t\n", str, areAllCharsUnique(str))
	}
}
