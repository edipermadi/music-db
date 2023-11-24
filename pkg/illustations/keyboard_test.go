package illustations_test

import (
	"image/png"
	"os"
	"testing"

	"github.com/edipermadi/music-db/pkg/illustations"
	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"github.com/edipermadi/music-db/pkg/theory/scale"
	"github.com/stretchr/testify/require"
)

func TestKeyboard(t *testing.T) {
	// generate image
	pitches := scale.Ionian.Pitches(pitch.FNatural)
	img, err := illustations.Keyboard(pitches)
	require.NoError(t, err)

	// create temporary file
	file, err := os.CreateTemp(os.TempDir(), "keyboard.*.png")
	require.NoError(t, err)
	defer func() {
		_ = file.Close()
		_ = os.Remove(file.Name())
	}()

	// save image as png
	require.NoError(t, png.Encode(file, img))
}
