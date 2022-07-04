package scale_test

import (
	"testing"

	"github.com/edipermadi/music-db/pkg/theory/scale"
	"github.com/stretchr/testify/assert"
)

func TestScaleWith4Notes_Values(t *testing.T) {
	type testCase struct {
		Scale        scale.Type
		Name         string
		Cardinality  int
		Number       int
		Perfection   int
		Imperfection int
		PitchClass   []int
	}

	testCases := []testCase{
		// 4.1
		{Scale: scale.Thaptic, Name: "Thaptic", Cardinality: 4, Number: 2193, Perfection: 2, Imperfection: 2, PitchClass: []int{0, 4, 7, 11}},
		{Scale: scale.Lothic, Name: "Lothic", Cardinality: 4, Number: 2328, Perfection: 2, Imperfection: 2, PitchClass: []int{0, 3, 7, 8}},
		{Scale: scale.Phratic, Name: "Phratic", Cardinality: 4, Number: 2244, Perfection: 2, Imperfection: 2, PitchClass: []int{0, 4, 5, 9}},
		{Scale: scale.Aerathic, Name: "Aerathic", Cardinality: 4, Number: 3144, Perfection: 2, Imperfection: 2, PitchClass: []int{0, 1, 5, 8}},

		// 4.2
		{Scale: scale.Epathic, Name: "Epathic", Cardinality: 4, Number: 2196, Perfection: 2, Imperfection: 2, PitchClass: []int{0, 4, 7, 9}},
		{Scale: scale.Mynic, Name: "Mynic", Cardinality: 4, Number: 2376, Perfection: 2, Imperfection: 2, PitchClass: []int{0, 3, 5, 8}},
		{Scale: scale.Rothic, Name: "Rothic", Cardinality: 4, Number: 2628, Perfection: 2, Imperfection: 2, PitchClass: []int{0, 2, 5, 9}},
		{Scale: scale.Eporic, Name: "Eporic", Cardinality: 4, Number: 2322, Perfection: 2, Imperfection: 2, PitchClass: []int{0, 3, 7, 10}},

		// 4.3
		{Scale: scale.Zyphic, Name: "Zyphic", Cardinality: 4, Number: 2185, Perfection: 1, Imperfection: 3, PitchClass: []int{0, 4, 8, 11}},
		{Scale: scale.Epogic, Name: "Epogic", Cardinality: 4, Number: 2200, Perfection: 1, Imperfection: 3, PitchClass: []int{0, 4, 7, 8}},
		{Scale: scale.Lanic, Name: "Lanic", Cardinality: 4, Number: 2440, Perfection: 1, Imperfection: 3, PitchClass: []int{0, 3, 4, 8}},
		{Scale: scale.Pyrric, Name: "Pyrric", Cardinality: 4, Number: 3140, Perfection: 1, Imperfection: 3, PitchClass: []int{0, 1, 5, 9}},

		// 4.4
		{Scale: scale.Aeoloric, Name: "Aeoloric", Cardinality: 4, Number: 2188, Perfection: 1, Imperfection: 3, PitchClass: []int{0, 4, 8, 9}},
		{Scale: scale.Gonic, Name: "Gonic", Cardinality: 4, Number: 2248, Perfection: 1, Imperfection: 3, PitchClass: []int{0, 4, 5, 8}},
		{Scale: scale.Dalic, Name: "Dalic", Cardinality: 4, Number: 3208, Perfection: 1, Imperfection: 3, PitchClass: []int{0, 1, 4, 8}},
		{Scale: scale.Dygic, Name: "Dygic", Cardinality: 4, Number: 2321, Perfection: 1, Imperfection: 3, PitchClass: []int{0, 3, 7, 11}},

		// 4.5
		{Scale: scale.Daric, Name: "Daric", Cardinality: 4, Number: 2194, Perfection: 1, Imperfection: 3, PitchClass: []int{0, 4, 7, 10}},
		{Scale: scale.Lonic, Name: "Lonic", Cardinality: 4, Number: 2344, Perfection: 1, Imperfection: 3, PitchClass: []int{0, 3, 6, 8}},
		{Scale: scale.Phradic, Name: "Phradic", Cardinality: 4, Number: 2372, Perfection: 1, Imperfection: 3, PitchClass: []int{0, 3, 5, 9}},
		{Scale: scale.Bolic, Name: "Bolic", Cardinality: 4, Number: 2596, Perfection: 1, Imperfection: 3, PitchClass: []int{0, 2, 6, 9}},

		// 4.6
		{Scale: scale.Saric, Name: "Saric", Cardinality: 4, Number: 2212, Perfection: 1, Imperfection: 3, PitchClass: []int{0, 4, 6, 9}},
		{Scale: scale.Zoptic, Name: "Zoptic", Cardinality: 4, Number: 2632, Perfection: 1, Imperfection: 3, PitchClass: []int{0, 2, 5, 8}},
		{Scale: scale.Aeraphic, Name: "Aeraphic", Cardinality: 4, Number: 2338, Perfection: 1, Imperfection: 3, PitchClass: []int{0, 3, 6, 10}},
		{Scale: scale.Byptic, Name: "Byptic", Cardinality: 4, Number: 2324, Perfection: 1, Imperfection: 3, PitchClass: []int{0, 3, 7, 9}},

		// 4.7
		{Scale: scale.Aeolic, Name: "Aeolic", Cardinality: 4, Number: 2186, Perfection: 0, Imperfection: 4, PitchClass: []int{0, 4, 8, 10}},
		{Scale: scale.Koptic, Name: "Koptic", Cardinality: 4, Number: 2216, Perfection: 0, Imperfection: 4, PitchClass: []int{0, 4, 6, 8}},
		{Scale: scale.Mixolyric, Name: "Mixolyric", Cardinality: 4, Number: 2696, Perfection: 0, Imperfection: 4, PitchClass: []int{0, 2, 4, 8}},
		{Scale: scale.Lydic, Name: "Lydic", Cardinality: 4, Number: 2594, Perfection: 0, Imperfection: 4, PitchClass: []int{0, 2, 6, 10}},

		// 4.8
		{Scale: scale.Stathic, Name: "Stathic", Cardinality: 4, Number: 2210, Perfection: 0, Imperfection: 4, PitchClass: []int{0, 4, 6, 10}},
		{Scale: scale.Dadic, Name: "Dadic", Cardinality: 4, Number: 2600, Perfection: 0, Imperfection: 4, PitchClass: []int{0, 2, 6, 8}},

		// 4.9
		{Scale: scale.Phrynic, Name: "Phrynic", Cardinality: 4, Number: 2340, Perfection: 0, Imperfection: 4, PitchClass: []int{0, 3, 6, 9}},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Name, tc.Scale.String())
			assert.Equal(t, tc.Cardinality, tc.Scale.Cardinality())
			assert.Equal(t, tc.Number, tc.Scale.Number())
			assert.Equal(t, tc.PitchClass, tc.Scale.PitchClass())

			result := tc.Scale.Perfection()
			assert.Equal(t, tc.Perfection, result.Perfection)
			assert.Equal(t, tc.Imperfection, result.Imperfection)
		})
	}
}
