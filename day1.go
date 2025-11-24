package main

import "fmt"

func Day() {
	var n = 20
	var s = "20"
	const og = "lover"
	fmt.Println("hello")
	fmt.Printf("n is type of %T \n", n)
	fmt.Printf("s is type of %T \n", s)
	fmt.Printf("og is type of %T", og)
	fmt.Println("og is type of %T", og) // will not work only printf can check type

	// if else
	if n := 3; n != 3 {
		fmt.Println("lool")
	}
	var x = 26
	if x < 50 {
		fmt.Println("less than 50")
	} else if x < 25 {
		fmt.Println("less than 25")
	} else {
		fmt.Println("i don't know more than 50")
	}
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}
}
