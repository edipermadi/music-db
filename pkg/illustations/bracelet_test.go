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

func TestBracelet(t *testing.T) {
	file, err := os.CreateTemp(os.TempDir(), "bracelet.*.png")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = os.Remove(file.Name()) }()

	// draw CNaturalIonian
	require.NoError(t, png.Encode(file, illustations.Bracelet(scale.Ionian.Pitches(pitch.CNatural))))
}
