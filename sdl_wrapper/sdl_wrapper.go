package sdl_wrapper

import "C"
import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"math"
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

	//window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
	//	w, h, sdl.WINDOW_SHOWN)
	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		w, h, sdl.WINDOW_SHOWN+sdl.WINDOW_RESIZABLE)

	if font, err = ttf.OpenFont("test.ttf", 32); err != nil {
		panic(err)
	}

	surface, err = window.GetSurface()
	if err != nil {
		panic(err)
	}

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE)
	if err != nil {
		renderer, err = window.GetRenderer()
		if err != nil {
			panic(err)
		}
	}

	time.Sleep(time.Millisecond * 10)

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
	SetColor(0, 0, 0)
	renderer.FillRect(&sdl.Rect{0, 0, 800, 600})
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
	return 'a'
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

func hline(x1, x2, y float32) {
	DrawLine(int32(x1), int32(y), int32(x2), int32(y))
}

func swapCoords(x1, y1, x2, y2 float32) (float32, float32, float32, float32) {
	return x2, y2, x1, y1
}

func FillTriangle(x1, y1, x2, y2, x3, y3 int32) {
	hx, hy, mx, my, lx, ly := float32(x1), float32(y1), float32(x2), float32(y2), float32(x3), float32(y3)
	if hy > my {
		hx, hy, mx, my = swapCoords(hx, hy, mx, my)
	}
	if my > ly {
		lx, ly, mx, my = swapCoords(lx, ly, mx, my)
	}
	if hy > my {
		hx, hy, mx, my = swapCoords(hx, hy, mx, my)
	}
	// assuming (hx, hy) as the highest
	x_hl := hx
	x_hm := hx
	if hy == my {
		x_hm = mx
	}
	dx_hl := (hx - lx) / (hy-ly)
	var dx_hm float32
	if hy - my == 0 {
		dx_hm = 0
	} else {
		dx_hm = (hx-mx)/(hy-my)
	}
	for y:=hy; y<=my; y++ {
		hline(x_hl, x_hm, y)
		x_hl += dx_hl
		x_hm += dx_hm
	}
	x_ml := x_hm
	var dx_ml float32
	if my-ly == 0 {
		dx_ml = 0
	} else {
		dx_ml = (mx-lx)/(my-ly)
	}
	for y:=my; y<=ly;y++{
		hline(x_hl, x_ml, y)
		x_hl+= dx_hl
		x_ml += dx_ml
	}
	fmt.Printf("hl %d ml %d hm %d \n", dx_hl, dx_ml, dx_hm)
}

func DrawPreciseCircle(x0, y0, radius int32) { // midpoint circle algorithm. Calculates each point of the circle.
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

func DrawApproxCircle(x0, y0, radius, desiredPointsCount int32) { // draws the circle as a polygon.
	pointsx := make([]int32, desiredPointsCount)
	pointsy := make([]int32, desiredPointsCount)
	anglePerPoint := (2 * 3.14159265358979323) / float64(desiredPointsCount)
	var pointNum int32
	for pointNum = 0; pointNum < desiredPointsCount; pointNum++ {
		pointAngle := anglePerPoint * float64(pointNum)
		x := float64(radius) * math.Sin(pointAngle)
		y := float64(radius) * math.Cos(pointAngle)
		pointsx[pointNum] = int32(x) + x0
		pointsy[pointNum] = int32(y) + y0
	}
	for i := 0; i < len(pointsx); i++ {
		indexNext := int32(i+1) % desiredPointsCount
		DrawLine(pointsx[i], pointsy[i], pointsx[indexNext], pointsy[indexNext])
	}
}

func FillApproxCircle(x0, y0, radius, desiredPointsCount int32) { // draws the circle as a polygon.
	pointsx := make([]int32, desiredPointsCount)
	pointsy := make([]int32, desiredPointsCount)
	anglePerPoint := (2 * 3.14159265358979323) / float64(desiredPointsCount)
	var pointNum int32
	for pointNum = 0; pointNum < desiredPointsCount; pointNum++ {
		pointAngle := anglePerPoint * float64(pointNum)
		x := float64(radius) * math.Sin(pointAngle)
		y := float64(radius) * math.Cos(pointAngle)
		pointsx[pointNum] = int32(x) + x0
		pointsy[pointNum] = int32(y) + y0
	}
	for i := 0; i < len(pointsx); i++ {
		indexNext := int32(i+1) % desiredPointsCount
		FillTriangle(x0, y0, pointsx[i], pointsy[i], pointsx[indexNext], pointsy[indexNext])
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
