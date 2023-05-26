package main

import "fmt"

func main() {
	name := "abil"

	switch name {
	case "eko":
		fmt.Println("nama eko")
	case "joko":
		fmt.Println("hy joko")
	default:
		fmt.Println("hy abil")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")

	case false:
		fmt.Println("mantap")
	}

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("nama teralu panjang")
	case length > 5:
		fmt.Println("nama terlalu panajng")
	default:

		fmt.Println("pas")
	}
}
