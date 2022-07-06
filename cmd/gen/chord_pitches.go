package main

import (
	"fmt"
	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"io"

	"github.com/edipermadi/music-db/pkg/theory/chord"
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
	id := int64(1)

	allChords := chord.AllChords()
	allPitches := pitch.AllPitches()
	maxAllChords := len(allChords)
	maxAllPitches := len(allPitches)

	_, _ = fmt.Fprintf(writer, "INSERT INTO chord_pitches (chord_id, root_id, pitch_id)\nVALUES\n")
	for i, v := range allChords {
		chordID := chordID(v)
		for j, w := range allPitches {
			rootID := pitchID(w)
			pitches := v.Pitches(w)
			maxPitches := len(pitches)
			for k, x := range pitches {
				pitchID := pitchID(x)
				chordPitchEntries = append(chordPitchEntries, chordPitchEntry{ID: id, ChordID: chordID, RootID: rootID, PitchID: pitchID})
				if i == maxAllChords-1 && j == maxAllPitches-1 && k == maxPitches-1 {
					_, _ = fmt.Fprintf(writer, "(%d, %d, %d);\n\n", chordID, rootID, pitchID)
				} else {
					_, _ = fmt.Fprintf(writer, "(%d, %d, %d),\n", chordID, rootID, pitchID)
				}
				id++
			}
		}
	}

	return nil
}
