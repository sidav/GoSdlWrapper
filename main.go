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
		sdl.Clear()

		r := uint8(rand.Int31n(256))
		g := uint8(rand.Int31n(256))
		b := uint8(rand.Int31n(256))

		sdl.SetColor(r, g, b)

		x := int32(400) // rand.Int31n(800)
		y := int32(300) // rand.Int31n(600)
		rad := rand.Int31n(500) + 11
		prec := time.Now()
		sdl.DrawPreciseCircle(x, y, rad-10)
		timePrec := time.Since(prec)
		app := time.Now()
		sdl.DrawApproxCircle(x, y, rad, 25)
		sdl.FillApproxCircle(x, y, rad, 52)
		timeapp := time.Since(app)
		sdl.PutString(0, 0, fmt.Sprintf("PREC %d, APP %d, diff %d", timePrec, timeapp, timePrec - timeapp))

		//x1 := int32(100)// rand.Int31n(500)
		//x2 := int32(200) // rand.Int31n(500)
		//x3 := int32(150) // rand.Int31n(500)
		//y1 := int32(100) // rand.Int31n(500)
		//y2 := int32(150) // rand.Int31n(500)
		//y3 := int32(400) // rand.Int31n(500)
		//sdl.FillTriangle(x1, y1, x2, y2, x3, y3)
		//fmt.Printf("%d %d %d %d %d %d \n", x1, y1, x2, y2, x3, y3)

		sdl.Flush()
		sdl.WaitKey()

		total++

		if time.Since(start) > time.Millisecond*1000 && total > 1000 {
			break
		}
	}

	fmt.Print(total)

	sdl.PutString(0, 0, "WOLOLO")

	time.Sleep(1000 * time.Millisecond)
}
