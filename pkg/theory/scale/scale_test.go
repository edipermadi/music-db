package scale_test

import (
	"fmt"
	"testing"

	"github.com/edipermadi/music-db/pkg/theory/degree"
	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"github.com/edipermadi/music-db/pkg/theory/scale"
	"github.com/stretchr/testify/assert"
)

func TestScale_Values(t *testing.T) {
	type testCase struct {
		Scale        scale.Type
		Name         string
		Cardinality  int
		Number       int
		Perfection   int
		Imperfection int
	}

	testCases := []testCase{
		{Scale: scale.Invalid, Name: "Invalid"},
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

func TestScale_Pitches(t *testing.T) {
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
		{
			Scale:   scale.Minoric,
			Tonic:   pitch.DSharp,
			Pitches: []pitch.Type{pitch.DSharp, pitch.GNatural, pitch.BNatural},
		},
		{
			Scale:   scale.Minoric,
			Tonic:   pitch.ENatural,
			Pitches: []pitch.Type{pitch.ENatural, pitch.GSharp, pitch.CNatural},
		},
		{
			Scale:   scale.Minoric,
			Tonic:   pitch.FNatural,
			Pitches: []pitch.Type{pitch.FNatural, pitch.ANatural, pitch.CSharp},
		},
		{
			Scale:   scale.Minoric,
			Tonic:   pitch.FSharp,
			Pitches: []pitch.Type{pitch.FSharp, pitch.ASharp, pitch.DNatural},
		},
		{
			Scale:   scale.Minoric,
			Tonic:   pitch.GNatural,
			Pitches: []pitch.Type{pitch.GNatural, pitch.BNatural, pitch.DSharp},
		},
		{
			Scale:   scale.Minoric,
			Tonic:   pitch.GSharp,
			Pitches: []pitch.Type{pitch.GSharp, pitch.CNatural, pitch.ENatural},
		},
		{
			Scale:   scale.Minoric,
			Tonic:   pitch.ANatural,
			Pitches: []pitch.Type{pitch.ANatural, pitch.CSharp, pitch.FNatural},
		},
		{
			Scale:   scale.Minoric,
			Tonic:   pitch.ASharp,
			Pitches: []pitch.Type{pitch.ASharp, pitch.DNatural, pitch.FSharp},
		},
		{
			Scale:   scale.Minoric,
			Tonic:   pitch.BNatural,
			Pitches: []pitch.Type{pitch.BNatural, pitch.DSharp, pitch.GNatural},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Scale%sWithTonic%s", tc.Scale, tc.Tonic), func(t *testing.T) {
			assert.Equal(t, tc.Pitches, tc.Scale.Pitches(tc.Tonic))
		})
	}
}

func TestType_RationalSymmetric(t *testing.T) {
	assert.True(t, scale.Chromatic.RotationalSymmetric())
	assert.True(t, scale.WholeTone.RotationalSymmetric())
	assert.True(t, scale.Phrynic.RotationalSymmetric())
	assert.False(t, scale.Ionian.RotationalSymmetric())
}

func TestAllScales(t *testing.T) {
	assert.NotEmpty(t, scale.AllScales())
}

func TestType_RotationalSymmetryLevel(t *testing.T) {
	assert.Equal(t, 1, scale.Chromatic.RotationalSymmetryLevel())
	assert.Equal(t, 2, scale.WholeTone.RotationalSymmetryLevel())

	assert.Equal(t, 3, scale.Phrynic.RotationalSymmetryLevel())
	assert.Equal(t, 3, scale.MinorDiminished.RotationalSymmetryLevel())
	assert.Equal(t, 3, scale.MajorDiminished.RotationalSymmetryLevel())

	assert.Equal(t, 4, scale.Minoric.RotationalSymmetryLevel())
	assert.Equal(t, 4, scale.Aerythimic.RotationalSymmetryLevel())
	assert.Equal(t, 4, scale.Stynygic.RotationalSymmetryLevel())
	assert.Equal(t, 4, scale.Ionythimic.RotationalSymmetryLevel())
	assert.Equal(t, 4, scale.Zydygic.RotationalSymmetryLevel())
	assert.Equal(t, 4, scale.Phronygic.RotationalSymmetryLevel())

	assert.Equal(t, 6, scale.Dadic.RotationalSymmetryLevel())
	assert.Equal(t, 6, scale.Stadimic.RotationalSymmetryLevel())
	assert.Equal(t, 6, scale.Dodimic.RotationalSymmetryLevel())
	assert.Equal(t, 6, scale.Zyrimic.RotationalSymmetryLevel())
	assert.Equal(t, 6, scale.Katogyllic.RotationalSymmetryLevel())
	assert.Equal(t, 6, scale.Stathic.RotationalSymmetryLevel())
	assert.Equal(t, 6, scale.Stylimic.RotationalSymmetryLevel())
	assert.Equal(t, 6, scale.Epidyllic.RotationalSymmetryLevel())
}

func TestType_ReflectiveSymmetric(t *testing.T) {
	type testCase struct {
		Scale scale.Type
		Axes  []int
	}

	testCases := []testCase{
		{Scale: scale.Minoric, Axes: []int{0, 2, 4, 6, 8, 10}},
		{Scale: scale.Koptic, Axes: []int{0, 6}},
		{Scale: scale.Poritonic, Axes: []int{0, 6}},
		{Scale: scale.Kadimic, Axes: []int{0, 6}},
		{Scale: scale.Phrynic, Axes: []int{0, 3, 6, 9}},
		{Scale: scale.Sylitonic, Axes: []int{0, 6}},
		{Scale: scale.Mocritonic, Axes: []int{0, 6}},
	}

	for _, tc := range testCases {
		t.Run(tc.Scale.String(), func(t *testing.T) {
			assert.True(t, tc.Scale.ReflectiveSymmetric())
			assert.Equal(t, tc.Axes, tc.Scale.ReflectiveSymmetryAxes())
		})
	}
}

