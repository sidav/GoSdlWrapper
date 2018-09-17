package sdl_wrapper

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

var (
	window sdl.Window
	// surface *sdl.Surface
	renderer *sdl.Renderer
)

// system funcs

func Init(title string, w, h int32) {

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		w, h, sdl.WINDOW_SHOWN)

	//surface, err = window.GetSurface()
	//if err != nil {
	//	panic(err)
	//}

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		panic(err)
	}

}

func Defer_me() {
	sdl.Quit()
	window.Destroy()
	renderer.Destroy()
}

// clear/flush funcs

func Flush() {
	// window.UpdateSurface()
	renderer.Present()
}

func Clear() {
	renderer.Clear()
}

// draw funcs

func SetColor(r, g, b uint8) {
	renderer.SetDrawColor(r, g, b, 255)
}

func DrawLine(x, y, x1, y1 int32) {
	renderer.DrawLine(x, y, x1, y1)
}

func DrawCircle(x0, y0, radius int32) {
	x := radius - 1
	var (
		y, dx, dy int32
	)
	y = 0
	dx = 1
	dy = 1
	err := dx - (radius << 1)

	for x >= y {
		renderer.DrawPoint(x0+x, y0+y)
		renderer.DrawPoint(x0+y, y0+x)
		renderer.DrawPoint(x0-y, y0+x)
		renderer.DrawPoint(x0-x, y0+y)
		renderer.DrawPoint(x0-x, y0-y)
		renderer.DrawPoint(x0-y, y0-x)
		renderer.DrawPoint(x0+y, y0-x)
		renderer.DrawPoint(x0+x, y0-y)

		if err <= 0 {
			y++
			err += dy
			dy += 2
		}

		if err > 0 {
			x--
			dx += 2
			err += dx - (radius << 1)
		}
	}
}

func FillCircle(x0, y0, radius int32) {
	x := radius - 1
	var (
		y, dx, dy int32
	)
	y = 0
	dx = 1
	dy = 1
	err := dx - (radius << 1)

	for x >= y {
		renderer.DrawLine(x0-y, y0-x, x0+y, y0-x)
		renderer.DrawLine(x0-x, y0-y, x0+x, y0-y)
		renderer.DrawLine(x0-x, y0+y, x0+x, y0+y)
		renderer.DrawLine(x0-y, y0+x, x0+y, y0+x)

		if err <= 0 {
			y++
			err += dy
			dy += 2
		}

		if err > 0 {
			x--
			dx += 2
			err += dx - (radius << 1)
		}
	}

}
