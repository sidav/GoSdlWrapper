package main

import (
	sdl "GoSdlWrapper/sdl_wrapper"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	sdl.Init("SDL WRAPPER TEST", 800, 600)

	defer sdl.Defer_me()

	sdl.Clear()

	start := time.Now()

	total := 0
	for {

		r := uint8(rand.Int31n(256))
		g := uint8(rand.Int31n(256))
		b := uint8(rand.Int31n(256))

		sdl.SetColor(r, g, b)
		x := rand.Int31n(800)
		y := rand.Int31n(600)
		rad := rand.Int31n(400)
		sdl.FillCircle(x, y, rad)

		total++

		if time.Since(start) > time.Millisecond*1000 {
			break
		}
	}

	fmt.Print(total)

	sdl.Flush()

	time.Sleep(5000 * time.Millisecond)

	//
	//surface.FillRect(nil, 0)
	//
	//rect := sdl.Rect{0, 0, 200, 200}
	//surface.FillRect(&rect, 0xffff0000)
	//window.UpdateSurface()
	//
	//running := true
	//var x, y, w, h int32
	//
	//for running {
	//	x++
	//	y++
	//	w++
	//	h++
	//	surface.FillRect(nil, 0)
	//
	//
	//	rect := sdl.Rect{x, y, w, h}
	//	surface.FillRect(&rect, 0xffff0000)
	//	window.UpdateSurface()
	//	time.Sleep(200 * time.Millisecond)
	//
	//	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
	//		switch event.(type) {
	//		case *sdl.QuitEvent:
	//			println("Quit")
	//			running = false
	//			break
	//		}
	//	}
	//}
}
