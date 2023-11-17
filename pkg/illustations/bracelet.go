package illustations

import (
	"image"
	"math"

	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"github.com/fogleman/gg"
)

func circleOfFifths() []pitch.Type {
	root := pitch.Invalid
	result := make([]pitch.Type, 12)
	for i := 0; i < 12; i++ {
		if i == 0 {
			root = pitch.CNatural
		} else {
			root = root.NextFifth()
		}

		result[i] = root
	}

	return result
}

func pitchSliceToMap(pitches []pitch.Type) map[pitch.Type]struct{} {
	pitchMap := make(map[pitch.Type]struct{})
	for _, v := range pitches {
		pitchMap[v] = struct{}{}
	}

	return pitchMap
}

func drawBracelet(pitches []pitch.Type, circle []pitch.Type) image.Image {
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

	// convert pitches into map
	pitchMap := pitchSliceToMap(pitches)

	step := math.Pi / 6
	for i, p := range circle {
		x := math.Sin(step*float64(i)) * 298
		y := math.Cos(step*float64(i)) * 298

		// small outer circle
		dc.DrawCircle(centerX+x, centerY-y, 65)
		dc.SetRGB(0, 0, 0)
		dc.Fill()

		// check if current number matches the pitch
		_, match := pitchMap[p]

		// small inner circle
		dc.DrawCircle(centerX+x, centerY-y, 60)

		// set color
		if match {
			_, hasNextFifth := pitchMap[p.NextFifth()]
			_, hasPreviousFifth := pitchMap[p.PreviousFifth()]
			switch {
			case hasPreviousFifth && hasNextFifth:
				dc.SetHexColor("#4caf50")
			case !hasPreviousFifth && hasNextFifth:
				dc.SetHexColor("#2196f3")
			case hasPreviousFifth && !hasNextFifth:
				dc.SetHexColor("#ff9800")
			default:
				dc.SetHexColor("#f44336")
			}
		} else {
			dc.SetRGB(1, 1, 1)
		}

		// fill
		dc.Fill()
	}

	return dc.Image()
}

func PitchClassBracelet(pitches []pitch.Type) image.Image {
	return drawBracelet(pitches, pitch.AllPitches())
}

func CircleOfFifthBracelet(pitches []pitch.Type) image.Image {
	return drawBracelet(pitches, circleOfFifths())
}
