package scale_test

import (
	"testing"

	"github.com/edipermadi/music-db/pkg/theory/scale"
	"github.com/stretchr/testify/assert"
)

func TestScaleWith3Notes_Values(t *testing.T) {
	type testCase struct {
		Scale           scale.Type
		Name            string
		Cardinality     int
		Number          int
		Perfection      int
		Imperfection    int
		PitchClass      []int
		IntervalPattern []int
	}

	testCases := []testCase{
		{Scale: scale.Minoric, Name: "Minoric", Cardinality: 3, Number: 2184, Perfection: 0, Imperfection: 3, PitchClass: []int{0, 4, 8}, IntervalPattern: []int{4, 4, 4}},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Name, tc.Scale.String())
			assert.Equal(t, tc.Cardinality, tc.Scale.Cardinality())
			assert.Equal(t, tc.Number, tc.Scale.Number())
			assert.Equal(t, tc.PitchClass, tc.Scale.PitchClass())
			assert.Equal(t, tc.IntervalPattern, tc.Scale.IntervalPattern())

			result := tc.Scale.Perfection()
			assert.Equal(t, tc.Perfection, result.Perfection)
			assert.Equal(t, tc.Imperfection, result.Imperfection)
		})
	}
}
