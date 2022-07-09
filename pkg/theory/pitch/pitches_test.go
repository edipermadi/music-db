package pitch_test

import (
	"testing"

	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"github.com/stretchr/testify/require"
)

func TestSlice_Equal(t *testing.T) {
	a := pitch.Slice{pitch.CNatural, pitch.ENatural, pitch.GNatural}
	b := pitch.Slice{pitch.ENatural, pitch.GNatural, pitch.CNatural}
	require.True(t, a.Equal(b))
}

func TestSlice_Signature(t *testing.T) {
	given := pitch.Slice{pitch.CNatural, pitch.ENatural, pitch.GNatural}
	expected := pitch.CNatural.Number() + pitch.ENatural.Number() + pitch.GNatural.Number()
	require.Equal(t, expected, given.Signature())
}

func TestSlice_Transpose(t *testing.T) {
	given := pitch.Slice{pitch.CNatural, pitch.ENatural, pitch.GNatural}
	expected := pitch.Slice{pitch.DNatural, pitch.FSharp, pitch.ANatural}
	require.Equal(t, expected, given.Transpose(2))
}

func TestSlice_Unique(t *testing.T) {
	given := pitch.Slice{pitch.CNatural, pitch.ENatural, pitch.GNatural, pitch.CNatural}
	expected := pitch.Slice{pitch.CNatural, pitch.ENatural, pitch.GNatural}
	require.Equal(t, expected, given.Unique())
}
