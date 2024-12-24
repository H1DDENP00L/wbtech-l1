package main

import (
	"context"
	"fmt"
	"l1/l1-7/customMap"
	"math/rand/v2"
	"sync"
	"time"
)

/*
	Реализовать конкурентную запись данных в map.

*/

// Функция для записи в customMap
func worker(ctx context.Context, id int, cm *customMap.CustomMap, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker [%d] finished it's job...\n", id)
			return

		default:
			cm.Put(fmt.Sprintf("%d", id), rand.IntN(1000))
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Создаем нашу мапу
	cm := customMap.NewCustomMap()

	// Задаем число горутин
	numberOfWorkers := 7

	wg.Add(numberOfWorkers)
	// Конкурентно записываем данные в мапу
	for i := 0; i < numberOfWorkers; i++ {
		go worker(ctx, i, cm, &wg)

	}
	fmt.Println("Concurrent recording of data to map...")
	time.Sleep(time.Second * 3)

	fmt.Println()
	cancel()
	wg.Wait()
	fmt.Printf("\n")
	customMap.DisplayMap(cm)
}
