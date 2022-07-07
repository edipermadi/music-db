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
	logger.Info("generating scale_pitches seed")

	var id int64 = 1
	for _, v := range scale.AllScales() {
		for _, w := range pitch.AllPitches() {
			pitchMap := make(map[pitch.Type]struct{})
			pitches := v.Pitches(w)
			for _, p := range pitches {
				pitchMap[p] = struct{}{}
			}

			for _, x := range pitches {
				_, perfect := pitchMap[x.NextFifth()]
				scaleID := findScaleID(v)
				tonicID := findPitchID(w)
				pitchID := findPitchID(x)

				scalePitchEntries = append(scalePitchEntries, scalePitchEntry{
					ID:        id,
					ScaleID:   scaleID,
					TonicID:   tonicID,
					PitchID:   pitchID,
					IsPerfect: perfect,
				})

				id++
			}
		}
	}

	_, _ = fmt.Fprintf(writer, "INSERT INTO scale_pitches (scale_id, tonic_id, pitch_id, is_perfect)\nVALUES\n")
	for i, v := range scalePitchEntries {
		if i < len(scalePitchEntries)-1 {
			_, _ = fmt.Fprintf(writer, "\t(%d, %d, %d, %t),\n", v.ScaleID, v.TonicID, v.PitchID, v.IsPerfect)
		} else {
			_, _ = fmt.Fprintf(writer, "\t(%d, %d, %d, %t);\n\n", v.ScaleID, v.TonicID, v.PitchID, v.IsPerfect)
		}
	}

	return nil
}
