package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	width  = 640
	height = 480
)

func main() {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("SDL2 example #12", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	primarySurface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	image, err := img.Load("globe.png")
	if err != nil {
		panic(err)
	}
	defer image.Free()

	primarySurface.FillRect(nil, sdl.MapRGB(primarySurface.Format, 192, 255, 192))

	dstRect := sdl.Rect{}

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			image.SetColorMod(byte(x*64), 255, byte(y*64))
			dstRect.X = int32(10 + x*100)
			dstRect.Y = int32(10 + y*100)

			err = image.Blit(nil, primarySurface, &dstRect)
			if err != nil {
				panic(err)
			}
		}
	}

	window.UpdateSurface()

	sdl.Delay(5000)
}
