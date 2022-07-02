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

	file, err := os.OpenFile(path.Clean(outFile), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		logger.With(zap.String("file", outFile)).Fatal("failed to open output file")
	}

	defer func() { _ = file.Close() }()
	if err := buildScale(logger, file); err != nil {
		logger.With(zap.String("file", outFile)).Fatal("failed to build scale seed")
	}
}
