package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/edipermadi/music-db/pkg/theory/scale"
	"go.uber.org/zap"
)

type scaleEntry struct {
	ID                       int64
	Name                     string
	Cardinality              int
	Number                   int
	Perfection               int
	Imperfection             int
	PitchClass               []int
	IntervalPattern          []int
	RotationalSymmetric      bool
	RotationalSymmetryLevel  int
	Palindromic              bool
	ReflectionalSymmetric    bool
	ReflectionalSymmetryAxes []int
}

var scaleEntries []scaleEntry

func buildScale(logger *zap.Logger, writer io.Writer) error {
	allScales := scale.AllScales()
	max := len(allScales)

	logger.Info("generating scale seed")
	_, _ = fmt.Fprintf(writer, "INSERT INTO scales (name, cardinality, number, perfection, imperfection, pitch_class, interval_pattern, rotational_symmetric, rotational_symmetry_level, palindromic, reflectional_symmetric, reflectional_symmetry_axes)\nVALUES\n")
	for i, v := range allScales {
		result := v.Perfection()
		pitchClass := v.PitchClass()
		encodedPitchClass, _ := json.Marshal(pitchClass)
		intervalPattern := v.IntervalPattern()
		encodedIntervalPattern, _ := json.Marshal(intervalPattern)
		rotationalSymmetric := v.RotationalSymmetric()
		rotationalSymmetryLevel := v.RotationalSymmetryLevel()
		palindromic := v.Palindromic()
		reflectiveSymmetric := v.ReflectiveSymmetric()
		reflectiveSymmetryAxes := v.ReflectiveSymmetryAxes()
		encodedReflectiveSymmetryAxes, _ := json.Marshal(reflectiveSymmetryAxes)
		scaleEntries = append(scaleEntries, scaleEntry{
			ID:                       int64(i + 1),
			Name:                     v.String(),
			Cardinality:              v.Cardinality(),
			Number:                   v.Number(),
			Perfection:               result.Perfection,
			Imperfection:             result.Imperfection,
			PitchClass:               pitchClass,
			IntervalPattern:          intervalPattern,
			RotationalSymmetric:      rotationalSymmetric,
			RotationalSymmetryLevel:  rotationalSymmetryLevel,
			Palindromic:              palindromic,
			ReflectionalSymmetric:    reflectiveSymmetric,
			ReflectionalSymmetryAxes: reflectiveSymmetryAxes,
		})

		if i < max-1 {
			_, _ = fmt.Fprintf(writer, "\t('%s', %d, %d, %d, %d, '%s', '%s', %t, %d, %t, %t, '%s'),\n", v.String(), v.Cardinality(), v.Number(), result.Perfection, result.Imperfection, encodedPitchClass, encodedIntervalPattern, rotationalSymmetric, rotationalSymmetryLevel, palindromic, reflectiveSymmetric, encodedReflectiveSymmetryAxes)
		} else {
			_, _ = fmt.Fprintf(writer, "\t('%s', %d, %d, %d, %d, '%s', '%s', %t, %d, %t, %t, '%s');\n\n", v.String(), v.Cardinality(), v.Number(), result.Perfection, result.Imperfection, encodedPitchClass, encodedIntervalPattern, rotationalSymmetric, rotationalSymmetryLevel, palindromic, reflectiveSymmetric, encodedReflectiveSymmetryAxes)
		}
	}

	return nil
}

func scaleID(wanted scale.Type) int64 {
	for _, entry := range scaleEntries {
		if entry.Number == wanted.Number() {
			return entry.ID
		}
	}

	return 0
}
