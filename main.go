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

	// sdl.PutString(0, 0, "WOLOLO")

	sdl.Flush()

	time.Sleep(5000 * time.Millisecond)
}
