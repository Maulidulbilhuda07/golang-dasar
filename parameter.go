package main

func sayHelloTo(firstName string, lastName string) {
	println("hello ", firstName, lastName)
}

func main() {
	firstName := "abil"
	sayHelloTo(firstName, "Bilhuda")
	sayHelloTo("maulidul", "Bilhuda")

}
