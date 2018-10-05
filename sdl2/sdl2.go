package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

//
// Variables
//
const winWidth, winHeight int = 80, 60

var red = color{255, 0, 0}
var green = color{0, 255, 0}
var blue = color{0, 0, 255}
var yellow = color{255, 255, 0}

var palet = [4]color{red, green, blue, yellow}

//
// Structures
//
type color struct {
	r, g, b byte
}

//
// Functions
//
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

func randomColor() color {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(palet))
	//fmt.Println(len(palet), r, palet[r])
	return palet[r]
}

func main() {

	window, err := sdl.CreateWindow("Window Title", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
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
			setPixel(x, y, randomColor(), array_of_pixels)
		}
	}

	tex.Update(nil, array_of_pixels, winWidth*4)
	renderer.Copy(tex, nil, nil)
	renderer.Present()

	sdl.Delay(2000)

}
