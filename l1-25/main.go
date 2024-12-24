package main

import (
	"fmt"
	"time"
)

/*
	Реализовать собственную функцию sleep.

*/

// Используем time.After, которая возвращает канал, в который придет значение по истечению заданного времени.
func mySleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	fmt.Println("Старт программы...")
	mySleep(4 * time.Second)
	fmt.Println("Прошло 4 секнуды")
	mySleep(2 * time.Second)
	fmt.Println("Прошло еще 2 секунды")
}
