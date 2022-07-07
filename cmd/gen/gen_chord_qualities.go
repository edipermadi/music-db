package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/edipermadi/music-db/pkg/theory/chord"
	"go.uber.org/zap"
)

type chordQualityEntry struct {
	ID              int64
	Name            string
	Cardinality     int
	Number          int
	PitchClass      []int
	IntervalPattern []int
}

var chordQualityEntries []chordQualityEntry

func buildChordQualitiesTableSeed(logger *zap.Logger, writer io.Writer) error {
	logger.Info("generating chord_qualities table seed")

	for i, v := range chord.AllQualities() {
		chordQualityEntries = append(chordQualityEntries, chordQualityEntry{
			ID:              int64(i + 1),
			Name:            v.String(),
			Cardinality:     v.Cardinality(),
			Number:          v.Number(),
			PitchClass:      v.PitchClass(),
			IntervalPattern: v.IntervalPattern(),
		})
	}

	_, _ = fmt.Fprintf(writer, "INSERT INTO chord_qualities (name, cardinality, number, pitch_class, interval_pattern)\nVALUES\n")
	for i, v := range chordQualityEntries {
		encodedPitchClass, _ := json.Marshal(v.PitchClass)
		encodedIntervalPattern, _ := json.Marshal(v.IntervalPattern)
		if i < len(chordQualityEntries)-1 {
			_, _ = fmt.Fprintf(writer, "('%s', %d, %d, '%s', '%s'),\n", v.Name, v.Cardinality, v.Number, encodedPitchClass, encodedIntervalPattern)
		} else {
			_, _ = fmt.Fprintf(writer, "('%s', %d, %d, '%s', '%s');\n\n", v.Name, v.Cardinality, v.Number, encodedPitchClass, encodedIntervalPattern)
		}
	}

	return nil
}

func findChordQualityID(wanted chord.Quality) int64 {
	for _, v := range chordQualityEntries {
		if v.Number == wanted.Number() {
			return v.ID
		}
	}

	return 0
}
