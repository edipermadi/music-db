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
		{pitch.CNatural, "CNatural", 2048, 261.63},
		{pitch.CSharp, "CSharp", 1024, 277.18},
		{pitch.DNatural, "DNatural", 512, 293.66},
		{pitch.DSharp, "DSharp", 256, 311.13},
		{pitch.ENatural, "ENatural", 128, 329.63},
		{pitch.FNatural, "FNatural", 64, 349.23},
		{pitch.FSharp, "FSharp", 32, 369.99},
		{pitch.GNatural, "GNatural", 16, 392.00},
		{pitch.GSharp, "GSharp", 8, 415.30},
		{pitch.ANatural, "ANatural", 4, 440.00},
		{pitch.ASharp, "ASharp", 2, 466.16},
		{pitch.BNatural, "BNatural", 1, 493.88},
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

		{Expected: pitch.DNatural, Given: pitch.DNatural, Amount: -12},
		{Expected: pitch.DSharp, Given: pitch.DNatural, Amount: -11},
		{Expected: pitch.ENatural, Given: pitch.DNatural, Amount: -10},
		{Expected: pitch.FNatural, Given: pitch.DNatural, Amount: -9},
		{Expected: pitch.FSharp, Given: pitch.DNatural, Amount: -8},
		{Expected: pitch.GNatural, Given: pitch.DNatural, Amount: -7},
		{Expected: pitch.GSharp, Given: pitch.DNatural, Amount: -6},
		{Expected: pitch.ANatural, Given: pitch.DNatural, Amount: -5},
		{Expected: pitch.ASharp, Given: pitch.DNatural, Amount: -4},
		{Expected: pitch.BNatural, Given: pitch.DNatural, Amount: -3},
		{Expected: pitch.CNatural, Given: pitch.DNatural, Amount: -2},
		{Expected: pitch.CSharp, Given: pitch.DNatural, Amount: -1},
		{Expected: pitch.DNatural, Given: pitch.DNatural, Amount: 0},
		{Expected: pitch.DSharp, Given: pitch.DNatural, Amount: 1},
		{Expected: pitch.ENatural, Given: pitch.DNatural, Amount: 2},
		{Expected: pitch.FNatural, Given: pitch.DNatural, Amount: 3},
		{Expected: pitch.FSharp, Given: pitch.DNatural, Amount: 4},
		{Expected: pitch.GNatural, Given: pitch.DNatural, Amount: 5},
		{Expected: pitch.GSharp, Given: pitch.DNatural, Amount: 6},
		{Expected: pitch.ANatural, Given: pitch.DNatural, Amount: 7},
		{Expected: pitch.ASharp, Given: pitch.DNatural, Amount: 8},
		{Expected: pitch.BNatural, Given: pitch.DNatural, Amount: 9},
		{Expected: pitch.CNatural, Given: pitch.DNatural, Amount: 10},
		{Expected: pitch.CSharp, Given: pitch.DNatural, Amount: 11},
		{Expected: pitch.DNatural, Given: pitch.DNatural, Amount: 12},

		{Expected: pitch.DSharp, Given: pitch.DSharp, Amount: -12},
		{Expected: pitch.ENatural, Given: pitch.DSharp, Amount: -11},
		{Expected: pitch.FNatural, Given: pitch.DSharp, Amount: -10},
		{Expected: pitch.FSharp, Given: pitch.DSharp, Amount: -9},
		{Expected: pitch.GNatural, Given: pitch.DSharp, Amount: -8},
		{Expected: pitch.GSharp, Given: pitch.DSharp, Amount: -7},
		{Expected: pitch.ANatural, Given: pitch.DSharp, Amount: -6},
		{Expected: pitch.ASharp, Given: pitch.DSharp, Amount: -5},
		{Expected: pitch.BNatural, Given: pitch.DSharp, Amount: -4},
		{Expected: pitch.CNatural, Given: pitch.DSharp, Amount: -3},
		{Expected: pitch.CSharp, Given: pitch.DSharp, Amount: -2},
		{Expected: pitch.DNatural, Given: pitch.DSharp, Amount: -1},
		{Expected: pitch.DSharp, Given: pitch.DSharp, Amount: 0},
		{Expected: pitch.ENatural, Given: pitch.DSharp, Amount: 1},
		{Expected: pitch.FNatural, Given: pitch.DSharp, Amount: 2},
		{Expected: pitch.FSharp, Given: pitch.DSharp, Amount: 3},
		{Expected: pitch.GNatural, Given: pitch.DSharp, Amount: 4},
		{Expected: pitch.GSharp, Given: pitch.DSharp, Amount: 5},
		{Expected: pitch.ANatural, Given: pitch.DSharp, Amount: 6},
		{Expected: pitch.ASharp, Given: pitch.DSharp, Amount: 7},
		{Expected: pitch.BNatural, Given: pitch.DSharp, Amount: 8},
		{Expected: pitch.CNatural, Given: pitch.DSharp, Amount: 9},
		{Expected: pitch.CSharp, Given: pitch.DSharp, Amount: 10},
		{Expected: pitch.DNatural, Given: pitch.DSharp, Amount: 11},
		{Expected: pitch.DSharp, Given: pitch.DSharp, Amount: 12},

		{Expected: pitch.ENatural, Given: pitch.ENatural, Amount: -12},
		{Expected: pitch.FNatural, Given: pitch.ENatural, Amount: -11},
		{Expected: pitch.FSharp, Given: pitch.ENatural, Amount: -10},
		{Expected: pitch.GNatural, Given: pitch.ENatural, Amount: -9},
		{Expected: pitch.GSharp, Given: pitch.ENatural, Amount: -8},
		{Expected: pitch.ANatural, Given: pitch.ENatural, Amount: -7},
		{Expected: pitch.ASharp, Given: pitch.ENatural, Amount: -6},
		{Expected: pitch.BNatural, Given: pitch.ENatural, Amount: -5},
		{Expected: pitch.CNatural, Given: pitch.ENatural, Amount: -4},
		{Expected: pitch.CSharp, Given: pitch.ENatural, Amount: -3},
		{Expected: pitch.DNatural, Given: pitch.ENatural, Amount: -2},
		{Expected: pitch.DSharp, Given: pitch.ENatural, Amount: -1},
		{Expected: pitch.ENatural, Given: pitch.ENatural, Amount: 0},
		{Expected: pitch.FNatural, Given: pitch.ENatural, Amount: 1},
		{Expected: pitch.FSharp, Given: pitch.ENatural, Amount: 2},
		{Expected: pitch.GNatural, Given: pitch.ENatural, Amount: 3},
		{Expected: pitch.GSharp, Given: pitch.ENatural, Amount: 4},
		{Expected: pitch.ANatural, Given: pitch.ENatural, Amount: 5},
		{Expected: pitch.ASharp, Given: pitch.ENatural, Amount: 6},
		{Expected: pitch.BNatural, Given: pitch.ENatural, Amount: 7},
		{Expected: pitch.CNatural, Given: pitch.ENatural, Amount: 8},
		{Expected: pitch.CSharp, Given: pitch.ENatural, Amount: 9},
		{Expected: pitch.DNatural, Given: pitch.ENatural, Amount: 10},
		{Expected: pitch.DSharp, Given: pitch.ENatural, Amount: 11},
		{Expected: pitch.ENatural, Given: pitch.ENatural, Amount: 12},

		{Expected: pitch.FNatural, Given: pitch.FNatural, Amount: -12},
		{Expected: pitch.FSharp, Given: pitch.FNatural, Amount: -11},
		{Expected: pitch.GNatural, Given: pitch.FNatural, Amount: -10},
		{Expected: pitch.GSharp, Given: pitch.FNatural, Amount: -9},
		{Expected: pitch.ANatural, Given: pitch.FNatural, Amount: -8},
		{Expected: pitch.ASharp, Given: pitch.FNatural, Amount: -7},
		{Expected: pitch.BNatural, Given: pitch.FNatural, Amount: -6},
		{Expected: pitch.CNatural, Given: pitch.FNatural, Amount: -5},
		{Expected: pitch.CSharp, Given: pitch.FNatural, Amount: -4},
		{Expected: pitch.DNatural, Given: pitch.FNatural, Amount: -3},
		{Expected: pitch.DSharp, Given: pitch.FNatural, Amount: -2},
		{Expected: pitch.ENatural, Given: pitch.FNatural, Amount: -1},
		{Expected: pitch.FNatural, Given: pitch.FNatural, Amount: 0},
		{Expected: pitch.FSharp, Given: pitch.FNatural, Amount: 1},
		{Expected: pitch.GNatural, Given: pitch.FNatural, Amount: 2},
		{Expected: pitch.GSharp, Given: pitch.FNatural, Amount: 3},
		{Expected: pitch.ANatural, Given: pitch.FNatural, Amount: 4},
		{Expected: pitch.ASharp, Given: pitch.FNatural, Amount: 5},
		{Expected: pitch.BNatural, Given: pitch.FNatural, Amount: 6},
		{Expected: pitch.CNatural, Given: pitch.FNatural, Amount: 7},
		{Expected: pitch.CSharp, Given: pitch.FNatural, Amount: 8},
		{Expected: pitch.DNatural, Given: pitch.FNatural, Amount: 9},
		{Expected: pitch.DSharp, Given: pitch.FNatural, Amount: 10},
		{Expected: pitch.ENatural, Given: pitch.FNatural, Amount: 11},
		{Expected: pitch.FNatural, Given: pitch.FNatural, Amount: 12},

		{Expected: pitch.FSharp, Given: pitch.FSharp, Amount: -12},
		{Expected: pitch.GNatural, Given: pitch.FSharp, Amount: -11},
		{Expected: pitch.GSharp, Given: pitch.FSharp, Amount: -10},
		{Expected: pitch.ANatural, Given: pitch.FSharp, Amount: -9},
		{Expected: pitch.ASharp, Given: pitch.FSharp, Amount: -8},
		{Expected: pitch.BNatural, Given: pitch.FSharp, Amount: -7},
		{Expected: pitch.CNatural, Given: pitch.FSharp, Amount: -6},
		{Expected: pitch.CSharp, Given: pitch.FSharp, Amount: -5},
		{Expected: pitch.DNatural, Given: pitch.FSharp, Amount: -4},
		{Expected: pitch.DSharp, Given: pitch.FSharp, Amount: -3},
		{Expected: pitch.ENatural, Given: pitch.FSharp, Amount: -2},
		{Expected: pitch.FNatural, Given: pitch.FSharp, Amount: -1},
		{Expected: pitch.FSharp, Given: pitch.FSharp, Amount: 0},
		{Expected: pitch.GNatural, Given: pitch.FSharp, Amount: 1},
		{Expected: pitch.GSharp, Given: pitch.FSharp, Amount: 2},
		{Expected: pitch.ANatural, Given: pitch.FSharp, Amount: 3},
		{Expected: pitch.ASharp, Given: pitch.FSharp, Amount: 4},
		{Expected: pitch.BNatural, Given: pitch.FSharp, Amount: 5},
		{Expected: pitch.CNatural, Given: pitch.FSharp, Amount: 6},
		{Expected: pitch.CSharp, Given: pitch.FSharp, Amount: 7},
		{Expected: pitch.DNatural, Given: pitch.FSharp, Amount: 8},
		{Expected: pitch.DSharp, Given: pitch.FSharp, Amount: 9},
		{Expected: pitch.ENatural, Given: pitch.FSharp, Amount: 10},
		{Expected: pitch.FNatural, Given: pitch.FSharp, Amount: 11},
		{Expected: pitch.FSharp, Given: pitch.FSharp, Amount: 12},

		{Expected: pitch.GNatural, Given: pitch.GNatural, Amount: -12},
		{Expected: pitch.GSharp, Given: pitch.GNatural, Amount: -11},
		{Expected: pitch.ANatural, Given: pitch.GNatural, Amount: -10},
		{Expected: pitch.ASharp, Given: pitch.GNatural, Amount: -9},
		{Expected: pitch.BNatural, Given: pitch.GNatural, Amount: -8},
		{Expected: pitch.CNatural, Given: pitch.GNatural, Amount: -7},
		{Expected: pitch.CSharp, Given: pitch.GNatural, Amount: -6},
		{Expected: pitch.DNatural, Given: pitch.GNatural, Amount: -5},
		{Expected: pitch.DSharp, Given: pitch.GNatural, Amount: -4},
		{Expected: pitch.ENatural, Given: pitch.GNatural, Amount: -3},
		{Expected: pitch.FNatural, Given: pitch.GNatural, Amount: -2},
		{Expected: pitch.FSharp, Given: pitch.GNatural, Amount: -1},
		{Expected: pitch.GNatural, Given: pitch.GNatural, Amount: 0},
		{Expected: pitch.GSharp, Given: pitch.GNatural, Amount: 1},
		{Expected: pitch.ANatural, Given: pitch.GNatural, Amount: 2},
		{Expected: pitch.ASharp, Given: pitch.GNatural, Amount: 3},
		{Expected: pitch.BNatural, Given: pitch.GNatural, Amount: 4},
		{Expected: pitch.CNatural, Given: pitch.GNatural, Amount: 5},
		{Expected: pitch.CSharp, Given: pitch.GNatural, Amount: 6},
		{Expected: pitch.DNatural, Given: pitch.GNatural, Amount: 7},
		{Expected: pitch.DSharp, Given: pitch.GNatural, Amount: 8},
		{Expected: pitch.ENatural, Given: pitch.GNatural, Amount: 9},
		{Expected: pitch.FNatural, Given: pitch.GNatural, Amount: 10},
		{Expected: pitch.FSharp, Given: pitch.GNatural, Amount: 11},
		{Expected: pitch.GNatural, Given: pitch.GNatural, Amount: 12},

		{Expected: pitch.GSharp, Given: pitch.GSharp, Amount: -12},
		{Expected: pitch.ANatural, Given: pitch.GSharp, Amount: -11},
		{Expected: pitch.ASharp, Given: pitch.GSharp, Amount: -10},
		{Expected: pitch.BNatural, Given: pitch.GSharp, Amount: -9},
		{Expected: pitch.CNatural, Given: pitch.GSharp, Amount: -8},
		{Expected: pitch.CSharp, Given: pitch.GSharp, Amount: -7},
		{Expected: pitch.DNatural, Given: pitch.GSharp, Amount: -6},
		{Expected: pitch.DSharp, Given: pitch.GSharp, Amount: -5},
		{Expected: pitch.ENatural, Given: pitch.GSharp, Amount: -4},
		{Expected: pitch.FNatural, Given: pitch.GSharp, Amount: -3},
		{Expected: pitch.FSharp, Given: pitch.GSharp, Amount: -2},
		{Expected: pitch.GNatural, Given: pitch.GSharp, Amount: -1},
		{Expected: pitch.GSharp, Given: pitch.GSharp, Amount: 0},
		{Expected: pitch.ANatural, Given: pitch.GSharp, Amount: 1},
		{Expected: pitch.ASharp, Given: pitch.GSharp, Amount: 2},
		{Expected: pitch.BNatural, Given: pitch.GSharp, Amount: 3},
		{Expected: pitch.CNatural, Given: pitch.GSharp, Amount: 4},
		{Expected: pitch.CSharp, Given: pitch.GSharp, Amount: 5},
		{Expected: pitch.DNatural, Given: pitch.GSharp, Amount: 6},
		{Expected: pitch.DSharp, Given: pitch.GSharp, Amount: 7},
		{Expected: pitch.ENatural, Given: pitch.GSharp, Amount: 8},
		{Expected: pitch.FNatural, Given: pitch.GSharp, Amount: 9},
		{Expected: pitch.FSharp, Given: pitch.GSharp, Amount: 10},
		{Expected: pitch.GNatural, Given: pitch.GSharp, Amount: 11},
		{Expected: pitch.GSharp, Given: pitch.GSharp, Amount: 12},

		{Expected: pitch.ANatural, Given: pitch.ANatural, Amount: -12},
		{Expected: pitch.ASharp, Given: pitch.ANatural, Amount: -11},
		{Expected: pitch.BNatural, Given: pitch.ANatural, Amount: -10},
		{Expected: pitch.CNatural, Given: pitch.ANatural, Amount: -9},
		{Expected: pitch.CSharp, Given: pitch.ANatural, Amount: -8},
		{Expected: pitch.DNatural, Given: pitch.ANatural, Amount: -7},
		{Expected: pitch.DSharp, Given: pitch.ANatural, Amount: -6},
		{Expected: pitch.ENatural, Given: pitch.ANatural, Amount: -5},
		{Expected: pitch.FNatural, Given: pitch.ANatural, Amount: -4},
		{Expected: pitch.FSharp, Given: pitch.ANatural, Amount: -3},
		{Expected: pitch.GNatural, Given: pitch.ANatural, Amount: -2},
		{Expected: pitch.GSharp, Given: pitch.ANatural, Amount: -1},
		{Expected: pitch.ANatural, Given: pitch.ANatural, Amount: 0},
		{Expected: pitch.ASharp, Given: pitch.ANatural, Amount: 1},
		{Expected: pitch.BNatural, Given: pitch.ANatural, Amount: 2},
		{Expected: pitch.CNatural, Given: pitch.ANatural, Amount: 3},
		{Expected: pitch.CSharp, Given: pitch.ANatural, Amount: 4},
		{Expected: pitch.DNatural, Given: pitch.ANatural, Amount: 5},
		{Expected: pitch.DSharp, Given: pitch.ANatural, Amount: 6},
		{Expected: pitch.ENatural, Given: pitch.ANatural, Amount: 7},
		{Expected: pitch.FNatural, Given: pitch.ANatural, Amount: 8},
		{Expected: pitch.FSharp, Given: pitch.ANatural, Amount: 9},
		{Expected: pitch.GNatural, Given: pitch.ANatural, Amount: 10},
		{Expected: pitch.GSharp, Given: pitch.ANatural, Amount: 11},
		{Expected: pitch.ANatural, Given: pitch.ANatural, Amount: 12},

		{Expected: pitch.ASharp, Given: pitch.ASharp, Amount: -12},
		{Expected: pitch.BNatural, Given: pitch.ASharp, Amount: -11},
		{Expected: pitch.CNatural, Given: pitch.ASharp, Amount: -10},
		{Expected: pitch.CSharp, Given: pitch.ASharp, Amount: -9},
		{Expected: pitch.DNatural, Given: pitch.ASharp, Amount: -8},
		{Expected: pitch.DSharp, Given: pitch.ASharp, Amount: -7},
		{Expected: pitch.ENatural, Given: pitch.ASharp, Amount: -6},
		{Expected: pitch.FNatural, Given: pitch.ASharp, Amount: -5},
		{Expected: pitch.FSharp, Given: pitch.ASharp, Amount: -4},
		{Expected: pitch.GNatural, Given: pitch.ASharp, Amount: -3},
		{Expected: pitch.GSharp, Given: pitch.ASharp, Amount: -2},
		{Expected: pitch.ANatural, Given: pitch.ASharp, Amount: -1},
		{Expected: pitch.ASharp, Given: pitch.ASharp, Amount: 0},
		{Expected: pitch.BNatural, Given: pitch.ASharp, Amount: 1},
		{Expected: pitch.CNatural, Given: pitch.ASharp, Amount: 2},
		{Expected: pitch.CSharp, Given: pitch.ASharp, Amount: 3},
		{Expected: pitch.DNatural, Given: pitch.ASharp, Amount: 4},
		{Expected: pitch.DSharp, Given: pitch.ASharp, Amount: 5},
		{Expected: pitch.ENatural, Given: pitch.ASharp, Amount: 6},
		{Expected: pitch.FNatural, Given: pitch.ASharp, Amount: 7},
		{Expected: pitch.FSharp, Given: pitch.ASharp, Amount: 8},
		{Expected: pitch.GNatural, Given: pitch.ASharp, Amount: 9},
		{Expected: pitch.GSharp, Given: pitch.ASharp, Amount: 10},
		{Expected: pitch.ANatural, Given: pitch.ASharp, Amount: 11},
		{Expected: pitch.ASharp, Given: pitch.ASharp, Amount: 12},

		{Expected: pitch.BNatural, Given: pitch.BNatural, Amount: -12},
		{Expected: pitch.CNatural, Given: pitch.BNatural, Amount: -11},
		{Expected: pitch.CSharp, Given: pitch.BNatural, Amount: -10},
		{Expected: pitch.DNatural, Given: pitch.BNatural, Amount: -9},
		{Expected: pitch.DSharp, Given: pitch.BNatural, Amount: -8},
		{Expected: pitch.ENatural, Given: pitch.BNatural, Amount: -7},
		{Expected: pitch.FNatural, Given: pitch.BNatural, Amount: -6},
		{Expected: pitch.FSharp, Given: pitch.BNatural, Amount: -5},
		{Expected: pitch.GNatural, Given: pitch.BNatural, Amount: -4},
		{Expected: pitch.GSharp, Given: pitch.BNatural, Amount: -3},
		{Expected: pitch.ANatural, Given: pitch.BNatural, Amount: -2},
		{Expected: pitch.ASharp, Given: pitch.BNatural, Amount: -1},
		{Expected: pitch.BNatural, Given: pitch.BNatural, Amount: 0},
		{Expected: pitch.CNatural, Given: pitch.BNatural, Amount: 1},
		{Expected: pitch.CSharp, Given: pitch.BNatural, Amount: 2},
		{Expected: pitch.DNatural, Given: pitch.BNatural, Amount: 3},
		{Expected: pitch.DSharp, Given: pitch.BNatural, Amount: 4},
		{Expected: pitch.ENatural, Given: pitch.BNatural, Amount: 5},
		{Expected: pitch.FNatural, Given: pitch.BNatural, Amount: 6},
		{Expected: pitch.FSharp, Given: pitch.BNatural, Amount: 7},
		{Expected: pitch.GNatural, Given: pitch.BNatural, Amount: 8},
		{Expected: pitch.GSharp, Given: pitch.BNatural, Amount: 9},
		{Expected: pitch.ANatural, Given: pitch.BNatural, Amount: 10},
		{Expected: pitch.ASharp, Given: pitch.BNatural, Amount: 11},
		{Expected: pitch.BNatural, Given: pitch.BNatural, Amount: 12},
	}

	for _, tc := range testCases {
		var title string
		if tc.Amount < 0 {
			title = fmt.Sprintf("%sTransposedByNegative%d", tc.Given, tc.Amount*-1)
		} else {
			title = fmt.Sprintf("%sTransposedBy%d", tc.Given, tc.Amount)
		}

		t.Run(title, func(t *testing.T) {
			assert.Equal(t, tc.Expected, tc.Given.Transpose(tc.Amount))
		})
	}
}

func TestAllPitches(t *testing.T) {
	assert.NotEmpty(t, pitch.AllPitches())
}

func TestType_NextFifth(t *testing.T) {
	assert.Equal(t, pitch.FNatural, pitch.CNatural.PreviousFifth())
	assert.Equal(t, pitch.ASharp, pitch.FNatural.PreviousFifth())
	assert.Equal(t, pitch.DSharp, pitch.ASharp.PreviousFifth())
	assert.Equal(t, pitch.GSharp, pitch.DSharp.PreviousFifth())
	assert.Equal(t, pitch.CSharp, pitch.GSharp.PreviousFifth())
	assert.Equal(t, pitch.FSharp, pitch.CSharp.PreviousFifth())
	assert.Equal(t, pitch.BNatural, pitch.FSharp.PreviousFifth())
	assert.Equal(t, pitch.ENatural, pitch.BNatural.PreviousFifth())
	assert.Equal(t, pitch.ANatural, pitch.ENatural.PreviousFifth())
	assert.Equal(t, pitch.DNatural, pitch.ANatural.PreviousFifth())
	assert.Equal(t, pitch.GNatural, pitch.DNatural.PreviousFifth())
	assert.Equal(t, pitch.CNatural, pitch.GNatural.PreviousFifth())
}
