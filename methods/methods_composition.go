package main

import "fmt"

type position struct {
	x float32
	y float32
}

// Composition of stucts:
// Car can inherit from the position struct.
// This is called composition.
type car struct {
	brand  string
	color  string
	doors  int
	wheels int
	position      // Composition: call position without a variable name
}

// Composition: We can now call .x and .y directly
func (c *car) drive() {
	c.x++
	c.y++
	fmt.Println(c)
}

// Composition: This method on position can also be used on a car object directly.
funct (p position) setX(x float32) {
	p.x = x
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


	// Composition:
	// Use method on position directly on the car object.
	fmt.Println(ferrari)


}
