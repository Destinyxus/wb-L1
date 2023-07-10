package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p Point) Distance(q Point) float64 {
	// Находим Евклидово расстояние между двумя точками
	distanceX := p.x - q.x
	distanceY := p.y - q.y
	// Далее вычисляем расстояние, используя теорему Пифагора
	return math.Sqrt(distanceX*distanceX + distanceY*distanceY)
}

func main() {
	p1 := Point{x: 1.5, y: 2.5}
	p2 := Point{x: 4.5, y: 6.5}

	distance := p1.Distance(p2)
	fmt.Printf("Distance is %.2f", distance)
}
