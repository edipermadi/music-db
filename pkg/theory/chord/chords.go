package chord

import "github.com/edipermadi/music-db/pkg/theory/pitch"

// Quality is a type for chord quality
type Quality int

// Chord quality enumerations
const (
	Invalid                                           Quality = iota
	Major                                             Quality = iota
	Minor                                             Quality = iota
	Power                                             Quality = iota
	Diminished                                        Quality = iota
	Augmented                                         Quality = iota
	MajorSuspendedSecond                              Quality = iota
	MajorSuspendedFourth                              Quality = iota
	MajorFlatFifth                                    Quality = iota
	MajorDoubleSharpFifth                             Quality = iota
	MinorSharpFifth                                   Quality = iota
	MinorDoubleFlatFifth                              Quality = iota
	MajorSuspendedSecondFlatFifth                     Quality = iota
	MajorSuspendedSecondDoubleFlatFifth               Quality = iota
	MajorSuspendedSecondSharpFifth                    Quality = iota
	MajorSuspendedFourthFlatFifth                     Quality = iota
	MajorSuspendedFourthSharpFifth                    Quality = iota
	MajorSuspendedFourthDoubleSharpFifth              Quality = iota
	MajorSuspendFlatFifthAddSharpFifth                Quality = iota
	MajorSuspendedSecondSuspendedFourth               Quality = iota
	Phrygian                                          Quality = iota
	Lydian                                            Quality = iota
	Locrian                                           Quality = iota
	Quartal                                           Quality = iota
	QuartalAugmented                                  Quality = iota
	DominantSeventh                                   Quality = iota
	MinorSeventh                                      Quality = iota
	MajorSeventh                                      Quality = iota
	MinorMajorSeventh                                 Quality = iota
	DiminishedSeventh                                 Quality = iota
	MinorSeventhFlatFifth                             Quality = iota
	AugmentedSeventh                                  Quality = iota
	AugmentedMajorSeventh                             Quality = iota
	DominantSeventhFlatFifth                          Quality = iota
	DominantSeventhSuspendedSecond                    Quality = iota
	DominantSeventhSuspendedFourth                    Quality = iota
	DominantSeventhSuspendedSecondFlatFifth           Quality = iota
	DominantSeventhSuspendedSecondSharpFifth          Quality = iota
	DominantSeventhSuspendedFourthFlatFifth           Quality = iota
	DominantSeventhSuspendedFourthSharpFifth          Quality = iota
	DominantSeventhSuspendSecondSuspendFourth         Quality = iota
	MajorSeventhSuspendedSecond                       Quality = iota
	MajorSeventhSuspendedSecondFlatFifth              Quality = iota
	MajorSeventhSuspendedFourth                       Quality = iota
	MajorSeventhSuspendedSecondSuspendedFourth        Quality = iota
	MinorSeventhSharpFifth                            Quality = iota
	MinorSeventhDoubleFlatFifth                       Quality = iota
	MinorMajorSeventhSharpFifth                       Quality = iota
	MinorMajorSeventhDoubleFlatFifthDoubleFlatSeventh Quality = iota
	MajorSeventhFlatFifth                             Quality = iota
	MajorSeventhDoubleSharpFifth                      Quality = iota
	DiminishedMajorSeventh                            Quality = iota
	MajorSeventhSuspendedFourthSharpFifth             Quality = iota
	MajorSeventhSuspendedFourthDoubleSharpFifth       Quality = iota
	PhrygianAugmentedSeventh                          Quality = iota
	LydianMajorSeventh                                Quality = iota
	MajorAddFourth                                    Quality = iota
	MinorAddFourth                                    Quality = iota
	DominantSeventhAddFourth                          Quality = iota
	MajorSeventhAddFourth                             Quality = iota
	MajorAddSharpFourth                               Quality = iota
	MinorAddSharpFourth                               Quality = iota
	DominantSeventhAddSharpFourth                     Quality = iota
	MajorSeventhAddSharpFourth                        Quality = iota
	SuspendedSecondFlatFifthAddSharpFifth             Quality = iota
	MajorAddSixth                                     Quality = iota
	MinorAddSixth                                     Quality = iota
	MajorAddSixthFlatFifth                            Quality = iota
	MajorAddSixthSuspendSecond                        Quality = iota
	MajorAddSixthSuspendFourth                        Quality = iota
	MajorAddSixthSuspendSecondFlatFifth               Quality = iota
	MajorAddSixthSuspendSecondDoubleFlatFifth         Quality = iota
	MajorAddNinth                                     Quality = iota
	MinorAddNinth                                     Quality = iota
	MajorAddSixthAddNinth                             Quality = iota
	MinorAddSixthAddNinth                             Quality = iota
	MajorAddFlatNinth                                 Quality = iota
	MinorAddFlatNinth                                 Quality = iota
	MajorAddSharpNinth                                Quality = iota
)

// String returns chord name
func (q Quality) String() string {
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
	}[q]
}

// Number returns chord number
func (q Quality) Number() int {
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
	}[q]
}

// Pitches returns chord pitches
func (q Quality) Pitches(root pitch.Type) []pitch.Type {
	amount := int(root - pitch.CNatural)
	pitches := make([]pitch.Type, 0)
	for _, v := range pitch.AllPitches() {
		if q.Number()&v.Number() == v.Number() {
			pitches = append(pitches, v.Transpose(amount))
		}
	}

	return pitches
}

// PitchClass returns chord pitch class
func (q Quality) PitchClass() []int {
	class := make([]int, 0)
	for _, v := range q.Pitches(pitch.CNatural) {
		class = append(class, int(v-1))
	}
	return class
}

// IntervalPattern returns chord interval pattern in semitones
func (q Quality) IntervalPattern() []int {
	var previous int
	pattern := make([]int, 0)
	for i, v := range q.PitchClass() {
		if i > 0 {
			pattern = append(pattern, v-previous)
		}
		previous = v
	}
	return pattern
}

// Cardinality returns count of pitches in the chord
func (q Quality) Cardinality() int {
	return len(q.Pitches(pitch.CNatural))
}

// AllQualities return all chord qualities
func AllQualities() []Quality {
	return []Quality{
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
