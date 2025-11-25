package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	var s []string
	fmt.Println("uninitialized", s, s == nil, len(s) == 0)
	s = make([]string, 3)
	fmt.Println("after slice", s, len(s))
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("after adding value", s)
	s = append(s, "d")
	fmt.Println("after appending value", s)

	c := make([]string, len(s))
	copy(c, s)
	print("i copied in c : ", c)      // gives hexa value of last variable and show total length actual result [4/4]0xc000072040 -- [4/4] - index/total
	print("\ni copied in c : ", c[3]) // gives actual value
	print("\ni copied in c : ", &c, "\n")

	l := c[1:3] // low index : high index -- doesn't return the value of high index
	fmt.Println(l)

	l = c[2:] // includes low point index
	fmt.Println(l)

	l = c[:3] // doesn't include high point index
	fmt.Println(l)

	t := []string{"dad", "manish", "sexy"}
	fmt.Println("initialize and declared: ", t)
	t2 := []string{"dad", "manish", "sexy"}

	if slices.Equal(t, t2) {
		fmt.Println("t == t2 ")
	}

	twoDim := make([][]int, 3)
	for i := range 3 {
		inner := i + 1
		twoDim[i] = make([]int, inner) //twodim[1] = ["",""]
		for j := range inner {
			twoDim[i][j] = i + j //10-1,11-2 20-2,3,4
		}
	}
	fmt.Println(twoDim)

	// Maps
	m := make(map[string]int) // map or hash or key-value pair
	m["k1"] = 1
	m["k2"] = 3
	fmt.Println("maps: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("length of m: ", len(m))

	delete(m, "k1") // deletes k1 key and value
	fmt.Println(m)

	clear(m) // removes all key value pair
	fmt.Println(m)

	ke, prs := m["k2"] // returns value and boolean if value exist then returns (value, true) if not present then (0, false)
	fmt.Println(ke, prs)

	n1 := map[string]int{"foo": 1, "baa": 2}
	fmt.Println("maps ", n1)
	n2 := map[string]int{"foo": 1, "baa": 2}
	if maps.Equal(n1, n2) {

		fmt.Println("n1 == n2")
	}

	fmt.Println("*****************************")
	fmt.Println(plus(10, 12))
	minus(10, 12)
	fmt.Println(pluss(10, 12, 13))
	res := plus(1, 3)
	fmt.Println(res)
}

func plus(a int, b int) int { // func func_name (parameter paramater_type) return type{}  -- incase of void leave return_type empty
	return a + b
}
func minus(a int, b int) { // no return or void
	fmt.Println(a - b)
}

func pluss(a, b, c int) int { // declare in one go -- abc will be int type
	return a + b
}

// func pluss(a := 16, b int, c int) int { // don't support place holder like use default value in case of a absent
// 	return a + b
// }
