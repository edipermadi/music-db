package illustations_test

import (
	"image/png"
	"log"
	"os"
	"testing"

	"github.com/edipermadi/music-db/pkg/illustations"
	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"github.com/edipermadi/music-db/pkg/theory/scale"
	"github.com/stretchr/testify/require"
)

func TestPitchClassBracelet(t *testing.T) {
	// generate image
	pitches := scale.Ionian.Pitches(pitch.CNatural)
	img, err := illustations.PitchClassBracelet(pitches)
	require.NoError(t, err)

	// create temporary file
	file, err := os.CreateTemp(os.TempDir(), "pitch-class-bracelet.*.png")
	if err != nil {
		log.Fatal(err)
	}

	// close and delete file later
	defer func() {
		_ = file.Close()
		_ = os.Remove(file.Name())
	}()

	// save image as png
	require.NoError(t, png.Encode(file, img))
}

func TestCircleOfFifthBracelet(t *testing.T) {
	// generate image
	pitches := scale.Ionian.Pitches(pitch.CNatural)
	img, err := illustations.CircleOfFifthBracelet(pitches)
	require.NoError(t, err)

	// create temporary file
	file, err := os.CreateTemp(os.TempDir(), "circle-of-fifth-bracelet.*.png")
	if err != nil {
		log.Fatal(err)
	}

	// close and delete file later
	defer func() {
		_ = file.Close()
		_ = os.Remove(file.Name())
	}()

	// save image as png
	require.NoError(t, png.Encode(file, img))
}
