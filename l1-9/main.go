package main

import "fmt"

/*
	Разработать конвейер чисел.
	Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
	после чего данные из второго канала должны выводиться в stdout
*/

// fillSourceChannel - функция для заполнения канала sourceCh элементами массива arr
func fillSourceChannel(arr []int, sourceCh chan int) {
	// Итерируемся по слайсу и отправляем значения в канал
	for _, item := range arr {
		sourceCh <- item
	}
	close(sourceCh)
}

// conveyor - конвейер получающий значения из sourceCh и отправляющая квадраты в sqCh
func conveyor(sourceCh, sqCh chan int) {
	for item := range sourceCh {
		sqCh <- item * item
	}
	close(sqCh)
}

// display - вывод
func display(ch2 <-chan int) {
	for item := range ch2 {
		fmt.Println(item)
	}
	fmt.Println()
}

func main() {

	// Создаем каналы
	sourceCh := make(chan int)
	sqCh := make(chan int)

	// Исходный массив
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Запускаем горутину для заполнения sourceCh
	go fillSourceChannel(arr, sourceCh)

	// Запускаем горутину для обработки данных и отправки в sqCh
	go conveyor(sourceCh, sqCh)

	// Выводим результат работы конвейера
	display(sqCh)
}
