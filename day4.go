package main

import "fmt"

func plus() (int, int) {
	return 4, 5
}

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	for index, num := range nums { // u can remove index by using _
		fmt.Println(index, "\n")
		fmt.Println(num, "\n")
		total += num
	}
	fmt.Println(total)
}

func main() {
	res1, res2 := plus()
	fmt.Println(res1)
	fmt.Println(res2)

	_, res3 := plus() // in case u don't need one variable
	fmt.Println(res3)
	// res5 := plus()	// gives error cause it returns 2 values
	fmt.Println("888888888888888")
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{7, 8, 9, 10, 11}
	sum(nums...)

}
