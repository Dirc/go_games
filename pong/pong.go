package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

//
// Variables
//
const winWidth, winHeight int = 800, 600

//
// Structures and Functions
//
type color struct {
	r, g, b byte
}

type pos struct {
	x, y float32
}

type ball struct {
	pos
	radius int
	xv     float32
	yv     float32
	color  color
}

func (ball *ball) draw(array_of_pixels []byte) {
	for y := -ball.radius; y < ball.radius; y++ {
		for x := -ball.radius; x < ball.radius; x++ {
			if x*x+y*y < ball.radius*ball.radius {
				setPixel(int(ball.x)+x, int(ball.y)+y, ball.color, array_of_pixels)
			}
		}
	}

}

type paddle struct {
	pos
	w     int
	h     int
	color color
}

func (paddle *paddle) draw(array_of_pixels []byte) {
	startX := int(paddle.x) - paddle.w/2
	startY := int(paddle.y) - paddle.h/2

	for y := 0; y < paddle.h; y++ {
		for x := 0; x < paddle.w; x++ {
			setPixel(startX+x, startY+y, paddle.color, array_of_pixels)
		}
	}
}

func setPixel(x, y int, c color, array_of_pixels []byte) {
	// index equals the n-th pixel
	index := (y*winWidth + x) * 4

	// Give the pixel it's RGBA color
	if index < len(array_of_pixels)-4 && index >= 0 {
		array_of_pixels[index] = c.r
		array_of_pixels[index+1] = c.g
		array_of_pixels[index+2] = c.b
	}
}

func main() {

	window, err := sdl.CreateWindow("Pong Game", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tex.Destroy()

	// Array of pixels
	// Example:
	// - An array of 9 pixels can represent a window of 3x3.
	// - Since every pixel consists of RGBA (A= transparancy),
	//   we take an array of 9*4 bytes to represent a window of 3x3.
	// More general:
	// We need a array of width * height * 4.
	array_of_pixels := make([]byte, winWidth*winHeight*4)

	// Make every pixel RED
	for y := 0; y < winHeight; y++ {
		for x := 0; x < winWidth; x++ {
			setPixel(x, y, color{0, 0, 0}, array_of_pixels)
		}
	}

	player1 := paddle{pos{30, float32(winHeight / 2)}, 20, 100, color{255, 255, 255}}
	player2 := paddle{pos{float32(winWidth - 30), float32(winHeight / 2)}, 20, 100, color{255, 255, 255}}

	ball := ball{pos{float32(winWidth / 2), float32(winHeight / 2)}, 20, 0, 0, color{255, 255, 255}}

	// Game loop
	// Ends with quit event
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		player1.draw(array_of_pixels)
		player2.draw(array_of_pixels)
		ball.draw(array_of_pixels)

		tex.Update(nil, array_of_pixels, winWidth*4)
		renderer.Copy(tex, nil, nil)
		renderer.Present()

		sdl.Delay(16)
	}

}
