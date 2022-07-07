package main

import (
	"fmt"
	"io"

	"github.com/edipermadi/music-db/pkg/theory/chord"
	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"go.uber.org/zap"
)

type chordPitchEntry struct {
	ID      int64
	ChordID int64
	RootID  int64
	PitchID int64
}

var chordPitchEntries []chordPitchEntry

func buildChordPitches(logger *zap.Logger, writer io.Writer) error {
	logger.Info("generating chord_pitches seed")

	id := int64(1)
	for _, v := range chord.AllChords() {
		chordID := findChordID(v)
		for _, w := range pitch.AllPitches() {
			rootID := findPitchID(w)
			for _, x := range v.Pitches(w) {
				chordPitchEntries = append(chordPitchEntries, chordPitchEntry{ID: id, ChordID: chordID, RootID: rootID, PitchID: findPitchID(x)})
				id++
			}
		}
	}

	_, _ = fmt.Fprintf(writer, "INSERT INTO chord_pitches (chord_id, root_id, pitch_id)\nVALUES\n")
	for i, v := range chordPitchEntries {
		if i < len(chordPitchEntries)-1 {
			_, _ = fmt.Fprintf(writer, "(%d, %d, %d),\n", v.ChordID, v.RootID, v.PitchID)
		} else {
			_, _ = fmt.Fprintf(writer, "(%d, %d, %d);\n\n", v.ChordID, v.RootID, v.PitchID)
		}
	}

	return nil
}
