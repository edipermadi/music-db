package chord

import "github.com/edipermadi/music-db/pkg/theory/pitch"

// Type is a type for chord
type Type int

const (
	Invalid                                           Type = iota
	Major                                             Type = iota
	Minor                                             Type = iota
	Power                                             Type = iota
	Diminished                                        Type = iota
	Augmented                                         Type = iota
	MajorSuspendedSecond                              Type = iota
	MajorSuspendedFourth                              Type = iota
	MajorFlatFifth                                    Type = iota
	MajorDoubleSharpFifth                             Type = iota
	MinorSharpFifth                                   Type = iota
	MinorDoubleFlatFifth                              Type = iota
	MajorSuspendedSecondFlatFifth                     Type = iota
	MajorSuspendedSecondDoubleFlatFifth               Type = iota
	MajorSuspendedSecondSharpFifth                    Type = iota
	MajorSuspendedFourthFlatFifth                     Type = iota
	MajorSuspendedFourthSharpFifth                    Type = iota
	MajorSuspendedFourthDoubleSharpFifth              Type = iota
	MajorSuspendFlatFifthAddSharpFifth                Type = iota
	MajorSuspendedSecondSuspendedFourth               Type = iota
	Phrygian                                          Type = iota
	Lydian                                            Type = iota
	Locrian                                           Type = iota
	Quartal                                           Type = iota
	QuartalAugmented                                  Type = iota
	DominantSeventh                                   Type = iota
	MinorSeventh                                      Type = iota
	MajorSeventh                                      Type = iota
	MinorMajorSeventh                                 Type = iota
	DiminishedSeventh                                 Type = iota
	MinorSeventhFlatFifth                             Type = iota
	AugmentedSeventh                                  Type = iota
	AugmentedMajorSeventh                             Type = iota
	DominantSeventhFlatFifth                          Type = iota
	DominantSeventhSuspendedSecond                    Type = iota
	DominantSeventhSuspendedFourth                    Type = iota
	DominantSeventhSuspendedSecondFlatFifth           Type = iota
	DominantSeventhSuspendedSecondSharpFifth          Type = iota
	DominantSeventhSuspendedFourthFlatFifth           Type = iota
	DominantSeventhSuspendedFourthSharpFifth          Type = iota
	DominantSeventhSuspendSecondSuspendFourth         Type = iota
	MajorSeventhSuspendedSecond                       Type = iota
	MajorSeventhSuspendedSecondFlatFifth              Type = iota
	MajorSeventhSuspendedFourth                       Type = iota
	MajorSeventhSuspendedSecondSuspendedFourth        Type = iota
	MinorSeventhSharpFifth                            Type = iota
	MinorSeventhDoubleFlatFifth                       Type = iota
	MinorMajorSeventhSharpFifth                       Type = iota
	MinorMajorSeventhDoubleFlatFifthDoubleFlatSeventh Type = iota
	MajorSeventhFlatFifth                             Type = iota
	MajorSeventhDoubleSharpFifth                      Type = iota
	DiminishedMajorSeventh                            Type = iota
	MajorSeventhSuspendedFourthSharpFifth             Type = iota
	MajorSeventhSuspendedFourthDoubleSharpFifth       Type = iota
	PhrygianAugmentedSeventh                          Type = iota
	LydianMajorSeventh                                Type = iota
	MajorAddFourth                                    Type = iota
	MinorAddFourth                                    Type = iota
	DominantSeventhAddFourth                          Type = iota
	MajorSeventhAddFourth                             Type = iota
	MajorAddSharpFourth                               Type = iota
	MinorAddSharpFourth                               Type = iota
	DominantSeventhAddSharpFourth                     Type = iota
	MajorSeventhAddSharpFourth                        Type = iota
	SuspendedSecondFlatFifthAddSharpFifth             Type = iota
	MajorAddSixth                                     Type = iota
	MinorAddSixth                                     Type = iota
	MajorAddSixthFlatFifth                            Type = iota
	MajorAddSixthSuspendSecond                        Type = iota
	MajorAddSixthSuspendFourth                        Type = iota
	MajorAddSixthSuspendSecondFlatFifth               Type = iota
	MajorAddSixthSuspendSecondDoubleFlatFifth         Type = iota
	MajorAddNinth                                     Type = iota
	MinorAddNinth                                     Type = iota
	MajorAddSixthAddNinth                             Type = iota
	MinorAddSixthAddNinth                             Type = iota
	MajorAddFlatNinth                                 Type = iota
	MinorAddFlatNinth                                 Type = iota
	MajorAddSharpNinth                                Type = iota
)

