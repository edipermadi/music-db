package main

import (
	"fmt"
	"io"

	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"github.com/edipermadi/music-db/pkg/theory/scale"
	"go.uber.org/zap"
)

type keyEntry struct {
	ID       int64
	Scale    scale.Type
	Tonic    pitch.Type
	ScaleID  int64
	TonicID  int64
	Name     string
	Number   int
	Balanced bool
	CenterX  float64
	CenterY  float64
	Pitches  []pitch.Type
}

var keyEntries []keyEntry

func buildKeysTableSeed(logger *zap.Logger, writer io.Writer) error {
	logger.Info("generating keys table seed")

	var id int64 = 1
	for _, v := range scale.AllScales() {
		scaleID := findScaleID(v)
		for _, w := range pitch.AllPitches() {
			centerX, centerY := v.CenterOfGravity(w)
			pitches := v.Pitches(w)
			keyEntries = append(keyEntries, keyEntry{
				ID:       id,
				Scale:    v,
				Tonic:    w,
				ScaleID:  scaleID,
				TonicID:  findPitchID(w),
				Name:     fmt.Sprintf("%s%s", w.String(), v.String()),
				Number:   pitch.Slice(pitches).Signature(),
				Balanced: v.Balanced(),
				CenterX:  centerX,
				CenterY:  centerY,
				Pitches:  pitches,
			})
			id++
		}
	}

	_, _ = fmt.Fprintf(writer, "INSERT INTO keys (scale_id, tonic_id, name, number, balanced, center_x, center_y)\nVALUES\n")
	for i, v := range keyEntries {
		if i < len(keyEntries)-1 {
			_, _ = fmt.Fprintf(writer, "\t(%d, %d, '%s', %d, %t, %.4f, %.4f),\n", v.ScaleID, v.TonicID, v.Name, v.Number, v.Balanced, v.CenterX, v.CenterY)
		} else {
			_, _ = fmt.Fprintf(writer, "\t(%d, %d, '%s', %d, %t, %.4f, %.4f);\n\n", v.ScaleID, v.TonicID, v.Name, v.Number, v.Balanced, v.CenterX, v.CenterY)
		}
	}

	return nil
}
