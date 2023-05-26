package main

import "fmt"

func main() {
	type NoKTP string
	type Maried bool

	var noKTPAbil NoKTP = "1234567890998"
	var marriedStatus Maried = true

	fmt.Println(noKTPAbil)
	fmt.Println(marriedStatus)
}
