package main

import "fmt"

type position struct {
	x float32
	y float32
}

type car struct {
	brand  string
	color  string
	doors  int
	wheels int
	pos    position
}

// Equivalent to a method, i.e. a function acting on a struct.
// In this case: a method for the car class
func (c *car) drive() {
	c.pos.x++
	c.pos.y++
	fmt.Println(c)
}

// More general Getters and Setters
// Setters:
// Need a pointer "*car"
func (c *car) setColor(color string) {
	c.color = color
}

// Getters:
// Need to specify the return type
func (c car) getColor() string {
	return c.color
}

func main() {

	ferrari := car{"Ferrari", "red", 2, 4, position{1, 3}}
	fmt.Println(ferrari)

	ferrari.drive()
	fmt.Println(ferrari)

	fmt.Println(ferrari.getColor())
	ferrari.setColor("yellow")
	fmt.Println(ferrari.getColor())

	fmt.Println(ferrari)

}
