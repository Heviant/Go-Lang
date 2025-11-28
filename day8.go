package main

import "fmt"

func main() {
	fmt.Println("Structs")

	love := person{"love", 13}

	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(love)
	fmt.Printf("name is %s", love.name)
}

type person struct {
	name string
	age  int
}
