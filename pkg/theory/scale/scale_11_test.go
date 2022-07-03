package scale_test

import (
	"testing"

	"github.com/edipermadi/music-db/pkg/theory/scale"
	"github.com/stretchr/testify/assert"
)

func TestScaleWith11Notes_Values(t *testing.T) {
	type testCase struct {
		Scale        scale.Type
		Name         string
		Cardinality  int
		Number       int
		Perfection   int
		Imperfection int
	}

	testCases := []testCase{
		// 11.1
		{Scale: scale.Aerycratic, Name: "Aerycratic", Cardinality: 11, Number: 3071, Perfection: 10, Imperfection: 1},
		{Scale: scale.Monatic, Name: "Monatic", Cardinality: 11, Number: 4094, Perfection: 10, Imperfection: 1},
		{Scale: scale.Solatic, Name: "Solatic", Cardinality: 11, Number: 4093, Perfection: 10, Imperfection: 1},
		{Scale: scale.Zylatic, Name: "Zylatic", Cardinality: 11, Number: 4091, Perfection: 10, Imperfection: 1},
		{Scale: scale.Mixolatic, Name: "Mixolatic", Cardinality: 11, Number: 4087, Perfection: 10, Imperfection: 1},
		{Scale: scale.Soratic, Name: "Soratic", Cardinality: 11, Number: 4079, Perfection: 10, Imperfection: 1},
		{Scale: scale.Godatic, Name: "Godatic", Cardinality: 11, Number: 4063, Perfection: 10, Imperfection: 1},
		{Scale: scale.Eptatic, Name: "Eptatic", Cardinality: 11, Number: 4031, Perfection: 10, Imperfection: 1},
		{Scale: scale.Ionatic, Name: "Ionatic", Cardinality: 11, Number: 3967, Perfection: 10, Imperfection: 1},
		{Scale: scale.Aeolatic, Name: "Aeolatic", Cardinality: 11, Number: 3839, Perfection: 10, Imperfection: 1},
		{Scale: scale.Thydatic, Name: "Thydatic", Cardinality: 11, Number: 3583, Perfection: 10, Imperfection: 1},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Name, tc.Scale.String())
			assert.Equal(t, tc.Cardinality, tc.Scale.Cardinality())
			assert.Equal(t, tc.Number, tc.Scale.Number())

			result := tc.Scale.Perfection()
			assert.Equal(t, tc.Perfection, result.Perfection)
			assert.Equal(t, tc.Imperfection, result.Imperfection)
		})
	}
}
