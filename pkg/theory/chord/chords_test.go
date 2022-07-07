package chord_test

import (
	"testing"

	"github.com/edipermadi/music-db/pkg/theory/chord"
	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"github.com/stretchr/testify/assert"
)

func TestChordType_Values(t *testing.T) {
	type testCase struct {
		Quality chord.Quality
		Name    string
		Pitches []pitch.Type
	}

	testCases := []testCase{
		{Quality: chord.Invalid, Name: "Invalid", Pitches: []pitch.Type{}},

		{Quality: chord.Major, Name: "Major", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.GNatural}},
		{Quality: chord.Minor, Name: "Minor", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.GNatural}},
		{Quality: chord.Power, Name: "Power", Pitches: []pitch.Type{pitch.CNatural, pitch.GNatural}},
		{Quality: chord.Diminished, Name: "Diminished", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.FSharp}},
		{Quality: chord.Augmented, Name: "Augmented", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.GSharp}},
		{Quality: chord.MajorSuspendedSecond, Name: "MajorSuspendedSecond", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.GNatural}},
		{Quality: chord.MajorSuspendedFourth, Name: "MajorSuspendedFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.GNatural}},
		{Quality: chord.MajorFlatFifth, Name: "MajorFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.FSharp}},
		{Quality: chord.MajorDoubleSharpFifth, Name: "MajorDoubleSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.ANatural}},
		{Quality: chord.MinorSharpFifth, Name: "MinorSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.GSharp}},

		{Quality: chord.MinorDoubleFlatFifth, Name: "MinorDoubleFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.FNatural}},
		{Quality: chord.MajorSuspendedSecondFlatFifth, Name: "MajorSuspendedSecondFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.FSharp}},
		{Quality: chord.MajorSuspendedSecondDoubleFlatFifth, Name: "MajorSuspendedSecondDoubleFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.FNatural}},
		{Quality: chord.MajorSuspendedSecondSharpFifth, Name: "MajorSuspendedSecondSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.GSharp}},
		{Quality: chord.MajorSuspendedFourthFlatFifth, Name: "MajorSuspendedFourthFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.FSharp}},
		{Quality: chord.MajorSuspendedFourthSharpFifth, Name: "MajorSuspendedFourthSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.GSharp}},
		{Quality: chord.MajorSuspendedFourthDoubleSharpFifth, Name: "MajorSuspendedFourthDoubleSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.ANatural}},
		{Quality: chord.MajorSuspendFlatFifthAddSharpFifth, Name: "MajorSuspendFlatFifthAddSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.FSharp, pitch.GSharp}},
		{Quality: chord.MajorSuspendedSecondSuspendedFourth, Name: "MajorSuspendedSecondSuspendedFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.FNatural, pitch.GNatural}},
		{Quality: chord.Phrygian, Name: "Phrygian", Pitches: []pitch.Type{pitch.CNatural, pitch.CSharp, pitch.GNatural}},

		{Quality: chord.Lydian, Name: "Lydian", Pitches: []pitch.Type{pitch.CNatural, pitch.FSharp, pitch.GNatural}},
		{Quality: chord.Locrian, Name: "Locrian", Pitches: []pitch.Type{pitch.CNatural, pitch.CSharp, pitch.FSharp}},
		{Quality: chord.Quartal, Name: "Quartal", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.ASharp}},
		{Quality: chord.QuartalAugmented, Name: "QuartalAugmented", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.BNatural}},
		{Quality: chord.DominantSeventh, Name: "DominantSeventh", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.GNatural, pitch.ASharp}},
		{Quality: chord.MinorSeventh, Name: "MinorSeventh", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.GNatural, pitch.ASharp}},
		{Quality: chord.MajorSeventh, Name: "MajorSeventh", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.GNatural, pitch.BNatural}},
		{Quality: chord.MinorMajorSeventh, Name: "MinorMajorSeventh", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.GNatural, pitch.BNatural}},
		{Quality: chord.DiminishedSeventh, Name: "DiminishedSeventh", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.FSharp, pitch.ANatural}},
		{Quality: chord.MinorSeventhFlatFifth, Name: "MinorSeventhFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.FSharp, pitch.ASharp}},

		{Quality: chord.AugmentedSeventh, Name: "AugmentedSeventh", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.GSharp, pitch.ASharp}},
		{Quality: chord.AugmentedMajorSeventh, Name: "AugmentedMajorSeventh", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.GSharp, pitch.BNatural}},
		{Quality: chord.DominantSeventhFlatFifth, Name: "DominantSeventhFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.FSharp, pitch.ASharp}},
		{Quality: chord.DominantSeventhSuspendedSecond, Name: "DominantSeventhSuspendedSecond", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.GNatural, pitch.ASharp}},
		{Quality: chord.DominantSeventhSuspendedFourth, Name: "DominantSeventhSuspendedFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.GNatural, pitch.ASharp}},
		{Quality: chord.DominantSeventhSuspendedSecondFlatFifth, Name: "DominantSeventhSuspendedSecondFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.FSharp, pitch.ASharp}},
		{Quality: chord.DominantSeventhSuspendedSecondSharpFifth, Name: "DominantSeventhSuspendedSecondSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.GSharp, pitch.ASharp}},
		{Quality: chord.DominantSeventhSuspendedFourthFlatFifth, Name: "DominantSeventhSuspendedFourthFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.FSharp, pitch.ASharp}},
		{Quality: chord.DominantSeventhSuspendedFourthSharpFifth, Name: "DominantSeventhSuspendedFourthSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.GSharp, pitch.ASharp}},
		{Quality: chord.DominantSeventhSuspendSecondSuspendFourth, Name: "DominantSeventhSuspendSecondSuspendFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.FNatural, pitch.GNatural, pitch.ASharp}},

		{Quality: chord.MajorSeventhSuspendedSecond, Name: "MajorSeventhSuspendedSecond", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.GNatural, pitch.BNatural}},
		{Quality: chord.MajorSeventhSuspendedSecondFlatFifth, Name: "MajorSeventhSuspendedSecondFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.FSharp, pitch.BNatural}},
		{Quality: chord.MajorSeventhSuspendedFourth, Name: "MajorSeventhSuspendedFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.GNatural, pitch.BNatural}},
		{Quality: chord.MajorSeventhSuspendedSecondSuspendedFourth, Name: "MajorSeventhSuspendedSecondSuspendedFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.FNatural, pitch.GNatural, pitch.BNatural}},
		{Quality: chord.MinorSeventhSharpFifth, Name: "MinorSeventhSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.GSharp, pitch.ASharp}},
		{Quality: chord.MinorSeventhDoubleFlatFifth, Name: "MinorSeventhDoubleFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.FNatural, pitch.ASharp}},
		{Quality: chord.MinorMajorSeventhSharpFifth, Name: "MinorMajorSeventhSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.GSharp, pitch.BNatural}},
		{Quality: chord.MinorMajorSeventhDoubleFlatFifthDoubleFlatSeventh, Name: "MinorMajorSeventhDoubleFlatFifthDoubleFlatSeventh", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.FNatural, pitch.ANatural}},
		{Quality: chord.MajorSeventhFlatFifth, Name: "MajorSeventhFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.FSharp, pitch.BNatural}},
		{Quality: chord.MajorSeventhDoubleSharpFifth, Name: "MajorSeventhDoubleSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.ANatural, pitch.BNatural}},

		{Quality: chord.DiminishedMajorSeventh, Name: "DiminishedMajorSeventh", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.FSharp, pitch.BNatural}},
		{Quality: chord.MajorSeventhSuspendedFourthSharpFifth, Name: "MajorSeventhSuspendedFourthSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.GSharp, pitch.BNatural}},
		{Quality: chord.MajorSeventhSuspendedFourthDoubleSharpFifth, Name: "MajorSeventhSuspendedFourthDoubleSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.ANatural, pitch.BNatural}},
		{Quality: chord.PhrygianAugmentedSeventh, Name: "PhrygianAugmentedSeventh", Pitches: []pitch.Type{pitch.CNatural, pitch.CSharp, pitch.GNatural, pitch.BNatural}},
		{Quality: chord.LydianMajorSeventh, Name: "LydianMajorSeventh", Pitches: []pitch.Type{pitch.CNatural, pitch.FSharp, pitch.GNatural, pitch.BNatural}},
		{Quality: chord.MajorAddFourth, Name: "MajorAddFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.FNatural, pitch.GNatural}},
		{Quality: chord.MinorAddFourth, Name: "MinorAddFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.FNatural, pitch.GNatural}},
		{Quality: chord.DominantSeventhAddFourth, Name: "DominantSeventhAddFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.FNatural, pitch.GNatural, pitch.ASharp}},
		{Quality: chord.MajorSeventhAddFourth, Name: "MajorSeventhAddFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.FNatural, pitch.GNatural, pitch.BNatural}},
		{Quality: chord.MajorAddSharpFourth, Name: "MajorAddSharpFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.FSharp, pitch.GNatural}},

		{Quality: chord.MinorAddSharpFourth, Name: "MinorAddSharpFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.FSharp, pitch.GNatural}},
		{Quality: chord.DominantSeventhAddSharpFourth, Name: "DominantSeventhAddSharpFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.FSharp, pitch.GNatural, pitch.ASharp}},
		{Quality: chord.MajorSeventhAddSharpFourth, Name: "MajorSeventhAddSharpFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.FSharp, pitch.GNatural, pitch.BNatural}},
		{Quality: chord.SuspendedSecondFlatFifthAddSharpFifth, Name: "SuspendedSecondFlatFifthAddSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.FSharp, pitch.GSharp}},
		{Quality: chord.MajorAddSixth, Name: "MajorAddSixth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.GNatural, pitch.ANatural}},
		{Quality: chord.MinorAddSixth, Name: "MinorAddSixth", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.GNatural, pitch.ANatural}},
		{Quality: chord.MajorAddSixthFlatFifth, Name: "MajorAddSixthFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.FSharp, pitch.ANatural}},
		{Quality: chord.MajorAddSixthSuspendSecond, Name: "MajorAddSixthSuspendSecond", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.GNatural, pitch.ANatural}},
		{Quality: chord.MajorAddSixthSuspendFourth, Name: "MajorAddSixthSuspendFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.GNatural, pitch.ANatural}},
		{Quality: chord.MajorAddSixthSuspendSecondFlatFifth, Name: "MajorAddSixthSuspendSecondFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.FSharp, pitch.ANatural}},

		{Quality: chord.MajorAddSixthSuspendSecondDoubleFlatFifth, Name: "MajorAddSixthSuspendSecondDoubleFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.FNatural, pitch.ANatural}},
		{Quality: chord.MajorAddNinth, Name: "MajorAddNinth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.ENatural, pitch.GNatural}},
		{Quality: chord.MinorAddNinth, Name: "MinorAddNinth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.DSharp, pitch.GNatural}},
		{Quality: chord.MajorAddSixthAddNinth, Name: "MajorAddSixthAddNinth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.ENatural, pitch.GNatural, pitch.ANatural}},
		{Quality: chord.MinorAddSixthAddNinth, Name: "MinorAddSixthAddNinth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.DSharp, pitch.GNatural, pitch.ANatural}},
		{Quality: chord.MajorAddFlatNinth, Name: "MajorAddFlatNinth", Pitches: []pitch.Type{pitch.CNatural, pitch.CSharp, pitch.ENatural, pitch.GNatural}},
		{Quality: chord.MinorAddFlatNinth, Name: "MinorAddFlatNinth", Pitches: []pitch.Type{pitch.CNatural, pitch.CSharp, pitch.DSharp, pitch.GNatural}},
		{Quality: chord.MajorAddSharpNinth, Name: "MajorAddSharpNinth", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.ENatural, pitch.GNatural}},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Name, tc.Quality.String())
			assert.Equal(t, pitch.Slice(tc.Pitches).Signature(), tc.Quality.Number())
			assert.Equal(t, tc.Pitches, tc.Quality.Pitches(pitch.CNatural))
		})
	}
}
