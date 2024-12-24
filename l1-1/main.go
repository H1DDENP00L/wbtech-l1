package main

import (
	"fmt"
	"math/rand/v2"
)

/*
Задание 1
	Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action
	от родительской структуры Human (аналог наследования).
*/

// Human - Наша родительская структура Human с полями и методами
type Human struct {
	name       string
	age        int
	employeeID int
}

func (h *Human) greeting() {
	fmt.Printf("Hello, guys! My name is %s, and I am %d years old. My eployee ID is %d\n", h.name, h.age, h.employeeID)
}

func (h *Human) changeEmployeeID() {

	h.employeeID = rand.IntN(1000)
	fmt.Printf("%s's Employee ID was changed...\n", h.name)
}

// Action - Структура куда встраиваем Human (чтобы унаследовать все поля и методы) `embedding`
type Action struct {
	Human
}

func main() {
	vova := &Human{"Vladimir", 22, 1}
	action := Action{Human: *vova}

	action.greeting()
	action.changeEmployeeID()
	action.greeting()

}
