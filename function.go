package main

import "fmt"

func sayHello() {

	fmt.Println("hello")

}

func main() {

	for i := 0; i < 6; i++ {
		sayHello()
	}

}
