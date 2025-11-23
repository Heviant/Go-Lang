package main

import (
	"fmt"
	"math"
)

const name string = "Loverr"

func main() {
	fmt.Println("Hello from learnCodeonLine.in")
	fmt.Println("1+1 = ", 1+1)
	fmt.Println("1.0/2.0 = ", 1.0/2.0)
	fmt.Println("false && true = ", false && true)
	fmt.Println("false && true = ", false || true)
	fmt.Println("!false = ", !false)
	var a = 1
	var v int = 3 //declare type
	fmt.Print(a)
	fmt.Print(v)

	f := "sdadasd" // := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple"
	fmt.Println(f)

	const n = 50000
	var d = 1 / n
	fmt.Println(d)
	fmt.Println(math.Sin(n))
	var i int = 1
	for i < 3 {
		println(i)
		i = i + 1
	}
	for j := 1; j <= 10; j++ {
		fmt.Println(j + 1)
	}
	for i := range 3 {
		fmt.Println(i)
	}
	for n := range 5 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	Day()
}
