package main

import (
	"fmt"
	"io"

	"github.com/edipermadi/music-db/pkg/theory/chord"
	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"github.com/edipermadi/music-db/pkg/theory/scale"
	"go.uber.org/zap"
)

type scalePitchChordsEntry struct {
	ID      int64
	ScaleID int64
	TonicID int64
	PitchID int64
	ChordID int64
	RootID  int64
}

var scalePitchChordsEntries []scalePitchChordsEntry

func buildScalePitchChords(logger *zap.Logger, writer io.Writer) error {
	logger.Info("generating scale_pitch_chords seed")

	var id int64 = 1
	for _, v := range scale.AllScales() {
		scaleID := findScaleID(v)
		for _, w := range pitch.AllPitches() {
			tonicID := findPitchID(w)
			scalePitches := v.Pitches(w)
			scalePitchesSignature := pitch.Slice(scalePitches).Signature()
			for _, x := range chord.AllChords() {
				chordID := findChordID(x)
				for _, y := range pitch.AllPitches() {
					rootID := findPitchID(y)
					chordPitches := x.Pitches(y)
					chordPitchesSignature := pitch.Slice(chordPitches).Signature()

					if scalePitchesSignature&chordPitchesSignature == chordPitchesSignature {
						scalePitchChordsEntries = append(scalePitchChordsEntries, scalePitchChordsEntry{
							ID:      id,
							ScaleID: scaleID,
							TonicID: tonicID,
							PitchID: rootID,
							ChordID: chordID,
							RootID:  rootID,
						})
					}
				}
			}
		}
	}

	_, _ = fmt.Fprintf(writer, "INSERT INTO scale_pitch_chords (scale_id, tonic_id, pitch_id, chord_id, root_id)\nVALUES\n")
	for i, v := range scalePitchChordsEntries {
		if i < len(scalePitchChordsEntries)-1 {
			_, _ = fmt.Fprintf(writer, "(%d, %d, %d, %d, %d),\n", v.ScaleID, v.TonicID, v.PitchID, v.ChordID, v.RootID)
		} else {
			_, _ = fmt.Fprintf(writer, "(%d, %d, %d, %d, %d);\n\n", v.ScaleID, v.TonicID, v.PitchID, v.ChordID, v.RootID)
		}
	}

	return nil
}
