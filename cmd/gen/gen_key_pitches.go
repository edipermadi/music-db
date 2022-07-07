package main

import (
	"fmt"
	"io"

	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"go.uber.org/zap"
)

func buildKeyPitchesTableSeed(logger *zap.Logger, writer io.Writer) error {
	logger.Info("generating key_pitches table seed")

	_, _ = fmt.Fprintf(writer, "INSERT INTO key_pitches (key_id, pitch_id, is_perfect)\nVALUES\n")
	for i, v := range keyEntries {
		pitchMap := make(map[pitch.Type]struct{})
		for _, p := range v.Pitches {
			pitchMap[p] = struct{}{}
		}

		for j, w := range v.Pitches {
			pitchID := findPitchID(w)
			_, perfect := pitchMap[w.NextFifth()]
			if i == len(keyEntries)-1 && j == len(v.Pitches)-1 {
				_, _ = fmt.Fprintf(writer, "\t(%d, %d, %t);\n\n", v.ID, pitchID, perfect)
			} else {
				_, _ = fmt.Fprintf(writer, "\t(%d, %d, %t),\n", v.ID, pitchID, perfect)
			}
		}

	}

	return nil
}
