package main

import (
	"fmt"
	"math"
)

type grometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	l, b float64
}

type circle struct {
	r float64
}

func (r rect) area() float64 {
	return r.l * r.b
}
func (r rect) perimeter() float64 {
	return 2 * (r.l + r.b)
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func measurement(g grometry) {
	fmt.Println(g)
	fmt.Println("area", g.area())
	fmt.Println("perimeter", g.perimeter())
}
func detectCircle(g grometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.r)
	}
}
func main() {

	r := rect{12, 23}
	c := circle{5}
	measurement(r)
	measurement(c)
	// fmt.Println("area of rectangle")
	// fmt.Println("perimeter of rectangle")
	// fmt.Println("area of circle")
	// fmt.Println("perimeter of circle")
	detectCircle(c)
	detectCircle(r)
}
