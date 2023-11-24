package pitch

import (
	"github.com/edipermadi/music-db/pkg/theory/degree"
	"github.com/edipermadi/music-db/pkg/theory/interval"
)

// Type is a type for pitch
type Type int

// Pitch name enumerations
const (
	Invalid  Type = iota
	CNatural Type = iota
	CSharp   Type = iota
	DNatural Type = iota
	DSharp   Type = iota
	ENatural Type = iota
	FNatural Type = iota
	FSharp   Type = iota
	GNatural Type = iota
	GSharp   Type = iota
	ANatural Type = iota
	ASharp   Type = iota
	BNatural Type = iota
)

// AllPitches returns all pitches
func AllPitches() []Type {
	return []Type{
		CNatural,
		CSharp,
		DNatural,
		DSharp,
		ENatural,
		FNatural,
		FSharp,
		GNatural,
		GSharp,
		ANatural,
		ASharp,
		BNatural,
	}
}

func CircleOfFifths() []Type {
	root := Invalid
	result := make([]Type, 12)
	for i := 0; i < 12; i++ {
		if i == 0 {
			root = CNatural
		} else {
			root = root.NextFifth()
		}

		result[i] = root
	}

	return result
}

func (p Type) String() string {
	if p < CNatural || p > BNatural {
		return "Invalid"
	}

	return [...]string{
		"Invalid",
		"CNatural",
		"CSharp",
		"DNatural",
		"DSharp",
		"ENatural",
		"FNatural",
		"FSharp",
		"GNatural",
		"GSharp",
		"ANatural",
		"ASharp",
		"BNatural",
	}[p]
}

// Frequency returns pitch frequency
func (p Type) Frequency() float64 {
	if p < CNatural || p > BNatural {
		return 0
	}

	return [...]float64{
		0,
		261.63,
		277.18,
		293.66,
		311.13,
		329.63,
		349.23,
		369.99,
		392.00,
		415.30,
		440.00,
		466.16,
		493.88,
	}[p]
}

// Number return patch number according to William Zeitler's numbering system
func (p Type) Number() int {
	return p.ZeitlerNumber()
}

// ZeitlerNumber return pitch numbering according to William Zeitler's system
func (p Type) ZeitlerNumber() int {
	if p < CNatural || p > BNatural {
		return 0
	}
	return 1 << (12 - p)
}

// RingNumber return pitch numbering according to Ian Ring's system
func (p Type) RingNumber() int {
	if p < CNatural || p > BNatural {
		return 0
	}

	return 1 << (p - 1)
}

// Transpose return transposed pitch
func (p Type) Transpose(amount int) Type {
	if p < CNatural || p > BNatural {
		return Invalid
	}

	if amount == 0 {
		return p
	}

	val := (((int(p)-1)+amount)%12 + 12) % 12
	return FromInt(val + 1)
}

func FromInt(v int) Type {
	if v < 1 || v > 12 {
		return Invalid
	}

	return Type(v)
}

// NextFifth returns next fifth
func (p Type) NextFifth() Type {
	return p.Transpose(interval.PerfectFifth.Semitones())
}

// PreviousFifth returns previous fifth
func (p Type) PreviousFifth() Type {
	return p.Transpose(-interval.PerfectFifth.Semitones())
}

// Tritone returns a tritone away from given pitch
func (p Type) Tritone() Type {
	return p.Transpose(interval.Tritone.Semitones())
}

// WithDegree represents a tuple of pitch with it's correspond degree
type WithDegree struct {
	Pitch  Type
	Degree degree.Type
}

// Func defines a pitch function that maps a pitch from one to another
type Func func(p Type) Type