// String returns chord name
func (c Type) String() string {
	return [...]string{
		"Invalid",
		"Major",
		"Minor",
		"Power",
		"Diminished",
		"Augmented",
		"MajorSuspendedSecond",
		"MajorSuspendedFourth",
		"MajorFlatFifth",
		"MajorDoubleSharpFifth",
		"MinorSharpFifth",
		"MinorDoubleFlatFifth",
		"MajorSuspendedSecondFlatFifth",
		"MajorSuspendedSecondDoubleFlatFifth",
		"MajorSuspendedSecondSharpFifth",
		"MajorSuspendedFourthFlatFifth",
		"MajorSuspendedFourthSharpFifth",
		"MajorSuspendedFourthDoubleSharpFifth",
		"MajorSuspendFlatFifthAddSharpFifth",
		"MajorSuspendedSecondSuspendedFourth",
		"Phrygian",
		"Lydian",
		"Locrian",
		"Quartal",
		"QuartalAugmented",
		"DominantSeventh",
		"MinorSeventh",
		"MajorSeventh",
		"MinorMajorSeventh",
		"DiminishedSeventh",
		"MinorSeventhFlatFifth",
		"AugmentedSeventh",
		"AugmentedMajorSeventh",
		"DominantSeventhFlatFifth",
		"DominantSeventhSuspendedSecond",
		"DominantSeventhSuspendedFourth",
		"DominantSeventhSuspendedSecondFlatFifth",
		"DominantSeventhSuspendedSecondSharpFifth",
		"DominantSeventhSuspendedFourthFlatFifth",
		"DominantSeventhSuspendedFourthSharpFifth",
		"DominantSeventhSuspendSecondSuspendFourth",
		"MajorSeventhSuspendedSecond",
		"MajorSeventhSuspendedSecondFlatFifth",
		"MajorSeventhSuspendedFourth",
		"MajorSeventhSuspendedSecondSuspendedFourth",
		"MinorSeventhSharpFifth",
		"MinorSeventhDoubleFlatFifth",
		"MinorMajorSeventhSharpFifth",
		"MinorMajorSeventhDoubleFlatFifthDoubleFlatSeventh",
		"MajorSeventhFlatFifth",
		"MajorSeventhDoubleSharpFifth",
		"DiminishedMajorSeventh",
		"MajorSeventhSuspendedFourthSharpFifth",
		"MajorSeventhSuspendedFourthDoubleSharpFifth",
		"PhrygianAugmentedSeventh",
		"LydianMajorSeventh",
		"MajorAddFourth",
		"MinorAddFourth",
		"DominantSeventhAddFourth",
		"MajorSeventhAddFourth",
		"MajorAddSharpFourth",
		"MinorAddSharpFourth",
		"DominantSeventhAddSharpFourth",
		"MajorSeventhAddSharpFourth",
		"SuspendedSecondFlatFifthAddSharpFifth",
		"MajorAddSixth",
		"MinorAddSixth",
		"MajorAddSixthFlatFifth",
		"MajorAddSixthSuspendSecond",
		"MajorAddSixthSuspendFourth",
		"MajorAddSixthSuspendSecondFlatFifth",
		"MajorAddSixthSuspendSecondDoubleFlatFifth",
		"MajorAddNinth",
		"MinorAddNinth",
		"MajorAddSixthAddNinth",
		"MinorAddSixthAddNinth",
		"MajorAddFlatNinth",
		"MinorAddFlatNinth",
		"MajorAddSharpNinth",
	}[c]
}

// Number returns chord number
func (c Type) Number() int {
	return [...]int{
		0,
		2192, 2320, 2064, 2336, 2184, 2576, 2128, 2208, 2180, 2312,
		2368, 2592, 2624, 2568, 2144, 2120, 2116, 2088, 2640, 3088,
		2096, 3104, 2114, 2113, 2194, 2322, 2193, 2321, 2340, 2338,
		2186, 2185, 2210, 2578, 2130, 2594, 2570, 2146, 2122, 2642,
		2577, 2593, 2129, 2641, 2314, 2370, 2313, 2372, 2209, 2181,
		2337, 2121, 2117, 3089, 2097, 2256, 2384, 2258, 2257, 2224,
		2352, 2226, 2225, 2600, 2196, 2324, 2212, 2580, 2132, 2596,
		2628, 2704, 2832, 2708, 2836, 3216, 3344, 2448,
	}[c]
}

