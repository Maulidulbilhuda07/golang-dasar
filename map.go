package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "abil",
		"address": "padang",
	}

	person["title"] = "programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	var book map[string]string = make(map[string]string)

	book["title"] = "buku golang"
	book["author"] = "abil"
	book["ups"] = "salah"
	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
