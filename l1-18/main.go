package main

import (
	"fmt"
	"l1/l1-18/counter"
	"sync"
)

/*
	Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
	По завершению программа должна выводить итоговое значение счетчика.
*/

// worker - воркер для инкрементирования счетчика
func worker(c *counter.Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	c.Inc()
}

func main() {

	// Создаем новый счетчик
	counter1 := counter.NewCounter()
	var wg sync.WaitGroup
	// Создаем 10 воркеров (для 10 горутин соответственно)
	numberOfWorkers := 10
	// Воркеры конкурентно инкрементят счетчик
	for i := 0; i < numberOfWorkers; i++ {
		wg.Add(1)
		// Счетчик будет увеличиваться в отдельной горутине
		go worker(counter1, &wg)
	}

	// Ждем завершения всех горутин
	wg.Wait()
	// Вывод результата...
	fmt.Printf("FINAL COUNTER VALUE = [%d]\n", counter1.GetCurrentValue())
}
