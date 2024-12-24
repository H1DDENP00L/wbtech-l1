package main

import (
	"fmt"
	"l1/l1-24/point"
)

/*
	Разработать программу нахождения расстояния между двумя точками,
	которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

func main() {
	pointA := point.NewPoint(10, 20)
	pointB := point.NewPoint(0, 30)

	fmt.Println("Расстояние от точки A о точки B =", pointA.CalcDistance(pointB))

}
