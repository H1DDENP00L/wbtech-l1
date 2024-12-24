package counter

import (
	"sync"
)

type Counter struct {
	sync.Mutex
	count int
}

// NewCounter - конструктор
func NewCounter() *Counter {
	c := Counter{count: 0}
	return &c
}

func (c *Counter) Inc() {
	c.Lock()         // Лочим мьютекс
	defer c.Unlock() // Анлочим мьютекс после завршения функции
	c.count++        // Увеличение счетчика
}

// GetCurrentValue - функция для получения текущего значения счетчика
func (c *Counter) GetCurrentValue() int {
	return c.count
}
