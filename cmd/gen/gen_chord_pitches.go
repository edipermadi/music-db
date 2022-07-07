package main

import (
	"fmt"
	"io"

	"go.uber.org/zap"
)

func buildChordPitchesTableSeed(logger *zap.Logger, writer io.Writer) error {
	logger.Info("generating chord_pitches table seed")

	_, _ = fmt.Fprintf(writer, "INSERT INTO chord_pitches (chord_id, pitch_id)\nVALUES\n")
	for i, v := range chordEntries {
		for j, w := range v.Pitches {
			pitchID := findPitchID(w)
			if i == len(chordEntries)-1 && j == len(v.Pitches)-1 {
				_, _ = fmt.Fprintf(writer, "(%d, %d);\n\n", v.ID, pitchID)
			} else {
				_, _ = fmt.Fprintf(writer, "(%d, %d),\n", v.ID, pitchID)
			}
		}
	}

	return nil
}
