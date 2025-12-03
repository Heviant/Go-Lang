package main

import "fmt"

type base struct {
	id int
}

func (a base) describe() string {
	return fmt.Sprintf("base with num %v", a.id)
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{
			id: 1,
		},
		str: "yoyo",
	}
	fmt.Printf("co={num: %v, str: %v}\n", co.id, co.str)
	fmt.Println(co.base.id)

	//Since container embeds base, the methods of base also become methods of a container.
	//  Here we invoke a method that was embedded from base directly on co.
	// basically container struct contains base which have describe so co can access all base function
	fmt.Println("describe func: ", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer: ", d.describe())
}
