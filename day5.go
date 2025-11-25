package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int { // 1
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intSeq()
	fmt.Println(newInt) // memory address
	fmt.Println(newInt())
}
