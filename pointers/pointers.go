package main

import "fmt"

// Why we need pointers?

func addOne(num int) {
	num = num + 1
}

func main() {
	x := 5
	fmt.Println(x)

	// Get the pionter of x with &
	xPtr := &x
	// var xPtr *int = &x
	fmt.Println(xPtr)

	addOne(x)
	fmt.Println(x)

}
