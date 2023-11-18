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
	file, err := os.CreateTemp(os.TempDir(), "pitch-class-bracelet.*.png")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = os.Remove(file.Name()) }()

	// draw CNaturalIonian
	require.NoError(t, png.Encode(file, illustations.PitchClassBracelet(scale.Ionian.Pitches(pitch.CNatural))))
}

func TestCircleOfFifthBracelet(t *testing.T) {
	file, err := os.CreateTemp(os.TempDir(), "circle-of-fifth-bracelet.*.png")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = os.Remove(file.Name()) }()

	// draw CNaturalIonian
	require.NoError(t, png.Encode(file, illustations.CircleOfFifthBracelet(scale.Ionian.Pitches(pitch.CNatural))))
}
