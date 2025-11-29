package main

import "fmt"

type dim struct {
	length  float64
	breadth float64
	height  float64
}

func rectangle(l float64, b float64) *dim {
	pt := dim{length: l, breadth: b, height: 0}
	return &pt
}

func square(l float64) *dim {
	pt := dim{l, l, l}
	return &pt
}

func (r *dim) area() float64 {
	return r.breadth * r.length
}

func (r *dim) sqarea() float64 { // pointer, if i do any change it will reflect on original
	return r.length * r.length
}

// for example
func (r *dim) ptr_sqarea() float64 { // pointer, if i do any change it will reflect on original
	r.length = 9
	return r.length
}

func (r dim) perimeter() float64 {
	return 2 * r.breadth * r.length
}

func (r dim) sqperimeter() float64 { // values, if i do any change it will only be true for this function and no change to original
	return 4 * r.length
}

// for eg
func (r dim) value_perimeter() float64 { // values, if i do any change it will only be true for this function and no change to original
	r.length = 8.0
	return r.length
}
func main() {
	fmt.Println("")
	rect := rectangle(14, 15)
	sq := square(4)
	fmt.Println("Area of rect: ", rect.area())
	fmt.Println("perimeter of rect: ", rect.perimeter())
	fmt.Println("Area of square: ", sq.sqarea())
	fmt.Println("Area of rect: ", sq.sqperimeter())
	// value
	fmt.Println("len of rect: ", sq.value_perimeter())
	fmt.Println("len of rect: ", sq.length) // no change in actual
	fmt.Println("****************************************** ")
	//  pointer
	fmt.Println("len of rect: ", sq.ptr_sqarea())
	fmt.Println("Area of rect: ", sq.length) // change in actual

}
