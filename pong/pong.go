package main

// Features
//
// - framerate
// - Score
// - Game over state
// - two player or PC
// - AI needs to be inperfect
// - resizing of the window

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

//
// Variables
//
const winWidth, winHeight int = 800, 600

var ballStartPos pos = pos{float32(winWidth / 2), float32(winHeight / 2)}
var ballStartXV float32 = 5
var ballStartYV float32 = 5

var paddleVelocity float32 = 5

var score1, score2 int = 0, 0

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
	radius float32
	xv     float32
	yv     float32
	color  color
}

func (ball *ball) draw(array_of_pixels []byte) {
	for y := -ball.radius; y < ball.radius; y++ {
		for x := -ball.radius; x < ball.radius; x++ {
			if x*x+y*y < ball.radius*ball.radius {
				setPixel(int(ball.x+x), int(ball.y+y), ball.color, array_of_pixels)
			}
		}
	}
}

func (ball *ball) update(leftPaddle *paddle, rightPaddle *paddle, score1 int, score2 int) {
	ball.x += ball.xv
	ball.y += ball.yv

	// Collisions
	// Collision: Top/bottom screen
	if ball.y-ball.radius < 0 || ball.y+ball.radius > float32(winHeight) {
		ball.yv = -ball.yv
	}
	// Collision: left/right screen - score
	if ball.x < 0 {
		ball.x = ballStartPos.x
		score2++
	}
	if int(ball.x) > winWidth {
		ball.y = ballStartPos.y
		score1++
	}
	// Collision: left paddle
	if ball.x-ball.radius < leftPaddle.x+leftPaddle.w/2 {
		// ball between paddle bottom and paddle top
		if ball.y+ball.radius < leftPaddle.y+leftPaddle.h/2 && ball.y-ball.radius > leftPaddle.y-leftPaddle.h/2 {
			ball.xv = -ball.xv
		}
	}
	// Collision: right paddle
	if ball.x+ball.radius > rightPaddle.x-rightPaddle.w/2 {
		// ball between paddle bottom and paddle top
		if ball.y+ball.radius < rightPaddle.y+rightPaddle.h/2 && ball.y-ball.radius > rightPaddle.y-rightPaddle.h/2 {
			ball.xv = -ball.xv
		}
	}
}

type paddle struct {
	pos
	w     float32
	h     float32
	color color
}

func (paddle *paddle) draw(array_of_pixels []byte) {
	startX := int(paddle.x - paddle.w/2)
	startY := int(paddle.y - paddle.h/2)

	for y := 0; y < int(paddle.h); y++ {
		for x := 0; x < int(paddle.w); x++ {
			setPixel(startX+x, startY+y, paddle.color, array_of_pixels)
		}
	}
}

func (paddle *paddle) update(keyState []uint8) {
	if keyState[sdl.SCANCODE_UP] != 0 {
		paddle.y -= paddleVelocity
	}
	if keyState[sdl.SCANCODE_DOWN] != 0 {
		paddle.y += paddleVelocity
	}
}

func (paddle *paddle) aiUpdate(ball *ball) {
	paddle.y = ball.y

}

// make all pixels black again
func clear(array_of_pixels []byte) {
	for i := range array_of_pixels {
		array_of_pixels[i] = 0
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

	// Give every pixel a color
	for y := 0; y < winHeight; y++ {
		for x := 0; x < winWidth; x++ {
			setPixel(x, y, color{0, 0, 0}, array_of_pixels)
		}
	}

	player1 := paddle{pos{30, float32(winHeight / 2)}, 20, 100, color{255, 255, 255}}
	player2 := paddle{pos{float32(winWidth - 30), float32(winHeight / 2)}, 20, 100, color{255, 255, 255}}

	ball := ball{ballStartPos, 20, ballStartXV, ballStartYV, color{255, 255, 255}}

	// Keyboard input
	// Define variable for keyState
	//   keyState := sdl.GetKeyboardState()
	//
	// Verify if UP key is used:
	//   if keyState[sdl.SCANCODE_UP] != 0 {}
	// Verify if DOWN key is used:
	//   if keyState[sdl.SCANCODE_DOWN] != 0 {}
	//
	// See for example: update function for paddle.

	keyState := sdl.GetKeyboardState() // array of bytes "uint8", is used in paddle.update()

	// Game loop
	// Ends with quit event
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		clear(array_of_pixels)

		//fmt.Println(score1, score2)

		player1.update(keyState)
		player2.aiUpdate(&ball)
		ball.update(&player1, &player2, score1, score2)

		player1.draw(array_of_pixels)
		player2.draw(array_of_pixels)
		ball.draw(array_of_pixels)

		tex.Update(nil, array_of_pixels, winWidth*4)
		renderer.Copy(tex, nil, nil)
		renderer.Present()

		sdl.Delay(16)
	}

}
