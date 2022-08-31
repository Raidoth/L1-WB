package main

import (
	"fmt"
	"math"
)

/*
24.	Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

//create struct Point with coordinate x,y
type Point struct {
	x, y float64
}

//initialization point
func CreatePoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

//function compute distance points
func (p *Point) DistancePoints(dot *Point) float64 {
	return math.Sqrt(math.Pow((dot.x-p.x), 2) + math.Pow((dot.y-p.y), 2))
}

func main() {
	//create point a
	a := CreatePoint(3, 1)
	//create point b
	b := CreatePoint(2, 3)
	//compute distance
	result := a.DistancePoints(b)
	//output result
	fmt.Printf("First dot: %.1f,%.1f\nSecond dot: %.1f,%.1f\nDistance: %.4f\n", a.x, a.y, b.x, b.y, result)

}
