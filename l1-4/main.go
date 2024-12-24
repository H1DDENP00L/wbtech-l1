package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
	Реализовать постоянную запись данных в канал (главный поток).
	Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
	Необходима возможность выбора количества воркеров при старте.
	Программа должна завершаться по нажатию Ctrl+C.
	Выбрать и обосновать способ завершения работы всех воркеров


*/

// worker - функция для горутины-воркера
func worker(ctx context.Context, dataCh chan int, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Программа завершена, worker [%d] остановлен\n", id)
			return
		case data, ok := <-dataCh:
			if !ok {
				return
			}
			fmt.Printf("Worker [%d], value: %d\n", id+1, data)
			time.Sleep(1 * time.Second)
		}
	}
}

// generator - функция для горутины-генератора
func generator(ctx context.Context, dataCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker finished it's work")
			return
		default:
			dataCh <- rand.Intn(100)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {

	var numberOfWorkers int
	fmt.Println("Enter number of worker: ")
	// Получаем количество воркеров от пользователя
	_, err := fmt.Scanf("%d", &numberOfWorkers)
	if err != nil {
		fmt.Println("Enter correct number of workers", err)
		return
	}

	// Создаем контекст с возможностью отмены
	ctx, cancel := context.WithCancel(context.Background())
	// Отменяем контекст при выходе из main
	defer cancel()

	// WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup

	dataChan := make(chan int, numberOfWorkers) // Создаем канал с буфером

	// Добавляем 1 горутину в waitGroup
	wg.Add(1)
	// Запускаем генератор в горутине
	go generator(ctx, dataChan, &wg)

	// Запускаем numberOfWorkers воркеров
	for i := 0; i < numberOfWorkers; i++ {
		wg.Add(1)                        // Увеличиваем счетчик waitGroup
		go worker(ctx, dataChan, &wg, i) // Каждый воркер запускается в своей отдельной горутине
	}

	stop := make(chan os.Signal, 1) // Создаем канал для сигналов os
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop // Ожидаем сигнала остановки

	fmt.Println("Trying to stop program...")
	cancel()        // Отменяем контекст (завершение генератора и воркеров)
	close(dataChan) // Закрываем канал (чтобы воркеры тоже завершились)
	wg.Wait()       // Ожидаем завершения всех горутин
	fmt.Println("Program finished with status 200 (OK)")
}
