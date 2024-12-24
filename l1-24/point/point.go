package point

import "math"

type Point struct {

	// Поля структуры начинаются с маленькой буквы, что инкапсулирует их от доступа извне пакета
	x, y float64
}

// NewPoint - конструктор для создания точки
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p1 *Point) CalcDistance(p2 *Point) float64 {

	// Евклидово расстояние
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}
