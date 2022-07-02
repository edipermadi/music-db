package main

import (
	"fmt"
	"io"

	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"github.com/edipermadi/music-db/pkg/theory/scale"
	"go.uber.org/zap"
)

type scalePitchEntry struct {
	ID      int64
	ScaleID int64
	TonicID int64
	PitchID int64
}

var scalePitchEntries []scalePitchEntry

func buildScalePitches(logger *zap.Logger, writer io.Writer) error {
	allScales := scale.AllScales()
	allPitches := pitch.AllPitches()

	var id int64 = 1
	_, _ = fmt.Fprintf(writer, "INSERT INTO scale_pitches (scale_id, tonic_id, pitch_id)\nVALUES\n")
	for i, v := range allScales {
		for j, w := range allPitches {
			logger.Info(fmt.Sprintf("%s - %s", v, w))
			pitches := v.Pitches(w)
			for k, x := range pitches {
				v1, v2, v3 := scaleID(v), pitchID(w), pitchID(x)

				scalePitchEntries = append(scalePitchEntries, scalePitchEntry{ID: id, ScaleID: v1, TonicID: v2, PitchID: v3})

				if i == len(allScales)-1 && j == len(allPitches)-1 && k == len(pitches)-1 {
					_, _ = fmt.Fprintf(writer, "\t(%d, %d, %d);\n\n", v1, v2, v3)
				} else {
					_, _ = fmt.Fprintf(writer, "\t(%d, %d, %d),\n", v1, v2, v3)
				}

				id++
			}
		}
	}

	return nil
}
