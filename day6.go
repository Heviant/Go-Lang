package main

// Recursion

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println("recursion")

	// pointer

	var ptr *int

	fmt.Println("value of pointer ", ptr)

	number := 23

	ptr = &number // saves or refrences memory address

	fmt.Println("value of ptr ", ptr)  // only have address , have same address as number i.e.: ptr, number == same address
	fmt.Println("value of ptr ", *ptr) // actually points to address value
	print(&number, "\n")               // memory address

	// ********
	fmt.Println("****************************************")

	data := fact(6)
	fmt.Println("factorial of 6: ", data)

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Print(fib(6))

}
