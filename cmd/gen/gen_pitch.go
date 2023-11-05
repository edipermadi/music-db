package main

import (
	"fmt"
	"io"

	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"go.uber.org/zap"
)

type pitchEntry struct {
	ID            int64
	Name          string
	ZeitlerNumber int
	RingNumber    int
	Frequency     float64
}

var pitchEntries []pitchEntry

func buildPitchesTableSeed(logger *zap.Logger, writer io.Writer) error {
	allPitches := pitch.AllPitches()
	max := len(allPitches)

	logger.Info("generating pitches table seed")
	_, _ = fmt.Fprintf(writer, "INSERT INTO pitches (name, zeitler_number, ring_number, frequency)\nVALUES\n")
	for i, v := range allPitches {
		pitchEntries = append(pitchEntries, pitchEntry{ID: int64(i + 1), Name: v.String(), ZeitlerNumber: v.ZeitlerNumber(), RingNumber: v.RingNumber(), Frequency: v.Frequency()})

		if i < max-1 {
			_, _ = fmt.Fprintf(writer, "\t('%s', %d, %d, %.2f),\n", v.String(), v.ZeitlerNumber(), v.RingNumber(), v.Frequency())
		} else {
			_, _ = fmt.Fprintf(writer, "\t('%s', %d, %d, %f);\n\n", v.String(), v.ZeitlerNumber(), v.RingNumber(), v.Frequency())
		}
	}

	return nil
}

func findPitchID(wanted pitch.Type) int64 {
	for _, entry := range pitchEntries {
		if entry.ZeitlerNumber == wanted.ZeitlerNumber() {
			return entry.ID
		}
	}

	return 0
}
