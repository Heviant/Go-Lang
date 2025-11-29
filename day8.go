package main

import "fmt"

func main() {
	fmt.Println("Structs")

	love := person{"love", 13}

	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(love)
	fmt.Printf("name is %s \n", love.name)
	fmt.Println(createUser())
}

type person struct {
	name string
	age  int
}

type user struct {
	username string
	password string
	active   bool
}

func createUser() *user {
	cus := user{username: "Love@gmail.com", password: "asdasd", active: true}
	return &cus
}
