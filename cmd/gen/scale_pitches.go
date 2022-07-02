package main

import (
	"fmt"
	"io"

	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"github.com/edipermadi/music-db/pkg/theory/scale"
	"go.uber.org/zap"
)

type scalePitchEntry struct {
	ID        int64
	ScaleID   int64
	TonicID   int64
	PitchID   int64
	IsPerfect bool
}

var scalePitchEntries []scalePitchEntry

func buildScalePitches(logger *zap.Logger, writer io.Writer) error {
	allScales := scale.AllScales()
	allPitches := pitch.AllPitches()

	logger.Info("generating scale_pitches seed")

	var id int64 = 1
	_, _ = fmt.Fprintf(writer, "INSERT INTO scale_pitches (scale_id, tonic_id, pitch_id, is_perfect)\nVALUES\n")
	for i, v := range allScales {
		for j, w := range allPitches {
			pitchMap := make(map[pitch.Type]struct{})
			pitches := v.Pitches(w)
			for _, p := range pitches {
				pitchMap[p] = struct{}{}
			}

			for k, x := range pitches {
				_, perfect := pitchMap[x.NextFifth()]
				v1, v2, v3 := scaleID(v), pitchID(w), pitchID(x)

				scalePitchEntries = append(scalePitchEntries, scalePitchEntry{ID: id, ScaleID: v1, TonicID: v2, PitchID: v3, IsPerfect: perfect})

				if i == len(allScales)-1 && j == len(allPitches)-1 && k == len(pitches)-1 {
					_, _ = fmt.Fprintf(writer, "\t(%d, %d, %d, %t);\n\n", v1, v2, v3, perfect)
				} else {
					_, _ = fmt.Fprintf(writer, "\t(%d, %d, %d, %t),\n", v1, v2, v3, perfect)
				}

				id++
			}
		}
	}

	return nil
}
