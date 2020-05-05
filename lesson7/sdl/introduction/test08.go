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

	window, err := sdl.CreateWindow("SDL2 example #8", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	primarySurface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	tempImage, err := img.Load("test1.bmp")
	if err != nil {
		panic(err)
	}
	defer tempImage.Free()

	convertedImage, err := tempImage.Convert(primarySurface.Format, 0)
	if err != nil {
		panic(err)
	}
	defer convertedImage.Free()

	primarySurface.FillRect(nil, sdl.MapRGB(primarySurface.Format, 192, 255, 192))

	dstRect := sdl.Rect{
		X: width / 3,
		Y: height / 3,
		W: convertedImage.W / 2,
		H: convertedImage.H / 2,
	}

	err = convertedImage.BlitScaled(nil, primarySurface, &dstRect)
	if err != nil {
		panic(err)
	}

	window.UpdateSurface()

	sdl.Delay(5000)
}
