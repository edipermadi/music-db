package main

import (
	"fmt"
	"io"

	"go.uber.org/zap"
)

func buildKeyPitchChordsTableSeed(logger *zap.Logger, writer io.Writer) error {
	logger.Info("generating key_pitch_chords seed")

	type entry struct {
		KeyID   int64
		PitchID int64
		ChordID int64
	}

	entries := make([]entry, 0)
	for _, key := range keyEntries {
		for _, chord := range chordEntries {
			if key.Number&chord.Number == chord.Number {
				entries = append(entries, entry{KeyID: key.ID, PitchID: chord.RootID, ChordID: chord.ID})
			}
		}
	}

	_, _ = fmt.Fprintf(writer, "INSERT INTO key_pitch_chords (key_id, pitch_id, chord_id)\nVALUES\n")
	for i, v := range entries {
		if i < len(entries)-1 {
			_, _ = fmt.Fprintf(writer, "(%d, %d, %d),\n", v.KeyID, v.PitchID, v.ChordID)
		} else {
			_, _ = fmt.Fprintf(writer, "(%d, %d, %d);\n\n", v.KeyID, v.PitchID, v.ChordID)
		}
	}

	return nil
}
