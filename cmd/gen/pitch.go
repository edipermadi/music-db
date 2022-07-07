package main

import (
	"fmt"
	"io"

	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"go.uber.org/zap"
)

type pitchEntry struct {
	ID        int64
	Name      string
	Number    int
	Frequency float64
}

var pitchEntries []pitchEntry

func buildPitch(logger *zap.Logger, writer io.Writer) error {
	allPitches := pitch.AllPitches()
	max := len(allPitches)

	logger.Info("generating pitch seed")
	_, _ = fmt.Fprintf(writer, "INSERT INTO pitches (name, number, frequency)\nVALUES\n")
	for i, v := range allPitches {
		pitchEntries = append(pitchEntries, pitchEntry{ID: int64(i + 1), Name: v.String(), Number: v.Number(), Frequency: v.Frequency()})

		if i < max-1 {
			_, _ = fmt.Fprintf(writer, "\t('%s', %d, %.2f),\n", v.String(), v.Number(), v.Frequency())
		} else {
			_, _ = fmt.Fprintf(writer, "\t('%s', %d, %f);\n\n", v.String(), v.Number(), v.Frequency())
		}
	}

	return nil
}

func findPitchID(wanted pitch.Type) int64 {
	for _, entry := range pitchEntries {
		if entry.Number == wanted.Number() {
			return entry.ID
		}
	}

	return 0
}
