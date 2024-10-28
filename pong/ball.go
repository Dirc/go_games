package main

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

func (ball *ball) reset(ballStartVelocity float32) {
	ball.xv = ballStartVelocity
	ball.yv = ballStartVelocity
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
		state = setState(leftPaddle.score, rightPaddle.score)
	} else if int(ball.x) > winWidth {
		leftPaddle.score++
		ball.pos = getCenter()
		state = setState(leftPaddle.score, rightPaddle.score)
	}
	// Collision: left paddle
	if ball.x-ball.radius < leftPaddle.x+leftPaddle.w/2 {
		// ball between paddle bottom and paddle top
		if ball.y+ball.radius < leftPaddle.y+leftPaddle.h/2 && ball.y-ball.radius > leftPaddle.y-leftPaddle.h/2 {
			ball.xv = -(ball.xv - acceleration)
			// Bugfix1: Ensure the ball bounces so that the above statement for x is not true again.
			// Bug1: Ball could move inside the paddle if the statement keeps being true for x.
			ball.x = leftPaddle.x + leftPaddle.w/2.0 + ball.radius
		}
	}
	// Collision: right paddle
	if ball.x+ball.radius > rightPaddle.x-rightPaddle.w/2 {
		// ball between paddle bottom and paddle top
		if ball.y+ball.radius < rightPaddle.y+rightPaddle.h/2 && ball.y-ball.radius > rightPaddle.y-rightPaddle.h/2 {
			ball.xv = -(ball.xv + acceleration)
			ball.x = rightPaddle.x - rightPaddle.w/2.0 - ball.radius
		}
	}
}
