package main

import (
	"fmt"
	"l1/l1-12/set"
)

/*
	Имеется последовательность строк - (cat, cat, dog, cat, tree)
	создать для нее собственное множество.


*/

func main() {
	// Создание пустого множества и вывод на экран []
	mySet := set.NewSet()
	mySet.DisplaySet()

	// Добавление элемента в множества
	mySet.Add("cat")
	// Проверка на наличие недобавленного элемента
	fmt.Println(mySet.DoesContains("dog"))
	// Вывод для проверки добавления элемента "cat"
	mySet.DisplaySet()

	// Создадим слайс чтобы затем добавить его в наше множество и проверить функционал
	appendedSlice := []string{"cat", "cat", "dog", "cat", "tree"}

	// Добавляем слайс в наше множество
	mySet.AddSlice(appendedSlice)

	// Вывод результата через метод GetItems()
	fmt.Println(mySet.GetItems())

}
