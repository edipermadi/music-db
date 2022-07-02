package main

import (
	"fmt"
	"io"

	"github.com/edipermadi/music-db/pkg/theory/scale"
	"go.uber.org/zap"
)

type scaleEntry struct {
	ID           int64
	Name         string
	Number       int
	Perfection   int
	Imperfection int
}

var scaleEntries []scaleEntry

func buildScale(logger *zap.Logger, writer io.Writer) error {
	allScales := scale.AllScales()
	max := len(allScales)

	logger.Info("generating scale seed")
	_, _ = fmt.Fprintf(writer, "INSERT INTO scales (name, number, perfection, imperfection)\nVALUES\n")
	for i, v := range allScales {
		result := v.Perfection()
		scaleEntries = append(scaleEntries, scaleEntry{ID: int64(i + 1), Name: v.String(), Number: v.Number(), Perfection: result.Perfection, Imperfection: result.Imperfection})

		if i < max-1 {
			_, _ = fmt.Fprintf(writer, "\t('%s', %d, %d, %d),\n", v.String(), v.Number(), result.Perfection, result.Imperfection)
		} else {
			_, _ = fmt.Fprintf(writer, "\t('%s', %d, %d, %d);\n\n", v.String(), v.Number(), result.Perfection, result.Imperfection)
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
