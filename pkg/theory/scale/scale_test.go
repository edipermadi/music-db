package scale_test

import (
	"fmt"
	"testing"

	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"github.com/edipermadi/music-db/pkg/theory/scale"
	"github.com/stretchr/testify/assert"
)

func TestScaleType_Values(t *testing.T) {
	type testCase struct {
		Scale  scale.Type
		Name   string
		Number int
	}

	testCases := []testCase{
		{scale.InvalidScale, "InvalidScale", 0},
		{scale.Minoric, "Minoric", 2184},

		{scale.Thaptic, "Thaptic", 2193},
		{scale.Lothic, "Lothic", 2328},
		{scale.Phratic, "Phratic", 2244},
		{scale.Aerathic, "Aerathic", 3144},

		{scale.Epathic, "Epathic", 2196},
		{scale.Mynic, "Mynic", 2376},
		{scale.Rothic, "Rothic", 2628},
		{scale.Eporic, "Eporic", 2322},

		{Scale: scale.Zyphic, Name: "Zyphic", Number: 2185},
		{Scale: scale.Epogic, Name: "Epogic", Number: 2200},
		{Scale: scale.Lanic, Name: "Lanic", Number: 2440},
		{Scale: scale.Pyrric, Name: "Pyrric", Number: 3140},

		{Scale: scale.Aeoloric, Name: "Aeoloric", Number: 2188},
		{Scale: scale.Gonic, Name: "Gonic", Number: 2248},
		{Scale: scale.Dalic, Name: "Dalic", Number: 3208},
		{Scale: scale.Dygic, Name: "Dygic", Number: 2321},

		{Scale: scale.Daric, Name: "Daric", Number: 2194},
		{Scale: scale.Lonic, Name: "Lonic", Number: 2344},
		{Scale: scale.Phradic, Name: "Phradic", Number: 2372},
		{Scale: scale.Bolic, Name: "Bolic", Number: 2596},

		{Scale: scale.Saric, Name: "Saric", Number: 2212},
		{Scale: scale.Zoptic, Name: "Zoptic", Number: 2632},
		{Scale: scale.Aeraphic, Name: "Aeraphic", Number: 2338},
		{Scale: scale.Byptic, Name: "Byptic", Number: 2324},

		{Scale: scale.Aeolic, Name: "Aeolic", Number: 2186},
		{Scale: scale.Koptic, Name: "Koptic", Number: 2216},
		{Scale: scale.Mixolyric, Name: "Mixolyric", Number: 2696},
		{Scale: scale.Lydic, Name: "Lydic", Number: 2594},

		{Scale: scale.Stathic, Name: "Stathic", Number: 2210},
		{Scale: scale.Dadic, Name: "Dadic", Number: 2600},

		{Scale: scale.Phrynic, Name: "Phrynic", Number: 2340},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Name, tc.Scale.String())
			assert.Equal(t, tc.Number, tc.Scale.Number())
		})
	}
}

func TestType_Pitches(t *testing.T) {
	type testCase struct {
		Scale   scale.Type
		Tonic   pitch.Type
		Pitches []pitch.Type
	}

	testCases := []testCase{
		{
			Scale:   scale.Minoric,
			Tonic:   pitch.CNatural,
			Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.GSharp},
		},
		{
			Scale:   scale.Minoric,
			Tonic:   pitch.CSharp,
			Pitches: []pitch.Type{pitch.CSharp, pitch.FNatural, pitch.ANatural},
		},
		{
			Scale:   scale.Minoric,
			Tonic:   pitch.DNatural,
			Pitches: []pitch.Type{pitch.DNatural, pitch.FSharp, pitch.ASharp},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Scale%sWithTonic%s", tc.Scale, tc.Tonic), func(t *testing.T) {
			assert.Equal(t, tc.Pitches, tc.Scale.Pitches(tc.Tonic))
		})
	}
}
