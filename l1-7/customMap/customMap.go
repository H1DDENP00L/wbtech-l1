package customMap

import (
	"fmt"
	"sync"
)

type CustomMap struct {
	// Мьютекс для безопасного обращения к общему объекту нескольих горутин
	sync.Mutex
	data map[string]int
}

// NewCustomMap - Конструктор
func NewCustomMap() *CustomMap {
	cm := CustomMap{data: make(map[string]int)}
	return &cm
}

// Put - добавление пары в мапу
func (cm *CustomMap) Put(k string, v int) {
	cm.Lock()
	cm.data[k] = v
	cm.Unlock()
}

// DisplayMap - красивый вывод данных
func DisplayMap(cm *CustomMap) {
	for k, v := range cm.data {
		fmt.Printf("|\tkey:%s | value:%d\t|\n", k, v)
	}
}
