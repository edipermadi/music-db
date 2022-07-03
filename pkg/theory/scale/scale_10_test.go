package scale_test

import (
	"testing"

	"github.com/edipermadi/music-db/pkg/theory/scale"
	"github.com/stretchr/testify/assert"
)

func TestScaleWith10Notes_Values(t *testing.T) {
	type testCase struct {
		Scale        scale.Type
		Name         string
		Cardinality  int
		Number       int
		Perfection   int
		Imperfection int
	}

	testCases := []testCase{
		// 10.1
		{Scale: scale.Aerycryllian, Name: "Aerycryllian", Cardinality: 10, Number: 3039, Perfection: 9, Imperfection: 1},
		{Scale: scale.Gadyllian, Name: "Gadyllian", Cardinality: 10, Number: 3966, Perfection: 9, Imperfection: 1},
		{Scale: scale.Solyllian, Name: "Solyllian", Cardinality: 10, Number: 3837, Perfection: 9, Imperfection: 1},
		{Scale: scale.Zyphyllian, Name: "Zyphyllian", Cardinality: 10, Number: 3579, Perfection: 9, Imperfection: 1},
		{Scale: scale.Garyllian, Name: "Garyllian", Cardinality: 10, Number: 3063, Perfection: 9, Imperfection: 1},
		{Scale: scale.Soryllian, Name: "Soryllian", Cardinality: 10, Number: 4062, Perfection: 9, Imperfection: 1},
		{Scale: scale.Godyllian, Name: "Godyllian", Cardinality: 10, Number: 4029, Perfection: 9, Imperfection: 1},
		{Scale: scale.Epityllian, Name: "Epityllian", Cardinality: 10, Number: 3963, Perfection: 9, Imperfection: 1},
		{Scale: scale.Ionyllian, Name: "Ionyllian", Cardinality: 10, Number: 3831, Perfection: 9, Imperfection: 1},
		{Scale: scale.Aeoryllian, Name: "Aeoryllian", Cardinality: 10, Number: 3567, Perfection: 9, Imperfection: 1},

		// 10.2
		{Scale: scale.Katoryllian, Name: "Katoryllian", Cardinality: 10, Number: 2559, Perfection: 8, Imperfection: 2},
		{Scale: scale.Dodyllian, Name: "Dodyllian", Cardinality: 10, Number: 4092, Perfection: 8, Imperfection: 2},
		{Scale: scale.Zogyllian, Name: "Zogyllian", Cardinality: 10, Number: 4089, Perfection: 8, Imperfection: 2},
		{Scale: scale.Madyllian, Name: "Madyllian", Cardinality: 10, Number: 4083, Perfection: 8, Imperfection: 2},
		{Scale: scale.Dycryllian, Name: "Dycryllian", Cardinality: 10, Number: 4071, Perfection: 8, Imperfection: 2},
		{Scale: scale.Aeogyllian, Name: "Aeogyllian", Cardinality: 10, Number: 4047, Perfection: 8, Imperfection: 2},
		{Scale: scale.Dydyllian, Name: "Dydyllian", Cardinality: 10, Number: 3999, Perfection: 8, Imperfection: 2},
		{Scale: scale.Thogyllian, Name: "Thogyllian", Cardinality: 10, Number: 3903, Perfection: 8, Imperfection: 2},
		{Scale: scale.Rygyllian, Name: "Rygyllian", Cardinality: 10, Number: 3711, Perfection: 8, Imperfection: 2},
		{Scale: scale.Bathyllian, Name: "Bathyllian", Cardinality: 10, Number: 3327, Perfection: 8, Imperfection: 2},

		// 10.3
		{Scale: scale.Sydyllian, Name: "Sydyllian", Cardinality: 10, Number: 2815, Perfection: 8, Imperfection: 2},
		{Scale: scale.Katogyllian, Name: "Katogyllian", Cardinality: 10, Number: 3070, Perfection: 8, Imperfection: 2},
		{Scale: scale.Mixodyllian, Name: "Mixodyllian", Cardinality: 10, Number: 4090, Perfection: 8, Imperfection: 2},
		{Scale: scale.Aeradyllian, Name: "Aeradyllian", Cardinality: 10, Number: 4085, Perfection: 8, Imperfection: 2},
		{Scale: scale.Ryptyllian, Name: "Ryptyllian", Cardinality: 10, Number: 4075, Perfection: 8, Imperfection: 2},
		{Scale: scale.Loptyllian, Name: "Loptyllian", Cardinality: 10, Number: 4055, Perfection: 8, Imperfection: 2},
		{Scale: scale.Kataphyllian, Name: "Kataphyllian", Cardinality: 10, Number: 4015, Perfection: 8, Imperfection: 2},
		{Scale: scale.Phradyllian, Name: "Phradyllian", Cardinality: 10, Number: 3935, Perfection: 8, Imperfection: 2},
		{Scale: scale.Dagyllian, Name: "Dagyllian", Cardinality: 10, Number: 3775, Perfection: 8, Imperfection: 2},
		{Scale: scale.Katyllian, Name: "Katyllian", Cardinality: 10, Number: 3455, Perfection: 8, Imperfection: 2},

		// 10.4
		{Scale: scale.Gothyllian, Name: "Gothyllian", Cardinality: 10, Number: 2943, Perfection: 8, Imperfection: 2},
		{Scale: scale.Lythyllian, Name: "Lythyllian", Cardinality: 10, Number: 3582, Perfection: 8, Imperfection: 2},
		{Scale: scale.Bacryllian, Name: "Bacryllian", Cardinality: 10, Number: 3069, Perfection: 8, Imperfection: 2},
		{Scale: scale.Aerygyllian, Name: "Aerygyllian", Cardinality: 10, Number: 4086, Perfection: 8, Imperfection: 2},
		{Scale: scale.Dathyllian, Name: "Dathyllian", Cardinality: 10, Number: 4077, Perfection: 8, Imperfection: 2},
		{Scale: scale.Boptyllian, Name: "Boptyllian", Cardinality: 10, Number: 4059, Perfection: 8, Imperfection: 2},
		{Scale: scale.Bagyllian, Name: "Bagyllian", Cardinality: 10, Number: 4023, Perfection: 8, Imperfection: 2},
		{Scale: scale.Mathyllian, Name: "Mathyllian", Cardinality: 10, Number: 3951, Perfection: 8, Imperfection: 2},
		{Scale: scale.Styptyllian, Name: "Styptyllian", Cardinality: 10, Number: 3807, Perfection: 8, Imperfection: 2},
		{Scale: scale.Zolyllian, Name: "Zolyllian", Cardinality: 10, Number: 3519, Perfection: 8, Imperfection: 2},

		// 10.5
		{Scale: scale.Staptyllian, Name: "Staptyllian", Cardinality: 10, Number: 3007, Perfection: 8, Imperfection: 2},
		{Scale: scale.Danyllian, Name: "Danyllian", Cardinality: 10, Number: 3838, Perfection: 8, Imperfection: 2},
		{Scale: scale.Goptyllian, Name: "Goptyllian", Cardinality: 10, Number: 3581, Perfection: 8, Imperfection: 2},
		{Scale: scale.Epocryllian, Name: "Epocryllian", Cardinality: 10, Number: 3067, Perfection: 8, Imperfection: 2},
		{Scale: scale.Rocryllian, Name: "Rocryllian", Cardinality: 10, Number: 4078, Perfection: 8, Imperfection: 2},
		{Scale: scale.Zyryllian, Name: "Zyryllian", Cardinality: 10, Number: 4061, Perfection: 8, Imperfection: 2},
		{Scale: scale.Sagyllian, Name: "Sagyllian", Cardinality: 10, Number: 4027, Perfection: 8, Imperfection: 2},
		{Scale: scale.Epinyllian, Name: "Epinyllian", Cardinality: 10, Number: 3959, Perfection: 8, Imperfection: 2},
		{Scale: scale.Katagyllian, Name: "Katagyllian", Cardinality: 10, Number: 3823, Perfection: 8, Imperfection: 2},
		{Scale: scale.Ragyllian, Name: "Ragyllian", Cardinality: 10, Number: 3551, Perfection: 8, Imperfection: 2},

		// 10.6
		{Scale: scale.Thydyllian, Name: "Thydyllian", Cardinality: 10, Number: 3055, Perfection: 8, Imperfection: 2},
		{Scale: scale.Epiryllian, Name: "Epiryllian", Cardinality: 10, Number: 4030, Perfection: 8, Imperfection: 2},
		{Scale: scale.Lyryllian, Name: "Lyryllian", Cardinality: 10, Number: 3965, Perfection: 8, Imperfection: 2},
		{Scale: scale.Mogyllian, Name: "Mogyllian", Cardinality: 10, Number: 3835, Perfection: 8, Imperfection: 2},
		{Scale: scale.Katodyllian, Name: "Katodyllian", Cardinality: 10, Number: 3575, Perfection: 8, Imperfection: 2},
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
