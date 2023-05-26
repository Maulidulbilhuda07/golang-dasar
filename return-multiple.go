package main

import "fmt"

func getFullName() (string, string) {
	return "maulidul", "bilhuda"

}

func main() {

	// untuk menhirsukan suatu return
	firstName, _ := getFullName()

	fmt.Println(firstName)
	//fmt.Println(lastName)
}
