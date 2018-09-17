package main

import (
	sdl "GoSdlWrapper/sdl_wrapper"
	"time"
)

func main() {

	sdl.Init("SDL WRAPPER TEST", 800, 600)

	defer sdl.Defer_me()

	sdl.Clear()

	sdl.FillCircle(400, 300, 250)

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
