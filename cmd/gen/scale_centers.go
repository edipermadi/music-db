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
	logger.Info("generating scale_centers seed")

	var id int64 = 1
	for _, v := range scale.AllScales() {
		scaleID := findScaleID(v)
		balanced := v.Balanced()
		for _, w := range pitch.AllPitches() {
			x, y := v.CenterOfGravity(w)
			scaleCenterEntries = append(scaleCenterEntries, scaleCenterEntry{
				ID:          id,
				ScaleID:     scaleID,
				TonicID:     findPitchID(w),
				Balanced:    balanced,
				CoordinateX: x, CoordinateY: y,
			})
			id++
		}
	}

	_, _ = fmt.Fprintf(writer, "INSERT INTO scale_centers (scale_id, tonic_id, balanced, coordinate_x, coordinate_y)\nVALUES\n")
	for i, v := range scaleCenterEntries {
		if i < len(scaleCenterEntries)-1 {
			_, _ = fmt.Fprintf(writer, "\t(%d, %d, %t, %.4f, %.4f),\n", v.ScaleID, v.TonicID, v.Balanced, v.CoordinateX, v.CoordinateY)
		} else {
			_, _ = fmt.Fprintf(writer, "\t(%d, %d, %t, %.4f, %.4f);\n\n", v.ScaleID, v.TonicID, v.Balanced, v.CoordinateX, v.CoordinateY)
		}
	}

	return nil
}
