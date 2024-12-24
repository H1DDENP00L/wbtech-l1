package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
	Реализовать все возможные способы остановки выполнения горутины.
*/

// stopWithChannels - останавливает горутину с помощью закрытия канала
func stopWithChannels(stopCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Goroutine1 starts doing it's job!")
	for {
		select {
		case <-stopCh:
			fmt.Println("!!!Goroutine1(ByChannel) stopped by closing channel\n\n")
			return

		default:
			fmt.Println("Goroutine1 is working...")
			time.Sleep(1 * time.Second)
		}
	}

}

// stopWithContext - останавливает горутину с помощью отмены контекста
func stopWithContext(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Goroutine2 starts doing it's job!")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("!!!Goroutine2(ByContext) stopped by closing context\n\n")
			return
		default:
			fmt.Println("Goroutine2 is working...")
			time.Sleep(time.Second)
		}
	}
}

// stopInsideGoroutine - останавливает горутину внутри самой горутины по условию
func stopInsideGoroutine(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Goroutine3 starts doing it's job!")
	for i := 0; i < 3; i++ {
		fmt.Println("Goroutine3 is working...")
		time.Sleep(1 * time.Second)
	}
	fmt.Println("!!!Goroutine3(ByItself) stopped by itself\n\n")
	return
}

// stopByFlag - останавливает горутину с помощью флага
func stopByFlag(stopFlag *bool, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Goroutine4 starts doing it's job!")
	for {
		if *stopFlag {
			fmt.Println("!!!Goroutine4(ByFlag) stopped by flag\n\n")
			return
		}
		fmt.Println("Goroutine 4 Wis working...")
		time.Sleep(time.Second * 1)
	}
}

func main() {

	// Создаем контекст с возможностью отмены
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup

	// Method 1 - stop by closing channel
	stopCh := make(chan bool)
	wg.Add(1)
	go stopWithChannels(stopCh, &wg) // Запуск горутины
	time.Sleep(time.Second * 2)
	// Закрываем канал
	close(stopCh)

	// Ожидаем завершения горутины
	wg.Wait()

	//Method 2 - by canceling context
	wg.Add(1)
	go stopWithContext(ctx, &wg) // Запуск горутины
	time.Sleep(time.Second * 2)

	// Отменяем конекст
	cancel()
	wg.Wait()

	//Method 3 - stop Goroutine inside itself
	wg.Add(1)
	go stopInsideGoroutine(&wg) // Запуск горутины
	wg.Wait()

	//Method 4 - stop Goroutine by flag
	flag := false
	go stopByFlag(&flag, &wg) // Запуск горутины
	wg.Add(1)
	time.Sleep(time.Second * 2)
	// Меняем флаг спустя 2 секунды для заверщения горутины
	flag = true
	wg.Wait()

}
