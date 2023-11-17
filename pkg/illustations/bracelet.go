package illustations

import (
	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"github.com/fogleman/gg"
	"image"
	"math"
	"slices"
)

func Bracelet(pitches []pitch.Type) image.Image {

	width := 800
	height := 800
	dc := gg.NewContext(width, height)
	centerX := float64(width) / 2
	centerY := float64(height) / 2

	// large outer circle
	dc.DrawCircle(centerX, centerY, 300)
	dc.SetRGB(0, 0, 0)
	dc.Fill()

	// large circle
	dc.DrawCircle(centerX, centerY, 295)
	dc.SetRGB(1, 1, 1)
	dc.Fill()

	step := math.Pi / 6
	for i := 0; i < 12; i++ {
		x := math.Sin(step*float64(i)) * 298
		y := math.Cos(step*float64(i)) * 298

		// small outer circle
		dc.DrawCircle(centerX+x, centerY-y, 60)
		dc.SetRGB(0, 0, 0)
		dc.Fill()

		// check if current number matches the pitch
		match := slices.ContainsFunc(pitches, func(p pitch.Type) bool {
			return i+1 == int(p)
		})

		// small inner circle
		dc.DrawCircle(centerX+x, centerY-y, 55)

		// set color
		if match {
			dc.SetRGB(0.9, 0, 0)
		} else {
			dc.SetRGB(1, 1, 1)
		}

		// fill
		dc.Fill()
	}

	return dc.Image()
}
