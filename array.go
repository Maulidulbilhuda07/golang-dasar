package main

import "fmt"

func main() {
	var names [2]string
	names[0] = "maulidul"
	names[1] = "bilhuda"

	fmt.Println(names[0])
	fmt.Println(names[1])

	var value = [3]int{
		80, 70, 89,
	}
	fmt.Println(value)
	fmt.Println(value[2])

	fmt.Println(len(names))
	fmt.Println(len(value))
}
