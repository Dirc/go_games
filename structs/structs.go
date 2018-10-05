package main

import "fmt"

type position struct {
	x float32
	y float32
}

type badGuy struct {
	name   string
	health int
	pos    position
}

func whereIsBadGuy(b badGuy) {
	x := b.pos.x
	y := b.pos.y
	fmt.Println(x, y)
}

func main() {

	/*
		var p position
		x := 3
		y := 4
	*/

	p := position{3, 4} // Note the curly brackets
	fmt.Println(p.x, p.y)

	badguy := badGuy{"Jabba The Hut", 100, p}
	fmt.Println(badguy)
	whereIsBadGuy(badguy)

}
