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

	window, err := sdl.CreateWindow("SDL2 example #13", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
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
	dstRect.Y = 10

	for y := 0; y < 4; y++ {
		image.SetAlphaMod(byte(y * 64))

		image.SetBlendMode(sdl.BLENDMODE_NONE)
		dstRect.X = 10
		image.Blit(nil, primarySurface, &dstRect)

		image.SetBlendMode(sdl.BLENDMODE_BLEND)
		dstRect.X += 100
		image.Blit(nil, primarySurface, &dstRect)

		image.SetBlendMode(sdl.BLENDMODE_ADD)
		dstRect.X += 100
		image.Blit(nil, primarySurface, &dstRect)

		image.SetBlendMode(sdl.BLENDMODE_MOD)
		dstRect.X += 100
		image.Blit(nil, primarySurface, &dstRect)

		dstRect.Y += 100
	}

	window.UpdateSurface()

	sdl.Delay(5000)
}
