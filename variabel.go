package main

import "fmt"

func main() {

	// var bisa di ganti denga :=

	var name string

	name = "maulidul"
	fmt.Println(name)
	name = "maulidul bilhuda"
	fmt.Println(name)

	//atau tidak perlu menyebut kan tipe data golang akan deek sendiri tipe data nya

	var name2 = "maulidul 2"
	fmt.Println(name2)
	var umur = 26
	fmt.Println(umur)

	negara := "Indonesia"
	fmt.Println(negara)

	var (
		firstName = "abil"
		lastName  = "Huda"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}
