package main

// Features
//
// - paddle stay on screen
// - ball cannot bounce behind paddle..
//
// - improve gameplay
//   - paddle angle bouncing
//   - Ball start in random direction
//   - Ball velocity increases
// - Game over state
// - two player or PC
// - AI needs to be inperfect
// - resizing of the window

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

//
// Variables
//
const winWidth, winHeight int = 800, 600

// Enum
type gameState int

const (
	start gameState = iota
	play
	gameover
)

var state = start

// ---- end enum

var ballStartVelocity float32 = 250
var paddleSpeed float32 = 250

var nums = [][]byte{
	{
		1, 1, 1,
		1, 0, 1,
		1, 0, 1,
		1, 0, 1,
		1, 1, 1,
	},
	{
		1, 1, 0,
		0, 1, 0,
		0, 1, 0,
		0, 1, 0,
		1, 1, 1,
	},
	{
		1, 1, 1,
		0, 0, 1,
		1, 1, 1,
		1, 0, 0,
		1, 1, 1,
	},
	{
		1, 1, 1,
		0, 0, 1,
		1, 1, 1,
		0, 0, 1,
		1, 1, 1,
	},
	{
		1, 0, 1,
		1, 0, 1,
		1, 1, 1,
		0, 0, 1,
		0, 0, 1,
	},
	{
		1, 1, 1,
		1, 0, 0,
		1, 1, 1,
		0, 0, 1,
		1, 1, 1,
	},
	{
		1, 1, 1,
		1, 0, 0,
		1, 1, 1,
		1, 0, 1,
		1, 1, 1,
	},
	{
		1, 1, 1,
		0, 0, 1,
		0, 0, 1,
		0, 0, 1,
		0, 0, 1,
	},
	{
		1, 1, 1,
		1, 0, 1,
		1, 1, 1,
		1, 0, 1,
		1, 1, 1,
	},
	{
		1, 1, 1,
		1, 0, 1,
		1, 1, 1,
		0, 0, 1,
		1, 1, 1,
	}}

//
// Structures and Functions
//
type color struct {
	r, g, b byte
}

type pos struct {
	x, y float32
}

func getCenter() pos {
	return pos{float32(winWidth / 2), float32(winHeight / 2)}
}

func drawNumber(pos pos, color color, size int, num int, array_of_pixels []byte) {
	startX := int(pos.x) - (size*3)/2
	startY := int(pos.y) - (size*5)/2

	for i, v := range nums[num] {
		if v == 1 {
			for y := startY; y < startY+size; y++ {
				for x := startX; x < startX+size; x++ {
					setPixel(x, y, color, array_of_pixels)
				}
			}
		}
		startX += size
		if (i+1)%3 == 0 {
			startY += size
			startX -= size * 3
		}
	}
}

// lerp: standard helper function to position objects on the screen
func lerp(a float32, b float32, pct float32) float32 {
	return a + pct*(b-a)
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

func (ball *ball) update(leftPaddle *paddle, rightPaddle *paddle, elapsedTime float32) {
	ball.x += ball.xv * elapsedTime
	ball.y += ball.yv * elapsedTime

	// Collisions
	// Collision: Top/bottom screen
	if ball.y-ball.radius < 0 || ball.y+ball.radius > float32(winHeight) {
		// Change direction
		ball.yv = -ball.yv
	}
	// Collision: left/right screen - score
	if ball.x < 0 {
		rightPaddle.score++
		ball.pos = getCenter()
		state = start
	} else if int(ball.x) > winWidth {
		leftPaddle.score++
		ball.pos = getCenter()
		state = start
	}
	// Collision: left paddle
	if ball.x-ball.radius < leftPaddle.x+leftPaddle.w/2 {
		// ball between paddle bottom and paddle top
		if ball.y+ball.radius < leftPaddle.y+leftPaddle.h/2 && ball.y-ball.radius > leftPaddle.y-leftPaddle.h/2 {
			ball.xv = -ball.xv
			// Bugfix1: Ensure the ball bouncens so that the above statement for x is not true again.
			// Bug1: Ball could move inside the paddle if the statement keeps being thue for x.
			ball.x = leftPaddle.x + leftPaddle.w/2.0 + ball.radius
		}
	}
	// Collision: right paddle
	if ball.x+ball.radius > rightPaddle.x-rightPaddle.w/2 {
		// ball between paddle bottom and paddle top
		if ball.y+ball.radius < rightPaddle.y+rightPaddle.h/2 && ball.y-ball.radius > rightPaddle.y-rightPaddle.h/2 {
			ball.xv = -ball.xv
			ball.x = rightPaddle.x - rightPaddle.w/2.0 - ball.radius
		}
	}
}

type paddle struct {
	pos
	w     float32
	h     float32
	speed float32
	score int
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

	numX := lerp(paddle.x, getCenter().x, 0.2)
	drawNumber(pos{numX, 35}, paddle.color, 10, paddle.score, array_of_pixels)
}

//func drawPaddle

func (paddle *paddle) update(keyState []uint8, elapsedTime float32) {
	if keyState[sdl.SCANCODE_UP] != 0 {
		paddle.y -= paddle.speed * elapsedTime
	}
	if keyState[sdl.SCANCODE_DOWN] != 0 {
		paddle.y += paddle.speed * elapsedTime
	}
}

func (paddle *paddle) aiUpdate(ball *ball, elapsedTime float32) {
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

	player1 := paddle{pos{30, float32(winHeight / 2)}, 20, 100, paddleSpeed, 0, color{255, 255, 255}}
	player2 := paddle{pos{float32(winWidth - 30), float32(winHeight / 2)}, 20, 100, paddleSpeed, 0, color{255, 255, 255}}

	ball := ball{getCenter(), 20, ballStartVelocity, ballStartVelocity, color{255, 255, 255}}

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

	var frameStart time.Time
	var elapsedTime float32

	// Game loop
	// Ends with quit event
	for {
		frameStart = time.Now()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		// Update
		if state == play {
			player1.update(keyState, elapsedTime)
			player2.aiUpdate(&ball, elapsedTime)
			ball.update(&player1, &player2, elapsedTime)
		} else if state == start {
			if keyState[sdl.SCANCODE_SPACE] != 0 {
				if player1.score == 3 || player2.score == 3 {
					player1.score = 0
					player2.score = 0
				}
				state = play
			}
		}

		// Draw
		clear(array_of_pixels)
		player1.draw(array_of_pixels)
		player2.draw(array_of_pixels)
		ball.draw(array_of_pixels)

		tex.Update(nil, array_of_pixels, winWidth*4)
		renderer.Copy(tex, nil, nil)
		renderer.Present()

		elapsedTime = float32(time.Since(frameStart).Seconds())
		// Naive framerate
		//sdl.Delay(16)
		// Improved framerate
		if elapsedTime < .005 {
			sdl.Delay(5 - uint32(elapsedTime/1000.0))
			elapsedTime = float32(time.Since(frameStart).Seconds())
		}
	}

}
