package main

import "fmt"

func main() {
	var name = "abil"

	if name == "abill" {
		fmt.Println(name)
	} else if name == "abil" {
		fmt.Println("hy abil")
	} else {
		fmt.Println("else")
	}

	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("masuk lese")

	}
}
