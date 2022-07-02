package theory_test

import (
	"testing"

	"github.com/edipermadi/music-db/pkg/theory"
	"github.com/stretchr/testify/assert"
)

func TestScaleType_Values(t *testing.T) {
	type testCase struct {
		Scale  theory.ScaleType
		Name   string
		Number int
	}

	testCases := []testCase{
		{theory.InvalidScale, "InvalidScale", 0},
		{theory.Minoric, "Minoric", 2184},

		{theory.Thaptic, "Thaptic", 2193},
		{theory.Lothic, "Lothic", 2328},
		{theory.Phratic, "Phratic", 2244},
		{theory.Aerathic, "Aerathic", 3144},

		{theory.Epathic, "Epathic", 2196},
		{theory.Mynic, "Mynic", 2376},
		{theory.Rothic, "Rothic", 2628},
		{theory.Eporic, "Eporic", 2322},

		{theory.Zyphic, "Zyphic", 2185},
		{theory.Epogic, "Epogic", 2200},
		{theory.Lanic, "Lanic", 2440},
		{theory.Pyrric, "Pyrric", 3140},

		{theory.Aeoloric, "Aeoloric", 2188},
		{theory.Gonic, "Gonic", 2248},
		{theory.Dalic, "Dalic", 3208},
		{theory.Dygic, "Dygic", 2321},

		{theory.Daric, "Daric", 2194},
		{theory.Lonic, "Lonic", 2344},
		{theory.Phradic, "Phradic", 2372},
		{theory.Bolic, "Bolic", 2596},

		{theory.Saric, "Saric", 2212},
		{theory.Zoptic, "Zoptic", 2632},
		{theory.Aeraphic, "Aeraphic", 2338},
		{theory.Byptic, "Byptic", 2324},

		{theory.Aeolic, "Aeolic", 2186},
		{theory.Koptic, "Koptic", 2216},
		{theory.Mixolyric, "Mixolyric", 2696},
		{theory.Lydic, "Lydic", 2594},

		{theory.Stathic, "Stathic", 2210},
		{theory.Dadic, "Dadic", 2600},

		{theory.Phrynic, "Phrynic", 2340},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Name, tc.Scale.String())
			assert.Equal(t, tc.Number, tc.Scale.Number())
		})
	}
}
