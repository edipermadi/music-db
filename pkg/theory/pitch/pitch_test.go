package pitch_test

import (
	"fmt"
	"testing"

	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"github.com/stretchr/testify/assert"
)

func TestPitchType_Value(t *testing.T) {
	type testCase struct {
		Pitch     pitch.Type
		Name      string
		Number    int
		Frequency float64
	}

	testCases := []testCase{
		{pitch.CNatural, "C", 2048, 261.63},
		{pitch.CSharp, "C#", 1024, 277.18},
		{pitch.DNatural, "D", 512, 293.66},
		{pitch.DSharp, "D#", 256, 311.13},
		{pitch.ENatural, "E", 128, 329.63},
		{pitch.FNatural, "F", 64, 349.23},
		{pitch.FSharp, "F#", 32, 369.99},
		{pitch.GNatural, "G", 16, 392.00},
		{pitch.GSharp, "G#", 8, 415.30},
		{pitch.ANatural, "A", 4, 440.00},
		{pitch.ASharp, "A#", 2, 466.16},
		{pitch.BNatural, "B", 1, 493.88},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Name, tc.Pitch.String())
			assert.Equal(t, tc.Number, tc.Pitch.Number())
			assert.Equal(t, tc.Frequency, tc.Pitch.Frequency())
		})
	}
}

func TestType_Transpose(t *testing.T) {
	type testCase struct {
		Expected pitch.Type
		Given    pitch.Type
		Amount   int
	}

	testCases := []testCase{
		{Expected: pitch.CNatural, Given: pitch.CNatural, Amount: -12},
		{Expected: pitch.CSharp, Given: pitch.CNatural, Amount: -11},
		{Expected: pitch.DNatural, Given: pitch.CNatural, Amount: -10},
		{Expected: pitch.DSharp, Given: pitch.CNatural, Amount: -9},
		{Expected: pitch.ENatural, Given: pitch.CNatural, Amount: -8},
		{Expected: pitch.FNatural, Given: pitch.CNatural, Amount: -7},
		{Expected: pitch.FSharp, Given: pitch.CNatural, Amount: -6},
		{Expected: pitch.GNatural, Given: pitch.CNatural, Amount: -5},
		{Expected: pitch.GSharp, Given: pitch.CNatural, Amount: -4},
		{Expected: pitch.ANatural, Given: pitch.CNatural, Amount: -3},
		{Expected: pitch.ASharp, Given: pitch.CNatural, Amount: -2},
		{Expected: pitch.BNatural, Given: pitch.CNatural, Amount: -1},
		{Expected: pitch.CNatural, Given: pitch.CNatural, Amount: 0},
		{Expected: pitch.CSharp, Given: pitch.CNatural, Amount: 1},
		{Expected: pitch.DNatural, Given: pitch.CNatural, Amount: 2},
		{Expected: pitch.DSharp, Given: pitch.CNatural, Amount: 3},
		{Expected: pitch.ENatural, Given: pitch.CNatural, Amount: 4},
		{Expected: pitch.FNatural, Given: pitch.CNatural, Amount: 5},
		{Expected: pitch.FSharp, Given: pitch.CNatural, Amount: 6},
		{Expected: pitch.GNatural, Given: pitch.CNatural, Amount: 7},
		{Expected: pitch.GSharp, Given: pitch.CNatural, Amount: 8},
		{Expected: pitch.ANatural, Given: pitch.CNatural, Amount: 9},
		{Expected: pitch.ASharp, Given: pitch.CNatural, Amount: 10},
		{Expected: pitch.BNatural, Given: pitch.CNatural, Amount: 11},
		{Expected: pitch.CNatural, Given: pitch.CNatural, Amount: 12},
		{Expected: pitch.CSharp, Given: pitch.CSharp, Amount: -12},
		{Expected: pitch.DNatural, Given: pitch.CSharp, Amount: -11},
		{Expected: pitch.DSharp, Given: pitch.CSharp, Amount: -10},
		{Expected: pitch.ENatural, Given: pitch.CSharp, Amount: -9},
		{Expected: pitch.FNatural, Given: pitch.CSharp, Amount: -8},
		{Expected: pitch.FSharp, Given: pitch.CSharp, Amount: -7},
		{Expected: pitch.GNatural, Given: pitch.CSharp, Amount: -6},
		{Expected: pitch.GSharp, Given: pitch.CSharp, Amount: -5},
		{Expected: pitch.ANatural, Given: pitch.CSharp, Amount: -4},
		{Expected: pitch.ASharp, Given: pitch.CSharp, Amount: -3},
		{Expected: pitch.BNatural, Given: pitch.CSharp, Amount: -2},
		{Expected: pitch.CNatural, Given: pitch.CSharp, Amount: -1},
		{Expected: pitch.CSharp, Given: pitch.CSharp, Amount: 0},
		{Expected: pitch.DNatural, Given: pitch.CSharp, Amount: 1},
		{Expected: pitch.DSharp, Given: pitch.CSharp, Amount: 2},
		{Expected: pitch.ENatural, Given: pitch.CSharp, Amount: 3},
		{Expected: pitch.FNatural, Given: pitch.CSharp, Amount: 4},
		{Expected: pitch.FSharp, Given: pitch.CSharp, Amount: 5},
		{Expected: pitch.GNatural, Given: pitch.CSharp, Amount: 6},
		{Expected: pitch.GSharp, Given: pitch.CSharp, Amount: 7},
		{Expected: pitch.ANatural, Given: pitch.CSharp, Amount: 8},
		{Expected: pitch.ASharp, Given: pitch.CSharp, Amount: 9},
		{Expected: pitch.BNatural, Given: pitch.CSharp, Amount: 10},
		{Expected: pitch.CNatural, Given: pitch.CSharp, Amount: 11},
		{Expected: pitch.CSharp, Given: pitch.CSharp, Amount: 12},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%sTransposedBy%d", tc.Given, tc.Amount), func(t *testing.T) {
			assert.Equal(t, tc.Expected, tc.Given.Transpose(tc.Amount))
		})
	}
}

func TestAllPitches(t *testing.T) {
	assert.NotEmpty(t, pitch.AllPitches())
}
