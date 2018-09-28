package sdl_wrapper

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"time"
)

var (
	window   sdl.Window
	surface  *sdl.Surface
	renderer *sdl.Renderer
	font     *ttf.Font
)

// system funcs

func Init(title string, w, h int32) {

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	if err := ttf.Init(); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		w, h, sdl.WINDOW_SHOWN)

	if font, err = ttf.OpenFont("test.ttf", 32); err != nil {
		panic(err)
	}

	surface, err = window.GetSurface()
	if err != nil {
		panic(err)
	}

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE)
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Millisecond * 1000)

}

func Defer_me() {
	font.Close()
	sdl.Quit()
	window.Destroy()
	renderer.Destroy()
}

///
///
/////
///////
/////////

// clear/flush funcs

func Flush() {
	renderer.Present()
	// window.UpdateSurface()
}

func Clear() {
	//renderer.Clear()
}

// draw funcs

func SetColor(r, g, b uint8) {
	renderer.SetDrawColor(r, g, b, 255)
}

func DrawLine(x, y, x1, y1 int32) {
	renderer.DrawLine(x, y, x1, y1)
}

func WaitKey() rune {
	break_loop := false
	for !break_loop {
		event := sdl.WaitEvent() // wait here until an event is in the event queue
		switch t := event.(type) {
		case *sdl.KeyboardEvent:
			fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
				t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
			break_loop = true
		}
	}

}

func workEvents() {
	event := sdl.WaitEvent() // wait here until an event is in the event queue
	switch t := event.(type) {
	case *sdl.MouseMotionEvent:
		fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
			t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)
	case *sdl.MouseButtonEvent:
		fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
			t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
	case *sdl.MouseWheelEvent:
		fmt.Printf("[%d ms] MouseWheel\ttype:%d\tid:%d\tx:%d\ty:%d\n",
			t.Timestamp, t.Type, t.Which, t.X, t.Y)
	case *sdl.KeyboardEvent:
		fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
			t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
	}
}

func PutString(x, y int32, str string) {
	var (
		solid *sdl.Surface
		err   error
	)

	if solid, err = font.RenderUTF8Solid(str, sdl.Color{255, 0, 0, 255}); err != nil {
		panic(err)
	}

	solid.GetColorKey()
	if err = solid.Blit(nil, surface, nil); err != nil {
		panic(err)
	}

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
