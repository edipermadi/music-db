package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/edipermadi/music-db/pkg/theory/chord"
	"go.uber.org/zap"
)

type chordEntry struct {
	ID              int64
	Name            string
	Cardinality     int
	Number          int
	PitchClass      []int
	IntervalPattern []int
}

var chordEntries []chordEntry

func buildChord(logger *zap.Logger, writer io.Writer) error {
	id := int64(1)

	allChords := chord.AllChords()

	_, _ = fmt.Fprintf(writer, "INSERT INTO chords (name, cardinality, number, pitch_class, interval_pattern)\nVALUES\n")
	for i, v := range allChords {
		name := v.String()
		cardinality := v.Cardinality()
		number := v.Number()
		pitchClass := v.PitchClass()
		encodedPitchClass, _ := json.Marshal(pitchClass)
		intervalPattern := v.IntervalPattern()
		encodedIntervalPattern, _ := json.Marshal(intervalPattern)
		chordEntries = append(chordEntries, chordEntry{ID: id, Name: name, Cardinality: cardinality, Number: number, PitchClass: pitchClass, IntervalPattern: intervalPattern})
		if i < len(allChords)-1 {
			_, _ = fmt.Fprintf(writer, "('%s', %d, %d, '%s', '%s'),\n", name, cardinality, number, encodedPitchClass, encodedIntervalPattern)
		} else {
			_, _ = fmt.Fprintf(writer, "('%s', %d, %d, '%s', '%s');\n\n", name, cardinality, number, encodedPitchClass, encodedIntervalPattern)
		}
		id++
	}

	return nil
}

func chordID(wanted chord.Type) int64 {
	for _, v := range chordEntries {
		if v.Number == wanted.Number() {
			return v.ID
		}
	}

	return 0
}