// Pitches returns chord pitches
func (c Type) Pitches(root pitch.Type) []pitch.Type {
	amount := int(root - pitch.CNatural)
	pitches := make([]pitch.Type, 0)
	for _, v := range pitch.AllPitches() {
		if c.Number()&v.Number() == v.Number() {
			pitches = append(pitches, v.Transpose(amount))
		}
	}

	return pitches
}

// PitchClass returns chord pitch class
func (c Type) PitchClass() []int {
	class := make([]int, 0)
	for _, v := range c.Pitches(pitch.CNatural) {
		class = append(class, int(v-1))
	}
	return class
}

// IntervalPattern returns chord interval pattern in semitones
func (c Type) IntervalPattern() []int {
	class := c.PitchClass()

	var previous int
	pattern := make([]int, 0)
	for i, v := range class {
		if i > 0 {
			pattern = append(pattern, v-previous)
		}
		previous = v
	}
	return pattern
}

// Cardinality returns count of pitches in the chord
func (c Type) Cardinality() int {
	return len(c.Pitches(pitch.CNatural))
}

// AllChords return all chord templates
func AllChords() []Type {
	return []Type{
		Major,
		Minor,
		Power,
		Diminished,
		Augmented,
		MajorSuspendedSecond,
		MajorSuspendedFourth,
		MajorFlatFifth,
		MajorDoubleSharpFifth,
		MinorSharpFifth,
		MinorDoubleFlatFifth,
		MajorSuspendedSecondFlatFifth,
		MajorSuspendedSecondDoubleFlatFifth,
		MajorSuspendedSecondSharpFifth,
		MajorSuspendedFourthFlatFifth,
		MajorSuspendedFourthSharpFifth,
		MajorSuspendedFourthDoubleSharpFifth,
		MajorSuspendFlatFifthAddSharpFifth,
		MajorSuspendedSecondSuspendedFourth,
		Phrygian,
		Lydian,
		Locrian,
		Quartal,
		QuartalAugmented,
		DominantSeventh,
		MinorSeventh,
		MajorSeventh,
		MinorMajorSeventh,
		DiminishedSeventh,
		MinorSeventhFlatFifth,
		AugmentedSeventh,
		AugmentedMajorSeventh,
		DominantSeventhFlatFifth,
		DominantSeventhSuspendedSecond,
		DominantSeventhSuspendedFourth,
		DominantSeventhSuspendedSecondFlatFifth,
		DominantSeventhSuspendedSecondSharpFifth,
		DominantSeventhSuspendedFourthFlatFifth,
		DominantSeventhSuspendedFourthSharpFifth,
		DominantSeventhSuspendSecondSuspendFourth,
		MajorSeventhSuspendedSecond,
		MajorSeventhSuspendedSecondFlatFifth,
		MajorSeventhSuspendedFourth,
		MajorSeventhSuspendedSecondSuspendedFourth,
		MinorSeventhSharpFifth,
		MinorSeventhDoubleFlatFifth,
		MinorMajorSeventhSharpFifth,
		MinorMajorSeventhDoubleFlatFifthDoubleFlatSeventh,
		MajorSeventhFlatFifth,
		MajorSeventhDoubleSharpFifth,
		DiminishedMajorSeventh,
		MajorSeventhSuspendedFourthSharpFifth,
		MajorSeventhSuspendedFourthDoubleSharpFifth,
		PhrygianAugmentedSeventh,
		LydianMajorSeventh,
		MajorAddFourth,
		MinorAddFourth,
		DominantSeventhAddFourth,
		MajorSeventhAddFourth,
		MajorAddSharpFourth,
		MinorAddSharpFourth,
		DominantSeventhAddSharpFourth,
		MajorSeventhAddSharpFourth,
		SuspendedSecondFlatFifthAddSharpFifth,
		MajorAddSixth,
		MinorAddSixth,
		MajorAddSixthFlatFifth,
		MajorAddSixthSuspendSecond,
		MajorAddSixthSuspendFourth,
		MajorAddSixthSuspendSecondFlatFifth,
		MajorAddSixthSuspendSecondDoubleFlatFifth,
		MajorAddNinth,
		MinorAddNinth,
		MajorAddSixthAddNinth,
		MinorAddSixthAddNinth,
		MajorAddFlatNinth,
		MinorAddFlatNinth,
		MajorAddSharpNinth,
	}
}
