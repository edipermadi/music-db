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
	ZeitlerNumber            int
	RingNumber               int
	Perfection               int
	Imperfection             int
	PitchClass               []int
	IntervalPattern          []int
	RotationalSymmetric      bool
	RotationalSymmetryLevel  int
	Palindromic              bool
	ReflectionalSymmetric    bool
	ReflectionalSymmetryAxes []int
	Balanced                 bool
	FifthGeneratorRootDegree int
}

var scaleEntries []scaleEntry

func buildScalesTableSeed(logger *zap.Logger, writer io.Writer) error {
	allScales := scale.AllScales()
	max := len(allScales)

	logger.Info("generating scale seed")
	_, _ = fmt.Fprintf(writer, "INSERT INTO scales (name, cardinality, zeitler_number, ring_number, perfection, imperfection, pitch_class, interval_pattern, rotational_symmetric, rotational_symmetry_level, palindromic, reflectional_symmetric, reflectional_symmetry_axes, balanced, fifth_generator_root_degree)\nVALUES\n")
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
		balanced := v.Balanced()
		fifthGenerator := v.FifthGeneratorRoot()
		fifthGeneratorRootDegree := int(fifthGenerator.Root().Degree)
		scaleEntries = append(scaleEntries, scaleEntry{
			ID:                       int64(i + 1),
			Name:                     v.String(),
			Cardinality:              v.Cardinality(),
			ZeitlerNumber:            v.ZeitlerNumber(),
			RingNumber:               v.RingNumber(),
			Perfection:               result.Perfection,
			Imperfection:             result.Imperfection,
			PitchClass:               pitchClass,
			IntervalPattern:          intervalPattern,
			RotationalSymmetric:      rotationalSymmetric,
			RotationalSymmetryLevel:  rotationalSymmetryLevel,
			Palindromic:              palindromic,
			ReflectionalSymmetric:    reflectiveSymmetric,
			ReflectionalSymmetryAxes: reflectiveSymmetryAxes,
			Balanced:                 balanced,
			FifthGeneratorRootDegree: fifthGeneratorRootDegree,
		})

		if i < max-1 {
			_, _ = fmt.Fprintf(writer, "\t('%s', %d, %d, %d, %d, %d, '%s', '%s', %t, %d, %t, %t, '%s', %t, %d),\n", v.String(), v.Cardinality(), v.ZeitlerNumber(), v.RingNumber(), result.Perfection, result.Imperfection, encodedPitchClass, encodedIntervalPattern, rotationalSymmetric, rotationalSymmetryLevel, palindromic, reflectiveSymmetric, encodedReflectiveSymmetryAxes, balanced, fifthGeneratorRootDegree)
		} else {
			_, _ = fmt.Fprintf(writer, "\t('%s', %d, %d, %d, %d, %d, '%s', '%s', %t, %d, %t, %t, '%s', %t, %d);\n\n", v.String(), v.Cardinality(), v.ZeitlerNumber(), v.RingNumber(), result.Perfection, result.Imperfection, encodedPitchClass, encodedIntervalPattern, rotationalSymmetric, rotationalSymmetryLevel, palindromic, reflectiveSymmetric, encodedReflectiveSymmetryAxes, balanced, fifthGeneratorRootDegree)
		}
	}

	return nil
}

func findScaleID(wanted scale.Type) int64 {
	for _, entry := range scaleEntries {
		if entry.ZeitlerNumber == wanted.ZeitlerNumber() {
			return entry.ID
		}
	}

	return 0
}
