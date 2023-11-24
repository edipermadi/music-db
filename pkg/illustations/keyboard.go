package illustations

import (
	"image"
	"slices"

	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"github.com/fogleman/gg"
)

func Keyboard(pitches []pitch.Type) (image.Image, error) {
	// https://bootcamp.uxdesign.cc/drawing-a-flat-piano-keyboard-in-illustrator-de07c74a64c6

	width := 450
	height := 250
	dc := gg.NewContext(width, height)
	if err := dc.LoadFontFace("DroidSansFallback.ttf", 16); err != nil {
		return nil, err
	}

	displayed := make(map[pitch.Type]struct{})

	// draw white keys
	{
		x := 0.0
		y := 0.0
		keyWidth := float64(width) / 7
		keyHeight := float64(height)
		padding := (x + keyWidth) / 2
		borderWidth := 3.0
		labels := []string{"C", "D", "E", "F", "G", "A", "B"}
		keyPitches := []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.ENatural, pitch.FNatural, pitch.GNatural, pitch.ANatural, pitch.BNatural}

		for i := 0; i < 7; i++ {
			// draw white key border
			dc.DrawRectangle(x, y, keyWidth, keyHeight)
			dc.SetRGB(0, 0, 0)
			dc.Fill()

			// draw white key
			if i == 6 {
				dc.DrawRectangle(x+borderWidth, y+borderWidth, keyWidth-(borderWidth*2), keyHeight-(borderWidth*2))
			} else {
				dc.DrawRectangle(x+borderWidth, y+borderWidth, keyWidth-borderWidth, keyHeight-(borderWidth*2))
			}

			dc.SetRGB(1, 1, 1)
			dc.Fill()

			// add label
			dc.SetRGB(0, 0, 0)
			dc.DrawStringAnchored(labels[i], x+padding, keyHeight-20, 0.5, 0.5)

			// check matching pitch
			keyPitch := keyPitches[i]
			matchingPitch := slices.ContainsFunc(pitches, func(p pitch.Type) bool {
				return keyPitch == p
			})

			// indicate if pitch is active
			if matchingPitch {
				if _, found := displayed[keyPitch]; !found {
					dc.DrawCircle(x+padding, keyHeight-50, 10)
					if keyPitch == pitches[0] {
						dc.SetHexColor("#2196f3")
					} else {
						dc.SetHexColor("#f44336")
					}
					dc.Fill()

					displayed[keyPitch] = struct{}{}
				}
			}

			// increase for the next key
			x += keyWidth
		}
	}

	// draw black keys
	{
		x := 0.0
		y := 0.0
		keyWidth := float64(width) / 12
		keyHeight := float64(height) * 0.6
		padding := (x + keyWidth) / 2

		displayMap := map[int]string{
			1:  "C♯",
			3:  "D♯",
			6:  "F♯",
			8:  "G♯",
			10: "A♯",
		}

		keyPitches := map[int]pitch.Type{
			1:  pitch.CSharp,
			3:  pitch.DSharp,
			6:  pitch.FSharp,
			8:  pitch.GSharp,
			10: pitch.ASharp,
		}

		for i := 0; i < 12; i++ {
			if label, draw := displayMap[i]; draw {
				// draw black key
				dc.DrawRectangle(x, y, keyWidth, keyHeight)
				dc.SetRGB(0, 0, 0)
				dc.Fill()

				// add label
				dc.SetRGB(1, 1, 1)
				dc.DrawStringAnchored(label, x+padding, keyHeight-20, 0.5, 0.5)

				// check matching pitch
				keyPitch := keyPitches[i]
				matchingPitch := slices.ContainsFunc(pitches, func(p pitch.Type) bool {
					return keyPitch == p
				})

				// indicate if pitch is active
				if matchingPitch {
					if _, found := displayed[keyPitch]; !found {
						dc.DrawCircle(x+padding, keyHeight-50, 10)
						if keyPitch == pitches[0] {
							dc.SetHexColor("#2196f3")
						} else {
							dc.SetHexColor("#f44336")
						}
						dc.Fill()

						// mark as displayed
						displayed[keyPitch] = struct{}{}
					}
				}

			}

			// increase for the next key
			x += keyWidth
		}
	}

	return dc.Image(), nil
}
