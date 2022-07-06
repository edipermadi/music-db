package main

import (
	"flag"
	"os"
	"path"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	defer func() { _ = logger.Sync() }()

	var outFile string
	flag.StringVar(&outFile, "output", "seed.sql", "path to output file")
	flag.Parse()

	file, err := os.OpenFile(path.Clean(outFile), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // #nosec
	if err != nil {
		logger.With(zap.String("file", outFile)).Fatal("failed to open output file")
	}

	defer func() { _ = file.Close() }()

	// build pitch
	if err := buildPitch(logger, file); err != nil {
		logger.With(zap.String("file", outFile)).Fatal("failed to build pitch seed")
	}

	// build chord
	if err := buildChord(logger, file); err != nil {
		logger.With(zap.String("file", outFile)).Fatal("failed to build chord seed")
	}

	// build chord pitches
	if err := buildChordPitches(logger, file); err != nil {
		logger.With(zap.String("file", outFile)).Fatal("failed to build chord pitche seed")
	}

	// build scale
	if err := buildScale(logger, file); err != nil {
		logger.With(zap.String("file", outFile)).Fatal("failed to build scale seed")
	}

	// build scale pitches
	if err := buildScalePitches(logger, file); err != nil {
		logger.With(zap.String("file", outFile)).Fatal("failed to build scale pitches seed")
	}

	// build scale centers
	if err := buildScaleCenters(logger, file); err != nil {
		logger.With(zap.String("file", outFile)).Fatal("failed to build scale centers seed")
	}
}
