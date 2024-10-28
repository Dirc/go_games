package main

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
	},
}

var alphabet = [][]byte{
	{
		1, 1, 1,
		1, 0, 1,
		1, 1, 1,
		1, 0, 1,
		1, 0, 1,
	},
	{
		1, 1, 1,
		1, 0, 0,
		1, 1, 1,
		1, 0, 0,
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
		1, 0, 1,
		1, 1, 1,
		1, 0, 1,
		1, 0, 1,
		1, 0, 1,
	},
	{
		1, 1, 1,
		1, 0, 1,
		1, 0, 1,
		1, 0, 1,
		1, 1, 1,
	},
	{
		1, 1, 1,
		1, 0, 1,
		1, 1, 1,
		1, 1, 0,
		1, 0, 1,
	},
	{
		1, 0, 1,
		1, 0, 1,
		1, 0, 1,
		1, 0, 1,
		0, 1, 0,
	},
}

// Structures and Functions
type color struct {
	r, g, b byte
}

type pos struct {
	x, y float32
}

func getCenter() pos {
	return pos{float32(winWidth / 2), float32(winHeight / 2)}
}

func drawCharacter(pos pos, color color, size int, characters [][]byte, num int, array_of_pixels []byte) {
	startX := int(pos.x) - (size*3)/2
	startY := int(pos.y) - (size*5)/2

	for i, v := range characters[num] {
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

func drawGameOver(size int, space int, array_of_pixels []byte) {
	numY := float32(winHeight) * 0.75

	minX := float32(winWidth) * 0.15
	maxX := float32(winWidth) * 0.85

	charSpacing := float32(size + space)

	drawCharacter(pos{minX, numY}, color{255, 255, 255}, size, alphabet, 2, array_of_pixels)
	drawCharacter(pos{minX + charSpacing*1, numY}, color{255, 255, 255}, size, alphabet, 0, array_of_pixels)
	drawCharacter(pos{minX + charSpacing*2, numY}, color{255, 255, 255}, size, alphabet, 3, array_of_pixels)
	drawCharacter(pos{minX + charSpacing*3, numY}, color{255, 255, 255}, size, alphabet, 1, array_of_pixels)
	drawCharacter(pos{maxX - charSpacing*3, numY}, color{255, 255, 255}, size, alphabet, 4, array_of_pixels)
	drawCharacter(pos{maxX - charSpacing*2, numY}, color{255, 255, 255}, size, alphabet, 6, array_of_pixels)
	drawCharacter(pos{maxX - charSpacing*1, numY}, color{255, 255, 255}, size, alphabet, 1, array_of_pixels)
	drawCharacter(pos{maxX, numY}, color{255, 255, 255}, size, alphabet, 5, array_of_pixels)
}
