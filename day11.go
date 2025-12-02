package main

import "fmt"

type day int

const (
	Sunday day = iota // auto increament from 0, iota initalize the first value as 0 and auto increase to other
	Monday
	Tuesday
)

var foodstatus = map[day]string{
	Sunday:  "Egg",
	Monday:  "No egg",
	Tuesday: "Meat",
}

func (s day) String() string {
	return foodstatus[s]
}

func main() {
	fmt.Println("Returns interger number: ", Monday) // fmt.Println("Returns integer number: ", Monday.String()) internal workings

	ns := transistion(Sunday)
	fmt.Println(ns) //

}

func transistion(s day) day {
	switch s {
	case Sunday:
		return Sunday
	case Monday, Tuesday:
		return Tuesday
	default:
		return Monday
	}
}
