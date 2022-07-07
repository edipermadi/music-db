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

	// build pitches table seed
	if err := buildPitchesTableSeed(logger, file); err != nil {
		logger.With(zap.String("file", outFile)).Fatal("failed to build pitches table seed")
	}

	// build chord_qualities seed
	if err := buildChordQualitiesTableSeed(logger, file); err != nil {
		logger.With(zap.String("file", outFile)).Fatal("failed to build chord_qualities table seed")
	}

	// build chords table seed
	if err := buildChordsTableSeed(logger, file); err != nil {
		logger.With(zap.String("file", outFile)).Fatal("failed to build chords table seed")
	}

	// build chord_pitches table seed
	if err := buildChordPitchesTableSeed(logger, file); err != nil {
		logger.With(zap.String("file", outFile)).Fatal("failed to build chord_pitches table seed")
	}

	// build scales table seed
	if err := buildScalesTableSeed(logger, file); err != nil {
		logger.With(zap.String("file", outFile)).Fatal("failed to build scale table seed")
	}

	// build keys table seed
	if err := buildKeysTableSeed(logger, file); err != nil {
		logger.With(zap.String("file", outFile)).Fatal("failed to build keys table seed")
	}

	// build key_pitches table seed
	if err := buildKeyPitchesTableSeed(logger, file); err != nil {
		logger.With(zap.String("file", outFile)).Fatal("failed to build key_pitches table seed")
	}

	// build scale pitch chords
	if err := buildKeyPitchChordsTableSeed(logger, file); err != nil {
		logger.With(zap.String("file", outFile)).Fatal("failed to build scale pitch chords seed")
	}
}