func TestType_Balanced(t *testing.T) {
	assert.True(t, scale.Minoric.Balanced())
	assert.True(t, scale.Dadic.Balanced())
	assert.True(t, scale.Daptitonic.Balanced())
	assert.True(t, scale.Stadimic.Balanced())
	assert.True(t, scale.Phrynic.Balanced())
	assert.True(t, scale.Zynitonic.Balanced())
	assert.True(t, scale.Dodimic.Balanced())
	assert.True(t, scale.Mocritonic.Balanced())
	assert.True(t, scale.Aerythimic.Balanced())
	assert.True(t, scale.Zyrimic.Balanced())
	assert.True(t, scale.Epadian.Balanced())
	assert.True(t, scale.Ionodian.Balanced())
	assert.True(t, scale.Katogyllic.Balanced())
	assert.True(t, scale.Chromatic.Balanced())
}

func TestType_FifthGeneratorRoot(t *testing.T) {
	type testCase struct {
		Title        string
		GivenScale   scale.Type
		ExpectedRoot pitch.WithDegree
	}

	testCases := []testCase{
		{
			Title:        "Lydian",
			GivenScale:   scale.Lydian,
			ExpectedRoot: pitch.WithDegree{Pitch: pitch.CNatural, Degree: degree.FromInt(1)},
		},
		{
			Title:        "Phrygian",
			GivenScale:   scale.Phrygian,
			ExpectedRoot: pitch.WithDegree{Pitch: pitch.CSharp, Degree: degree.FromInt(2)},
		},
		{
			Title:        "Dorian",
			GivenScale:   scale.Dorian,
			ExpectedRoot: pitch.WithDegree{Pitch: pitch.DSharp, Degree: degree.FromInt(3)},
		},
		{
			Title:        "Ionian",
			GivenScale:   scale.Ionian,
			ExpectedRoot: pitch.WithDegree{Pitch: pitch.FNatural, Degree: degree.FromInt(4)},
		},
		{
			Title:        "Locrian",
			GivenScale:   scale.Locrian,
			ExpectedRoot: pitch.WithDegree{Pitch: pitch.FSharp, Degree: degree.FromInt(5)},
		},
		{
			Title:        "Aeolian",
			GivenScale:   scale.Aeolian,
			ExpectedRoot: pitch.WithDegree{Pitch: pitch.GSharp, Degree: degree.FromInt(6)},
		},
		{
			Title:        "Mixolydian",
			GivenScale:   scale.Mixolydian,
			ExpectedRoot: pitch.WithDegree{Pitch: pitch.ASharp, Degree: degree.FromInt(7)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			assert.Equal(t, tc.ExpectedRoot, tc.GivenScale.FifthGeneratorRoot())
		})
	}
}

func TestType_FifthGeneratorRootWithTonic(t *testing.T) {
	type testCase struct {
		Title        string
		GivenScale   scale.Type
		GivenTonic   pitch.Type
		ExpectedRoot pitch.WithDegree
	}

	testCases := []testCase{
		{
			Title:        "Lydian",
			GivenScale:   scale.Lydian,
			GivenTonic:   pitch.FNatural,
			ExpectedRoot: pitch.WithDegree{Pitch: pitch.FNatural, Degree: degree.FromInt(1)},
		},
		{
			Title:        "Phrygian",
			GivenScale:   scale.Phrygian,
			GivenTonic:   pitch.ENatural,
			ExpectedRoot: pitch.WithDegree{Pitch: pitch.FNatural, Degree: degree.FromInt(2)},
		},
		{
			Title:        "Dorian",
			GivenScale:   scale.Dorian,
			GivenTonic:   pitch.DNatural,
			ExpectedRoot: pitch.WithDegree{Pitch: pitch.FNatural, Degree: degree.FromInt(3)},
		},
		{
			Title:        "Ionian",
			GivenScale:   scale.Ionian,
			GivenTonic:   pitch.CNatural,
			ExpectedRoot: pitch.WithDegree{Pitch: pitch.FNatural, Degree: degree.FromInt(4)},
		},
		{
			Title:        "Locrian",
			GivenScale:   scale.Locrian,
			GivenTonic:   pitch.BNatural,
			ExpectedRoot: pitch.WithDegree{Pitch: pitch.FNatural, Degree: degree.FromInt(5)},
		},
		{
			Title:        "Aeolian",
			GivenScale:   scale.Aeolian,
			GivenTonic:   pitch.ANatural,
			ExpectedRoot: pitch.WithDegree{Pitch: pitch.FNatural, Degree: degree.FromInt(6)},
		},
		{
			Title:        "Mixolydian",
			GivenScale:   scale.Mixolydian,
			GivenTonic:   pitch.GNatural,
			ExpectedRoot: pitch.WithDegree{Pitch: pitch.FNatural, Degree: degree.FromInt(7)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			assert.Equal(t, tc.ExpectedRoot, tc.GivenScale.FifthGeneratorRootWithTonic(tc.GivenTonic))
		})
	}
}
