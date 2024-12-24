package main

import (
	"fmt"
)

/*
	Дана переменная int64.
	Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

// SetBit - устанавливает значение бита (0 или 1) в заданном числе по заданному номеру бита
func SetBit(number int64, bitNumber int64, bitValue int64) (int64, error) {
	// Проверка на корректный value
	if bitValue != 0 && bitValue != 1 {
		return number, fmt.Errorf("bit value must be 0 or 1")
	}
	// Проверка на корректный bitNumber
	if bitNumber < 0 {
		return number, fmt.Errorf("bit number must be >=0")
	}

	// Устанавливаем значение в i-й бит (1 или 0)
	if bitValue == 0 {
		number &^= 1 << bitNumber
	} else {
		number |= 1 << bitNumber
	}
	return number, nil // возвращаем результат и nil если все хорошо
}

func main() {
	var number int64    // Исходное число
	var bitNumber int64 // Номер бита
	var bitValue int64  // Значение бита

	// Считываем ввод пользователя number
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Printf("input error: %v\n", err)
		return
	}

	// Считываем ввод пользователя bitNumber
	fmt.Print("Enter the bit number (>=0): ")
	_, err = fmt.Scan(&bitNumber)
	if err != nil {
		fmt.Printf("input error: %v\n", err)
		return
	}

	// Считываем ввод пользователя bitValue
	fmt.Print("Enter a value for the bit (0 or 1): ")
	_, err = fmt.Scan(&bitValue)
	if err != nil {
		fmt.Printf("input error: %v\n", err)
		return
	}

	// Вызываем функцию SetBit и обрабатываем ошибку
	fmt.Println()
	fmt.Printf("Number:     %d [ %b ]\n", number, number)
	newNumber, err := SetBit(number, bitNumber, bitValue)
	if err != nil {
		fmt.Printf("input error: %v\n", err)
		return
	}
	fmt.Printf("New Number: %d [ %b ]\n", newNumber, newNumber)
}
