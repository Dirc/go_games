package main

// Features
// 1. Print number of tries
// 2. See if you can tell if the user is lying

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Scan for terminal input
	scanner := bufio.NewScanner(os.Stdin)

	low := 1
	high := 100
	count := 0

	fmt.Println("Please think of a number between", low, "and", high)
	fmt.Println("Press ENTER when ready")
	scanner.Scan()

	for {
		count++
		guess := (low + high) / 2
		fmt.Println("Try", count)
		fmt.Println("I guess the number is", guess)
		fmt.Println("Is that:")
		fmt.Println("a) too high?")
		fmt.Println("b) too low?")
		fmt.Println("c) correct?")
		scanner.Scan()
		response := scanner.Text()

		if response == "a" {
			high = guess - 1
		} else if response == "b" {
			low = guess + 1
		} else if response == "c" {
			fmt.Println("I won!")
			break
		} else {
			fmt.Println("Invalid input")
		}
	}

}
