package main

import "fmt"

type color struct {
	r, g, b byte
}

var red = color{255, 0, 0}

func main() {

	fmt.Println(red.r, red.g, red.b)

}
