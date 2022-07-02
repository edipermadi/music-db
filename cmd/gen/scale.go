package main

import (
	"fmt"
	"io"

	"github.com/edipermadi/music-db/pkg/theory"
	"go.uber.org/zap"
)

type scaleEntry struct {
	ID     int64
	Name   string
	Number int
}

var scaleEntries []scaleEntry

func buildScale(logger *zap.Logger, writer io.Writer) error {
	allScales := []theory.ScaleType{
		theory.Minoric,
		theory.Thaptic,
		theory.Lothic,
		theory.Phratic,
		theory.Aerathic,
		theory.Epathic,
		theory.Mynic,
		theory.Rothic,
		theory.Eporic,
		theory.Zyphic,
		theory.Epogic,
		theory.Lanic,
		theory.Pyrric,
		theory.Aeoloric,
		theory.Gonic,
		theory.Dalic,
		theory.Dygic,
		theory.Daric,
		theory.Lonic,
		theory.Phradic,
		theory.Bolic,
		theory.Saric,
		theory.Zoptic,
		theory.Aeraphic,
		theory.Byptic,
		theory.Aeolic,
		theory.Koptic,
		theory.Mixolyric,
		theory.Lydic,
		theory.Stathic,
		theory.Dadic,
		theory.Phrynic,
	}

	_, _ = fmt.Fprintf(writer, "INSERT INTO scales (name, number)\nVALUES\n")

	max := len(allScales)
	for i, v := range allScales {
		scaleEntries = append(scaleEntries, scaleEntry{int64(i + 1), v.String(), v.Number()})

		if i < max-1 {
			_, _ = fmt.Fprintf(writer, "\t('%s', %d),\n", v.String(), v.Number())
		} else {
			_, _ = fmt.Fprintf(writer, "\t('%s', %d);\n", v.String(), v.Number())
		}
	}

	return nil
}
