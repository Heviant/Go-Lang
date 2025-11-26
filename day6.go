package main

// Recursion

import "fmt"

func main() {
	fmt.Println("recursion")

	// pointer

	var ptr *int

	fmt.Println("value of pointer ", ptr)

	number := 23

	ptr = &number // saves or refrences memory address

	fmt.Println("value of ptr ", ptr)  // only have address , have same address as number i.e.: ptr, number == same address
	fmt.Println("value of ptr ", *ptr) // actually points to address value
	print(&number)                     // memory address
}
