package interval

// Type defines a type of music interval
type Type int

const (
	Invalid       Type = iota
	PerfectUnison Type = iota
	MinorSecond   Type = iota
	MajorSecond   Type = iota
	MinorThird    Type = iota
	MajorThird    Type = iota
	PerfectFourth Type = iota
	Tritone       Type = iota
	PerfectFifth  Type = iota
	MinorSixth    Type = iota
	MajorSixth    Type = iota
	MinorSeventh  Type = iota
	MajorSeventh  Type = iota
	PerfectOctave Type = iota
)

func (i Type) String() string {
	if i < PerfectUnison || i > PerfectOctave {
		return "Invalid"
	}

	return [...]string{
		"Invalid",
		"PerfectUnison",
		"MinorSecond",
		"MajorSecond",
		"MinorThird",
		"MajorThird",
		"PerfectFourth",
		"Tritone",
		"PerfectFifth",
		"MinorSixth",
		"MajorSixth",
		"MinorSeventh",
		"MajorSeventh",
		"PerfectOctave",
	}[i]
}

// Semitones return count of semitones
func (i Type) Semitones() int {
	if i < PerfectUnison || i > PerfectOctave {
		return 0
	}

	return int(i) - 1
}
