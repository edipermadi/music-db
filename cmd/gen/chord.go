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
	logger.Info("generating chords seed")

	for i, v := range chord.AllChords() {
		id := int64(i + 1)
		name := v.String()
		cardinality := v.Cardinality()
		number := v.Number()
		pitchClass := v.PitchClass()
		intervalPattern := v.IntervalPattern()
		chordEntries = append(chordEntries, chordEntry{ID: id, Name: name, Cardinality: cardinality, Number: number, PitchClass: pitchClass, IntervalPattern: intervalPattern})
	}

	_, _ = fmt.Fprintf(writer, "INSERT INTO chords (name, cardinality, number, pitch_class, interval_pattern)\nVALUES\n")
	for i, v := range chordEntries {
		encodedPitchClass, _ := json.Marshal(v.PitchClass)
		encodedIntervalPattern, _ := json.Marshal(v.IntervalPattern)
		if i < len(chordEntries)-1 {
			_, _ = fmt.Fprintf(writer, "('%s', %d, %d, '%s', '%s'),\n", v.Name, v.Cardinality, v.Number, encodedPitchClass, encodedIntervalPattern)
		} else {
			_, _ = fmt.Fprintf(writer, "('%s', %d, %d, '%s', '%s');\n\n", v.Name, v.Cardinality, v.Number, encodedPitchClass, encodedIntervalPattern)
		}
	}

	return nil
}

func findChordID(wanted chord.Type) int64 {
	for _, v := range chordEntries {
		if v.Number == wanted.Number() {
			return v.ID
		}
	}

	return 0
}
