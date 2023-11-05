package pitch

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

func (p Type) String() string {
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
	if p < CNatural {
		return 0
	}
	return 1 << (12 - p)
}

// RingNumber return pitch numbering according to Ian Ring's system
func (p Type) RingNumber() int {
	if p < CNatural {
		return 0
	}

	return 1 << (p - 1)
}

// Transpose return transposed pitch
func (p Type) Transpose(amount int) Type {
	if p < CNatural {
		return Invalid
	}

	if amount == 0 {
		return p
	}

	val := (((int(p)-1)+amount)%12 + 12) % 12
	return fromInt(val + 1)
}

func fromInt(v int) Type {
	switch v {
	case 1:
		return CNatural
	case 2:
		return CSharp
	case 3:
		return DNatural
	case 4:
		return DSharp
	case 5:
		return ENatural
	case 6:
		return FNatural
	case 7:
		return FSharp
	case 8:
		return GNatural
	case 9:
		return GSharp
	case 10:
		return ANatural
	case 11:
		return ASharp
	case 12:
		return BNatural
	default:
		return Invalid
	}
}

// NextFifth returns next fifth
func (p Type) NextFifth() Type {
	switch p {
	case CNatural:
		return GNatural
	case CSharp:
		return GSharp
	case DNatural:
		return ANatural
	case DSharp:
		return ASharp
	case ENatural:
		return BNatural
	case FNatural:
		return CNatural
	case FSharp:
		return CSharp
	case GNatural:
		return DNatural
	case GSharp:
		return DSharp
	case ANatural:
		return ENatural
	case ASharp:
		return FNatural
	case BNatural:
		return FSharp
	default:
		return Invalid
	}
}

// PreviousFifth returns previous fifth
func (p Type) PreviousFifth() Type {
	switch p {
	case CNatural:
		return FNatural
	case CSharp:
		return FSharp
	case DNatural:
		return GNatural
	case DSharp:
		return GSharp
	case ENatural:
		return ANatural
	case FNatural:
		return ASharp
	case FSharp:
		return BNatural
	case GNatural:
		return CNatural
	case GSharp:
		return CSharp
	case ANatural:
		return DNatural
	case ASharp:
		return DSharp
	case BNatural:
		return ENatural
	default:
		return Invalid
	}
}

// Tritone returns a tritone away from given pitch
func (p Type) Tritone() Type {
	switch p {
	case CNatural:
		return FSharp
	case CSharp:
		return GNatural
	case DNatural:
		return GSharp
	case DSharp:
		return ANatural
	case ENatural:
		return ASharp
	case FNatural:
		return BNatural
	case FSharp:
		return CNatural
	case GNatural:
		return CSharp
	case GSharp:
		return DNatural
	case ANatural:
		return DSharp
	case ASharp:
		return ENatural
	case BNatural:
		return FNatural
	default:
		return Invalid
	}
}
