package set

import "fmt"

type MySet struct {
	contains map[string]bool
	items    []string
}

// NewSet - Конструктор
func NewSet() *MySet {
	return &MySet{
		contains: make(map[string]bool),
		items:    make([]string, 0),
	}
}

// DisplaySet - Вывод всего множества
func (set *MySet) DisplaySet() {
	fmt.Println(set.items)
}

// GetItems - функция чтобы вернуть все значения
func (set *MySet) GetItems() []string {
	return set.items
}

// AddSlice - Функция вставки слайса
func (set *MySet) AddSlice(sl []string) {
	for _, item := range sl {
		set.Add(item)
	}
}

// Add - Функция добавление значения
func (set *MySet) Add(item string) {
	// Проверка на наличие добавляемого элемента в множестве (множество - набор УНИКАЛЬНЫХ значения)
	if status := set.DoesContains(item); !status {
		set.items = append(set.items, item)
		set.contains[item] = true
	}
}

// Remove - Функция удаления значения
func (set *MySet) Remove(item string) {
	if status := set.DoesContains(item); status {
		delete(set.contains, item)
		set.contains[item] = false
	}
}

// DoesContains - Функция проверяющая наличие элемента в множестве
func (set *MySet) DoesContains(item string) bool {
	_, status := set.contains[item]
	return status
}
