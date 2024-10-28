package main

import "github.com/veandco/go-sdl2/sdl"

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

	// Draw paddle
	for y := 0; y < int(paddle.h); y++ {
		for x := 0; x < int(paddle.w); x++ {
			setPixel(startX+x, startY+y, paddle.color, array_of_pixels)
		}
	}

	// Draw score
	numX := lerp(paddle.x, getCenter().x, 0.2)
	drawCharacter(pos{numX, 35}, paddle.color, 10, nums, paddle.score, array_of_pixels)
}

func (paddle *paddle) update1(keyState []uint8, elapsedTime float32) {
	if keyState[sdl.SCANCODE_A] != 0 {
		if paddle.y-paddle.h/2 <= 0 {
			paddle.y = paddle.h / 2
		} else if paddle.y-paddle.h/2 > 0 {
			paddle.y -= paddle.speed * elapsedTime
		}
	}
	if keyState[sdl.SCANCODE_Z] != 0 {
		if paddle.y+paddle.h/2 >= float32(winHeight) {
			paddle.y = float32(winHeight) - paddle.h/2
		} else if paddle.y+paddle.h/2 <= float32(winHeight) {
			paddle.y += paddle.speed * elapsedTime
		}
	}
}

func (paddle *paddle) update2(keyState []uint8, elapsedTime float32) {
	if keyState[sdl.SCANCODE_UP] != 0 {
		if paddle.y-paddle.h/2 <= 0 {
			paddle.y = paddle.h / 2
		} else if paddle.y-paddle.h/2 > 0 {
			paddle.y -= paddle.speed * elapsedTime
		}
	}
	if keyState[sdl.SCANCODE_DOWN] != 0 {
		if paddle.y+paddle.h/2 >= float32(winHeight) {
			paddle.y = float32(winHeight) - paddle.h/2
		} else if paddle.y+paddle.h/2 <= float32(winHeight) {
			paddle.y += paddle.speed * elapsedTime
		}
	}
}

func (paddle *paddle) aiUpdate(ball *ball, elapsedTime float32) {
	paddle.y = ball.y
}
