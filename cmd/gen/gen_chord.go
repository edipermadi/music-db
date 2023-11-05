package main

import (
	"fmt"
	"io"

	"github.com/edipermadi/music-db/pkg/theory/chord"
	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"go.uber.org/zap"
)

type chordEntry struct {
	ID             int64
	ChordQuality   chord.Quality
	Root           pitch.Type
	ChordQualityID int64
	RootID         int64
	Name           string
	ZeitlerNumber  int // Numbering according to Willam Zeitler's system
	RingNumber     int // Numbering according to Ian Ring's system
	Pitches        pitch.Slice
}

var chordEntries []chordEntry

func buildChordsTableSeed(logger *zap.Logger, writer io.Writer) error {
	logger.Info("generating chords table seed")

	id := int64(1)
	for _, v := range chord.AllQualities() {
		chordQualityID := findChordQualityID(v)
		for _, w := range pitch.AllPitches() {
			pitches := v.Pitches(w)
			chordEntries = append(chordEntries, chordEntry{
				ID:             id,
				ChordQuality:   v,
				Root:           w,
				ChordQualityID: chordQualityID,
				RootID:         findPitchID(w),
				Name:           fmt.Sprintf("%s%s", w.String(), v.String()),
				ZeitlerNumber:  pitch.Slice(pitches).ZeitlerSignature(),
				RingNumber:     pitch.Slice(pitches).RingSignature(),
				Pitches:        pitches,
			})
			id++
		}
	}

	_, _ = fmt.Fprintf(writer, "INSERT INTO chords (chord_quality_id, root_id, name, zeitler_number, ring_number)\nVALUES\n")
	for i, v := range chordEntries {
		if i < len(chordEntries)-1 {
			_, _ = fmt.Fprintf(writer, "(%d, %d, '%s', %d, %d),\n", v.ChordQualityID, v.RootID, v.Name, v.ZeitlerNumber, v.RingNumber)
		} else {
			_, _ = fmt.Fprintf(writer, "(%d, %d, '%s', %d, %d);\n\n", v.ChordQualityID, v.RootID, v.Name, v.ZeitlerNumber, v.RingNumber)
		}
	}

	return nil
}
