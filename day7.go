package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}

	tot := 0
	for _, data := range arr {
		tot = data + tot
	}
	fmt.Println("total: ", tot)

	for i, d := range arr {
		if d == 4 {
			fmt.Println("index of d is: ", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana", "c": "pine"}

	for key, data := range kvs {
		fmt.Printf("%T --> %T\n", key, data)
	}

	for k := range kvs {
		fmt.Println("key: ", k)
	}

	for i, c := range "kvs" {
		fmt.Println(i, c)
	}

}
