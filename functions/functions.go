package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodbye(name string) {
	fmt.Println("Bye", name)
}

func beSociable(name string) {
	sayHello(name)
	fmt.Println("How is the weather?")
	sayGoodbye(name)
}

func main() {
	beSociable("Mijke")
	beSociable("Alice")
}
