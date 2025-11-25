package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Weekday()) // time gives all time and Now gives today "2025-11-22 09:19:45.8637846 +0545 +0545 m=+0.000000001"
	// Day() - gives date, Weekday() - gives week name
	fmt.Println(time.Wednesday) // returns wednesday

	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("today is holiday")
	default:
		fmt.Println("workday")
	}

	t := time.Now()
	switch {
	case t.Hour() > 12:
		fmt.Println("it is afternoon")
	default:
		fmt.Println("it is before noon")
	}

	whatami := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("it is boolena")
		case int:
			fmt.Println("it is integer")
		case string:
			fmt.Println("it is string")
		default:
			fmt.Println(" no idea ", t)
		}
	}
	whatami("loll")

	// Array

	var a [5]int
	fmt.Println("default value", a)
	a[3] = 10
	fmt.Println("after set value", a)

	b := [5]int{1, 2, 3, 4}
	fmt.Println(b)
	dynamicArray := [...]int{1, 2, 3, 4, 5, 6, 9, 4} //dynamicArray compiler counts the total number
	fmt.Println("dynamic value ", dynamicArray, " total length", len(dynamicArray))
	eachto := [...]int{100, 4: 200, 300} // up to index 4 it puts 0 in empty place
	fmt.Println("index no. : then number result ", eachto)

	arr := [2][2]int{
		{1, 2},
		{2, 2},
	}
	fmt.Println(arr)
	fmt.Println(arr[1]) // row
	var twoD [3][3]int
	for i := range 3 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("matrix ", twoD)

	noofuser := 133
	fmt.Println(noofuser)
	fmt.Printf("type is %T", noofuser)

}
