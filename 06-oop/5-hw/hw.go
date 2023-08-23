package hw

import (
	"fmt"
	"math"
)

// По условиям задачи, координаты не могут быть меньше 0.
// geom, переименован с маленькой буквы, объект используется только внутри
type geom struct {
	X1, Y1, X2, Y2 float64
}

// CalculateDistance Переименован geom -> g рекомендации наименивания
// CalculateDistance geom -> *geom меняем на указатель, напрямую работать с объектом (не создавать копию)
func (g *geom) CalculateDistance() (distance float64) {

	if g.X1 < 0 || g.X2 < 0 || g.Y1 < 0 || g.Y2 < 0 {
		fmt.Println("Координаты не могут быть меньше нуля")
		return -1
	} else {
		distance = math.Sqrt(math.Pow(g.X2-g.X1, 2) + math.Pow(g.Y2-g.Y1, 2))
	}

	// возврат расстояния между точками
	return distance
}
