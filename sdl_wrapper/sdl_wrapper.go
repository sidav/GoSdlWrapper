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

func Init(title string, windowW, windowH, renderW, renderH int32) {

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	if err := ttf.Init(); err != nil {
		panic(err)
	}

	//window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
	//	windowW, windowH, sdl.WINDOW_SHOWN)
	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		windowW, windowH, sdl.WINDOW_SHOWN+sdl.WINDOW_RESIZABLE)

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

	renderer.SetLogicalSize(renderW, renderH)

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
		// FillTriangleslope(x0, y0, pointsx[i], pointsy[i], pointsx[indexNext], pointsy[indexNext])
		fillTriangle(x0, y0, pointsx[i], pointsy[i], pointsx[indexNext], pointsy[indexNext])
	}
}

func FillPreciseCircle(x0, y0, radius int32) {
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

// Fill a triangle - slope method
// Original Author: Adafruit Industries
func swapCoords(x1, y1, x2, y2 int32) (int32, int32, int32, int32) {
	return x2, y2, x1, y1
}

func hline(x1, x2, y int32) {
	for x := x1; x < x2; x++ {
		renderer.DrawPoint(x, y)
	}
	DrawLine((x1), (y), (x2), (y))
}
func FillTriangleslope(x0, y0, x1, y1, x2, y2 int32) {
	var a, b, y, last int32
	// Sort coordinates by Y order (y2 >= y1 >= y0)
	if y0 > y1 {
		x0, y0, x1, y1 = swapCoords(x0, y0, x1, y1)
	}
	if y1 > y2 {
		x2, y2, x1, y1 = swapCoords(x2, y2, x1, y1)
	}
	if y0 > y1 {
		x0, y0, x1, y1 = swapCoords(x0, y0, x1, y1)
	}

	if (y0 == y2) { // All on same line case
		a = x0
		b = x0
		if (x1 < a) {
			a = x1
		} else if (x1 > b) {
			b = x1
		}
		if (x2 < a) {
			a = x2
		} else if (x2 > b) {
			b = x2
		}
		hline(a, b, y0)
		return
	}

	dx01 := x1 - x0
	dy01 := y1 - y0
	dx02 := x2 - x0
	dy02 := y2 - y0
	dx12 := x2 - x1
	dy12 := y2 - y1
	var sa, sb int32

	// For upper part of triangle, find scanline crossings for segment
	// 0-1 and 0-2.  If y1=y2 (flat-bottomed triangle), the scanline y
	// is included here (and second loop will be skipped, avoiding a /
	// error there), otherwise scanline y1 is skipped here and handle
	// in the second loop...which also avoids a /0 error here if y0=y
	// (flat-topped triangle)
	if (y1 == y2) {
		last = y1 // Include y1 scanline
	} else {
		last = y1 - 1
	} // Skip it

	for y = y0; y <= last; y++ {
		a = x0 + sa/dy01
		b = x0 + sb/dy02
		sa += dx01
		sb += dx02
		// longhand a = x0 + (x1 - x0) * (y - y0) / (y1 - y0)
		//          b = x0 + (x2 - x0) * (y - y0) / (y2 - y0)
		hline(a, b, y)
	}

	// For lower part of triangle, find scanline crossings for segment
	// 0-2 and 1-2.  This loop is skipped if y1=y2
	sa = dx12 * (y - y1)
	sb = dx02 * (y - y0)
	for ; y <= y2; y++ {
		a = x1 + sa/dy12
		b = x0 + sb/dy02
		sa += dx12
		sb += dx02
		// longhand a = x1 + (x2 - x1) * (y - y1) / (y2 - y1)
		//          b = x0 + (x2 - x0) * (y - y0) / (y2 - y0)
		hline(a, b, y)
	}
}

// Fill a triangle - Bresenham method
// Original from http://www.sunshine2k.de/coding/java/TriangleRasterization/TriangleRasterization.html
func fillTriangle(x1, y1, x2, y2, x3, y3 int32) {
	var t1x, t2x, y, minx, maxx, t1xp, t2xp int32
	changed1 := false;
	changed2 := false;
	var signx1, signx2, dx1, dy1, dx2, dy2 int32
	var e1, e2 int32
	// Sort vertices
	if y1 > y2 {
		x1, y1, x2, y2 = swapCoords(x1, y1, x2, y2)
	}
	if y2 > y3 {
		x3, y3, x2, y2 = swapCoords(x3, y3, x2, y2)
	}
	if y1 > y2 {
		x1, y1, x2, y2 = swapCoords(x1, y1, x2, y2)
	}

	t1x = x2
	t2x = x2
	y = y1 // Starting points

	dx1 = (x2 - x1)
	if (dx1 < 0) {
		dx1 = -dx1;
		signx1 = -1;
	} else {
		signx1 = 1;
	}
	dy1 = (y2 - y1);

	dx2 = (x3 - x1);
	if (dx2 < 0) {
		dx2 = -dx2;
		signx2 = -1;
	} else {
		signx2 = 1;
	}
	dy2 = (y3 - y1);
	if (dy1 > dx1) { // swap values
		dx1, dy1 = dy1, dx1
		changed1 = true;
	}
	if (dy2 > dx2) { // swap values
		dx2, dy2 = dy2, dx2
		changed2 = true;
	}

	e2 = (dx2 >> 1);
	// Flat top, just process the second half
	if !(y1 == y2) {
		e1 = (dx1 >> 1);
		var i int32
		for i = 0; i < dx1; {
			t1xp = 0;
			t2xp = 0;
			if (t1x < t2x) {
				minx = t1x;
				maxx = t2x;
			} else {
				minx = t2x;
				maxx = t1x;
			}
			// process first line until y value is about to change
			for (i < dx1) {
				i++;
				e1 += dy1;
				for (e1 >= dx1) {
					e1 -= dx1;
					if (changed1) {
						t1xp = signx1
					} else {
						goto next1
					}
				}
				if (changed1) {
					break
				} else
				{
					t1x += signx1
				}
			}
			// Move line
		next1:
			// process second line until y value is about to change
			for {
				e2 += dy2;
				for (e2 >= dx2) {
					e2 -= dx2;
					if (changed2) {
						t2xp = signx2
					} else {
						goto next2;
					}
				}
				if (changed2) {
					break
				} else {
					t2x += signx2
				}
			}
		next2:
			if (minx > t1x) {
				minx = t1x
			};
			if (minx > t2x) {
				minx = t2x
			};
			if (maxx < t1x) {
				maxx = t1x
			};
			if (maxx < t2x) {
				maxx = t2x
			};
			hline(minx, maxx, y); // Draw line from min to max points found on the y
			// Now increase y
			if (!changed1) {
				t1x += signx1
			};
			t1x += t1xp;
			if (!changed2) {
				t2x += signx2
			};
			t2x += t2xp;
			y += 1;
			if (y == y3) {
				break
			}
		}
	}
// next:
	// Second half
	dx1 = (x3 - x2);
	if (dx1 < 0) {
		dx1 = -dx1;
		signx1 = -1;
	} else {
		signx1 = 1;
	}
	dy1 = (y3 - y2);
	t1x = x2;
	if (dy1 > dx1) { // swap values
		dy1, dx1 = dx1, dy1
		changed1 = true;
	} else {
		changed1 = false;
	}
	e1 = (dx1 >> 1);
	var i int32
	for i = 0; i <= dx1; i++ {
		t1xp = 0;
		t2xp = 0;
		if (t1x < t2x) {
			minx = t1x;
			maxx = t2x;
		} else {
			minx = t2x;
			maxx = t1x;
		}
		// process first line until y value is about to change
		for (i < dx1) {
			e1 += dy1;
			for (e1 >= dx1) {
				e1 -= dx1;
				if (changed1) {
					t1xp = signx1;
					break;
				} else {
					goto next3;
				}
			}
			if (changed1) {
				break;
			} else          {
				t1x += signx1
			}
			if (i < dx1) {
				i++
			};
		}
	next3:
		// process second line until y value is about to change
		for (t2x != x3) {
			e2 += dy2;
			for (e2 >= dx2) {
				e2 -= dx2;
				if (changed2) {
					t2xp = signx2
				} else {
					goto next4
				}
			}
			if (changed2) {
				break;
			} else {
				t2x += signx2
			};
		}
	next4:
		if (minx > t1x) {
			minx = t1x
		};
		if (minx > t2x) {
			minx = t2x
		};
		if (maxx < t1x) {
			maxx = t1x
		};
		if (maxx < t2x) {
			maxx = t2x
		};
		hline(minx, maxx, y); // Draw line from min to max points found on the y
		// Now increase y
		if (!changed1) {
			t1x += signx1
		};
		t1x += t1xp;
		if (!changed2) {
			t2x += signx2
		};
		t2x += t2xp;
		y += 1;
		if (y > y3) {
			return;
		}
	}
}
