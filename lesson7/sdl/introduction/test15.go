package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	width  = 640
	height = 480
)

func putpixel(surface *sdl.Surface, x int32, y int32, r byte, g byte, b byte) {
	if x >= 0 && x < surface.W && y >= 0 && y < surface.H {
		switch surface.Format.BitsPerPixel {
		case 24:
			index := x*3 + y*surface.Pitch
			pixels := surface.Pixels()
			pixels[index] = b
			pixels[index+1] = g
			pixels[index+2] = r
		case 32:
			index := x*4 + y*surface.Pitch
			pixels := surface.Pixels()
			pixels[index] = b
			pixels[index+1] = g
			pixels[index+2] = r
		}
	}
}

func main() {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("SDL2 example #15", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	primarySurface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	primarySurface.FillRect(nil, sdl.MapRGB(primarySurface.Format, 192, 255, 192))

	dstRect := sdl.Rect{}
	dstRect.Y = 10

	primarySurface.Lock()
	var x, y int32
	for y = 0; y < primarySurface.H; y++ {
		for x = 0; x < primarySurface.W; x++ {
			r := byte(255 * x / primarySurface.W)
			g := byte(128)
			b := byte(255 * y / primarySurface.H)
			putpixel(primarySurface, x, y, r, g, b)
		}
	}
	primarySurface.Unlock()

	window.UpdateSurface()

	sdl.Delay(5000)
}
