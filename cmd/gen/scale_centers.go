package main

import (
	"fmt"
	"io"

	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"github.com/edipermadi/music-db/pkg/theory/scale"
	"go.uber.org/zap"
)

type scaleCenterEntry struct {
	ID          int64
	ScaleID     int64
	TonicID     int64
	Balanced    bool
	CoordinateX float64
	CoordinateY float64
}

var scaleCenterEntries []scaleCenterEntry

func buildScaleCenters(logger *zap.Logger, writer io.Writer) error {
	allScales := scale.AllScales()
	allPitches := pitch.AllPitches()

	logger.Info("generating scale_centers seed")

	var id int64 = 1
	_, _ = fmt.Fprintf(writer, "INSERT INTO scale_centers (scale_id, tonic_id, balanced, coordinate_x, coordinate_y)\nVALUES\n")
	for i, v := range allScales {
		scaleID := scaleID(v)
		balanced := v.Balanced()
		for j, w := range allPitches {
			tonicID := pitchID(w)
			x, y := v.CenterOfGravity(w)
			scaleCenterEntries = append(scaleCenterEntries, scaleCenterEntry{ID: id, ScaleID: scaleID, TonicID: tonicID, Balanced: balanced, CoordinateX: x, CoordinateY: y})

			if i == len(allScales)-1 && j == len(allPitches)-1 {
				_, _ = fmt.Fprintf(writer, "\t(%d, %d, %t, %.4f, %.4f);\n\n", scaleID, tonicID, balanced, x, y)
			} else {
				_, _ = fmt.Fprintf(writer, "\t(%d, %d, %t, %.4f, %.4f),\n", scaleID, tonicID, balanced, x, y)
			}

			id++
		}
	}

	return nil
}
