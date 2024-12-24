package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

/*
	Разработать программу,
	которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
	По истечению N секунд программа должна завершаться.
*/

// writer - функция для записи значений в канал из горутины
func writer(ctx context.Context, dataCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Program stops writing...")
			return
		default:
			dataCh <- rand.IntN(100)
			time.Sleep(1 * time.Second)
		}
	}
}

// reader - функция для чтения значений из канала горутиной
func reader(ctx context.Context, processCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		// Проверяем на отмену контекста
		case <-ctx.Done():
			fmt.Println("Program stops reading...")
			return
		case value, ok := <-processCh:
			if !ok {
				return
			}
			fmt.Printf("channel get value: %d\n", value)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	var dur int
	fmt.Println("Enter timer duration (seconds): ")

	// Считываем длительность с клавиатуры
	_, err := fmt.Scanf("%d", &dur)
	if err != nil {
		fmt.Println("wrong timer value", err)
		return
	}

	// Создаем контекст
	ctx := context.Background()

	// Создаем контекст с таймаутом
	withTimeout, cancel := context.WithTimeout(ctx, time.Duration(dur)*time.Second)
	// Отменяем контекст при завершении main
	defer cancel()

	// WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup
	// Создаем небуферизованный канал
	processCh := make(chan int)

	// Увеличиваем счетчик waitGroup на 2 (т.к 2 горутины)
	wg.Add(2)
	// Запускаем две горутины
	go reader(withTimeout, processCh, &wg)
	go writer(withTimeout, processCh, &wg)

	// Ожидаем завершения горутин
	wg.Wait()
}
