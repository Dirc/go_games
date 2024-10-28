package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"time"
	"unsafe"
)

// Variables
const winWidth, winHeight int = 800, 600

// Enum
type gameState int

const (
	start gameState = iota
	play
	gameover
	newgame
)

var state = start

// ---- end enum

var paddleSpeed float32 = 250

const (
	acceleration      float32 = 20
	ballStartVelocity float32 = 250
)

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

	// Give every pixel a color: black
	for y := 0; y < winHeight; y++ {
		for x := 0; x < winWidth; x++ {
			setPixel(x, y, color{0, 0, 0}, array_of_pixels)
		}
	}

	// Draw initial
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
			player1.update1(keyState, elapsedTime)
			//player2.aiUpdate(&ball, elapsedTime)
			player2.update2(keyState, elapsedTime)
			ball.update(&player1, &player2, elapsedTime)
		} else if state == start {
			// reset paddles and ball.velocity
			player1.pos = pos{30, float32(winHeight / 2)}
			player2.pos = pos{float32(winWidth - 30), float32(winHeight / 2)}
			ball.reset(ballStartVelocity)
			// Hit SPACE to start
			if keyState[sdl.SCANCODE_SPACE] != 0 {
				state = play
			}
		} else if state == gameover {
			// Draw GAME OVER
			//charX = lerp(0,winWidth, 0.5)
			clear(array_of_pixels)
			drawGameOver(15, 60, array_of_pixels)

			// Update texture and render
			tex.Update(nil, unsafe.Pointer(&array_of_pixels[0]), winWidth*4)
			renderer.Copy(tex, nil, nil)
			renderer.Present()

			// Hit SPACE to start
			if keyState[sdl.SCANCODE_SPACE] != 0 {
				state = newgame
			}
		} else if state == newgame {
			// reset paddles
			player1.pos = pos{30, float32(winHeight / 2)}
			player2.pos = pos{float32(winWidth - 30), float32(winHeight / 2)}
			player1.score = 0
			player2.score = 0
			// Hit SPACE to start
			if keyState[sdl.SCANCODE_SPACE] != 0 {
				state = play
			}
		}

		// Draw
		clear(array_of_pixels)
		player1.draw(array_of_pixels)
		player2.draw(array_of_pixels)
		ball.draw(array_of_pixels)

		tex.Update(nil, unsafe.Pointer(&array_of_pixels[0]), winWidth*4)
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
