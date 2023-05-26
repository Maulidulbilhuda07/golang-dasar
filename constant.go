package main

import "fmt"

func main() {
	//constant data tidak dapat berubah

	// const firstName = "Maulidul"
	// const lastName = "Bilhuda"

	// multple variabel

	const (
		firstName = "Maulidul"
		lastName  = "Bilhuda"
	)

	//error
	// firstName = "bil"
	// lastName = "huda"

	fmt.Println(firstName)
	fmt.Println(lastName)

}
