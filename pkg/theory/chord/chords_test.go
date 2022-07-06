package chord_test

import (
	"testing"

	"github.com/edipermadi/music-db/pkg/theory/chord"
	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"github.com/stretchr/testify/assert"
)

func TestChordType_Values(t *testing.T) {
	type testCase struct {
		Chord   chord.Type
		Name    string
		Pitches []pitch.Type
	}

	testCases := []testCase{
		{Chord: chord.Invalid, Name: "Invalid", Pitches: []pitch.Type{}},

		{Chord: chord.Major, Name: "Major", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.GNatural}},
		{Chord: chord.Minor, Name: "Minor", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.GNatural}},
		{Chord: chord.Power, Name: "Power", Pitches: []pitch.Type{pitch.CNatural, pitch.GNatural}},
		{Chord: chord.Diminished, Name: "Diminished", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.FSharp}},
		{Chord: chord.Augmented, Name: "Augmented", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.GSharp}},
		{Chord: chord.MajorSuspendedSecond, Name: "MajorSuspendedSecond", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.GNatural}},
		{Chord: chord.MajorSuspendedFourth, Name: "MajorSuspendedFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.GNatural}},
		{Chord: chord.MajorFlatFifth, Name: "MajorFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.FSharp}},
		{Chord: chord.MajorDoubleSharpFifth, Name: "MajorDoubleSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.ANatural}},
		{Chord: chord.MinorSharpFifth, Name: "MinorSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.GSharp}},

		{Chord: chord.MinorDoubleFlatFifth, Name: "MinorDoubleFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.FNatural}},
		{Chord: chord.MajorSuspendedSecondFlatFifth, Name: "MajorSuspendedSecondFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.FSharp}},
		{Chord: chord.MajorSuspendedSecondDoubleFlatFifth, Name: "MajorSuspendedSecondDoubleFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.FNatural}},
		{Chord: chord.MajorSuspendedSecondSharpFifth, Name: "MajorSuspendedSecondSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.GSharp}},
		{Chord: chord.MajorSuspendedFourthFlatFifth, Name: "MajorSuspendedFourthFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.FSharp}},
		{Chord: chord.MajorSuspendedFourthSharpFifth, Name: "MajorSuspendedFourthSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.GSharp}},
		{Chord: chord.MajorSuspendedFourthDoubleSharpFifth, Name: "MajorSuspendedFourthDoubleSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.ANatural}},
		{Chord: chord.MajorSuspendFlatFifthAddSharpFifth, Name: "MajorSuspendFlatFifthAddSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.FSharp, pitch.GSharp}},
		{Chord: chord.MajorSuspendedSecondSuspendedFourth, Name: "MajorSuspendedSecondSuspendedFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.FNatural, pitch.GNatural}},
		{Chord: chord.Phrygian, Name: "Phrygian", Pitches: []pitch.Type{pitch.CNatural, pitch.CSharp, pitch.GNatural}},

		{Chord: chord.Lydian, Name: "Lydian", Pitches: []pitch.Type{pitch.CNatural, pitch.FSharp, pitch.GNatural}},
		{Chord: chord.Locrian, Name: "Locrian", Pitches: []pitch.Type{pitch.CNatural, pitch.CSharp, pitch.FSharp}},
		{Chord: chord.Quartal, Name: "Quartal", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.ASharp}},
		{Chord: chord.QuartalAugmented, Name: "QuartalAugmented", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.BNatural}},
		{Chord: chord.DominantSeventh, Name: "DominantSeventh", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.GNatural, pitch.ASharp}},
		{Chord: chord.MinorSeventh, Name: "MinorSeventh", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.GNatural, pitch.ASharp}},
		{Chord: chord.MajorSeventh, Name: "MajorSeventh", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.GNatural, pitch.BNatural}},
		{Chord: chord.MinorMajorSeventh, Name: "MinorMajorSeventh", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.GNatural, pitch.BNatural}},
		{Chord: chord.DiminishedSeventh, Name: "DiminishedSeventh", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.FSharp, pitch.ANatural}},
		{Chord: chord.MinorSeventhFlatFifth, Name: "MinorSeventhFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.FSharp, pitch.ASharp}},

		{Chord: chord.AugmentedSeventh, Name: "AugmentedSeventh", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.GSharp, pitch.ASharp}},
		{Chord: chord.AugmentedMajorSeventh, Name: "AugmentedMajorSeventh", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.GSharp, pitch.BNatural}},
		{Chord: chord.DominantSeventhFlatFifth, Name: "DominantSeventhFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.FSharp, pitch.ASharp}},
		{Chord: chord.DominantSeventhSuspendedSecond, Name: "DominantSeventhSuspendedSecond", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.GNatural, pitch.ASharp}},
		{Chord: chord.DominantSeventhSuspendedFourth, Name: "DominantSeventhSuspendedFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.GNatural, pitch.ASharp}},
		{Chord: chord.DominantSeventhSuspendedSecondFlatFifth, Name: "DominantSeventhSuspendedSecondFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.FSharp, pitch.ASharp}},
		{Chord: chord.DominantSeventhSuspendedSecondSharpFifth, Name: "DominantSeventhSuspendedSecondSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.GSharp, pitch.ASharp}},
		{Chord: chord.DominantSeventhSuspendedFourthFlatFifth, Name: "DominantSeventhSuspendedFourthFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.FSharp, pitch.ASharp}},
		{Chord: chord.DominantSeventhSuspendedFourthSharpFifth, Name: "DominantSeventhSuspendedFourthSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.GSharp, pitch.ASharp}},
		{Chord: chord.DominantSeventhSuspendSecondSuspendFourth, Name: "DominantSeventhSuspendSecondSuspendFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.FNatural, pitch.GNatural, pitch.ASharp}},

		{Chord: chord.MajorSeventhSuspendedSecond, Name: "MajorSeventhSuspendedSecond", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.GNatural, pitch.BNatural}},
		{Chord: chord.MajorSeventhSuspendedSecondFlatFifth, Name: "MajorSeventhSuspendedSecondFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.FSharp, pitch.BNatural}},
		{Chord: chord.MajorSeventhSuspendedFourth, Name: "MajorSeventhSuspendedFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.GNatural, pitch.BNatural}},
		{Chord: chord.MajorSeventhSuspendedSecondSuspendedFourth, Name: "MajorSeventhSuspendedSecondSuspendedFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.FNatural, pitch.GNatural, pitch.BNatural}},
		{Chord: chord.MinorSeventhSharpFifth, Name: "MinorSeventhSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.GSharp, pitch.ASharp}},
		{Chord: chord.MinorSeventhDoubleFlatFifth, Name: "MinorSeventhDoubleFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.FNatural, pitch.ASharp}},
		{Chord: chord.MinorMajorSeventhSharpFifth, Name: "MinorMajorSeventhSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.GSharp, pitch.BNatural}},
		{Chord: chord.MinorMajorSeventhDoubleFlatFifthDoubleFlatSeventh, Name: "MinorMajorSeventhDoubleFlatFifthDoubleFlatSeventh", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.FNatural, pitch.ANatural}},
		{Chord: chord.MajorSeventhFlatFifth, Name: "MajorSeventhFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.FSharp, pitch.BNatural}},
		{Chord: chord.MajorSeventhDoubleSharpFifth, Name: "MajorSeventhDoubleSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.ANatural, pitch.BNatural}},

		{Chord: chord.DiminishedMajorSeventh, Name: "DiminishedMajorSeventh", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.FSharp, pitch.BNatural}},
		{Chord: chord.MajorSeventhSuspendedFourthSharpFifth, Name: "MajorSeventhSuspendedFourthSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.GSharp, pitch.BNatural}},
		{Chord: chord.MajorSeventhSuspendedFourthDoubleSharpFifth, Name: "MajorSeventhSuspendedFourthDoubleSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.ANatural, pitch.BNatural}},
		{Chord: chord.PhrygianAugmentedSeventh, Name: "PhrygianAugmentedSeventh", Pitches: []pitch.Type{pitch.CNatural, pitch.CSharp, pitch.GNatural, pitch.BNatural}},
		{Chord: chord.LydianMajorSeventh, Name: "LydianMajorSeventh", Pitches: []pitch.Type{pitch.CNatural, pitch.FSharp, pitch.GNatural, pitch.BNatural}},
		{Chord: chord.MajorAddFourth, Name: "MajorAddFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.FNatural, pitch.GNatural}},
		{Chord: chord.MinorAddFourth, Name: "MinorAddFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.FNatural, pitch.GNatural}},
		{Chord: chord.DominantSeventhAddFourth, Name: "DominantSeventhAddFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.FNatural, pitch.GNatural, pitch.ASharp}},
		{Chord: chord.MajorSeventhAddFourth, Name: "MajorSeventhAddFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.FNatural, pitch.GNatural, pitch.BNatural}},
		{Chord: chord.MajorAddSharpFourth, Name: "MajorAddSharpFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.FSharp, pitch.GNatural}},

		{Chord: chord.MinorAddSharpFourth, Name: "MinorAddSharpFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.FSharp, pitch.GNatural}},
		{Chord: chord.DominantSeventhAddSharpFourth, Name: "DominantSeventhAddSharpFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.FSharp, pitch.GNatural, pitch.ASharp}},
		{Chord: chord.MajorSeventhAddSharpFourth, Name: "MajorSeventhAddSharpFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.FSharp, pitch.GNatural, pitch.BNatural}},
		{Chord: chord.SuspendedSecondFlatFifthAddSharpFifth, Name: "SuspendedSecondFlatFifthAddSharpFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.FSharp, pitch.GSharp}},
		{Chord: chord.MajorAddSixth, Name: "MajorAddSixth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.GNatural, pitch.ANatural}},
		{Chord: chord.MinorAddSixth, Name: "MinorAddSixth", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.GNatural, pitch.ANatural}},
		{Chord: chord.MajorAddSixthFlatFifth, Name: "MajorAddSixthFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.FSharp, pitch.ANatural}},
		{Chord: chord.MajorAddSixthSuspendSecond, Name: "MajorAddSixthSuspendSecond", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.GNatural, pitch.ANatural}},
		{Chord: chord.MajorAddSixthSuspendFourth, Name: "MajorAddSixthSuspendFourth", Pitches: []pitch.Type{pitch.CNatural, pitch.FNatural, pitch.GNatural, pitch.ANatural}},
		{Chord: chord.MajorAddSixthSuspendSecondFlatFifth, Name: "MajorAddSixthSuspendSecondFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.FSharp, pitch.ANatural}},

		{Chord: chord.MajorAddSixthSuspendSecondDoubleFlatFifth, Name: "MajorAddSixthSuspendSecondDoubleFlatFifth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.FNatural, pitch.ANatural}},
		{Chord: chord.MajorAddNinth, Name: "MajorAddNinth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.ENatural, pitch.GNatural}},
		{Chord: chord.MinorAddNinth, Name: "MinorAddNinth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.DSharp, pitch.GNatural}},
		{Chord: chord.MajorAddSixthAddNinth, Name: "MajorAddSixthAddNinth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.ENatural, pitch.GNatural, pitch.ANatural}},
		{Chord: chord.MinorAddSixthAddNinth, Name: "MinorAddSixthAddNinth", Pitches: []pitch.Type{pitch.CNatural, pitch.DNatural, pitch.DSharp, pitch.GNatural, pitch.ANatural}},
		{Chord: chord.MajorAddFlatNinth, Name: "MajorAddFlatNinth", Pitches: []pitch.Type{pitch.CNatural, pitch.CSharp, pitch.ENatural, pitch.GNatural}},
		{Chord: chord.MinorAddFlatNinth, Name: "MinorAddFlatNinth", Pitches: []pitch.Type{pitch.CNatural, pitch.CSharp, pitch.DSharp, pitch.GNatural}},
		{Chord: chord.MajorAddSharpNinth, Name: "MajorAddSharpNinth", Pitches: []pitch.Type{pitch.CNatural, pitch.DSharp, pitch.ENatural, pitch.GNatural}},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Name, tc.Chord.String())
			assert.Equal(t, pitch.Slice(tc.Pitches).Signature(), tc.Chord.Number())
			assert.Equal(t, tc.Pitches, tc.Chord.Pitches(pitch.CNatural))
		})
	}
}
