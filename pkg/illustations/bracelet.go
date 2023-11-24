package illustations

import (
	"image"
	"math"

	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"github.com/fogleman/gg"
)

func pitchSliceToMap(pitches []pitch.Type) map[pitch.Type]struct{} {
	pitchMap := make(map[pitch.Type]struct{})
	for _, v := range pitches {
		pitchMap[v] = struct{}{}
	}

	return pitchMap
}

func drawBracelet(pitches []pitch.Type, circle []pitch.Type, labels map[pitch.Type]string) (image.Image, error) {
	width := 400
	height := 400
	dc := gg.NewContext(width, height)
	if err := dc.LoadFontFace("DroidSansFallback.ttf", 24); err != nil {
		return nil, err
	}

	centerX := float64(width) / 2
	centerY := float64(height) / 2

	// large outer circle
	dc.DrawCircle(centerX, centerY, 150)
	dc.SetRGB(0, 0, 0)
	dc.Fill()

	// large circle
	dc.DrawCircle(centerX, centerY, 148)
	dc.SetRGB(1, 1, 1)
	dc.Fill()

	// convert pitches into map
	pitchMap := pitchSliceToMap(pitches)

	step := math.Pi / 6
	for i, p := range circle {
		x := math.Sin(step*float64(i)) * 145
		y := math.Cos(step*float64(i)) * 145

		// small outer circle
		dc.DrawCircle(centerX+x, centerY-y, 32)
		dc.SetRGB(0, 0, 0)
		dc.Fill()

		// check if current number matches the pitch
		_, match := pitchMap[p]

		// small inner circle
		dc.DrawCircle(centerX+x, centerY-y, 30)

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

		// draw text
		dc.SetRGB(0, 0, 0)
		dc.DrawStringAnchored(labels[p], centerX+x, centerY-y, 0.5, 0.5)
	}

	return dc.Image(), nil
}

var pitchClassLabels = map[pitch.Type]string{
	pitch.CNatural: "C",
	pitch.CSharp:   "C♯",
	pitch.DNatural: "D",
	pitch.DSharp:   "D♯",
	pitch.ENatural: "E",
	pitch.FNatural: "F",
	pitch.FSharp:   "F♯",
	pitch.GNatural: "G",
	pitch.GSharp:   "G♯",
	pitch.ANatural: "A",
	pitch.ASharp:   "A♯",
	pitch.BNatural: "B",
}

var circleOfFifthLabels = map[pitch.Type]string{
	pitch.CNatural: "C",
	pitch.CSharp:   "D♭",
	pitch.DNatural: "D",
	pitch.DSharp:   "E♭",
	pitch.ENatural: "E",
	pitch.FNatural: "F",
	pitch.FSharp:   "G♭",
	pitch.GNatural: "G",
	pitch.GSharp:   "A♭",
	pitch.ANatural: "A",
	pitch.ASharp:   "B♭",
	pitch.BNatural: "B",
}

func PitchClassBracelet(pitches []pitch.Type) (image.Image, error) {
	return drawBracelet(pitches, pitch.AllPitches(), pitchClassLabels)
}

func CircleOfFifthBracelet(pitches []pitch.Type) (image.Image, error) {
	return drawBracelet(pitches, pitch.CircleOfFifths(), circleOfFifthLabels)
}
