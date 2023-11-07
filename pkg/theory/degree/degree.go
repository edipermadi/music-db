package degree

// Type is pitch degree type
type Type int

// Pitch degree name enumerations
const (
	Invalid     Type = iota
	Tonic       Type = iota
	Supertonic  Type = iota
	Mediant     Type = iota
	Subdominant Type = iota
	Dominant    Type = iota
	Subtonic    Type = iota
	LeadingTone Type = iota
)

// String returns string representation of pitch degree
func (d Type) String() string {
	return [...]string{
		"Invalid",
		"Tonic",
		"Supertonic",
		"Mediant",
		"Subdominant",
		"Dominant",
		"Subtonic",
		"LeadingTone",
	}[d]
}

// FromInt returns pitch degree from tonic
func FromInt(v int) Type {
	switch v {
	case 1:
		return Tonic
	case 2:
		return Supertonic
	case 3:
		return Mediant
	case 4:
		return Subdominant
	case 5:
		return Dominant
	case 6:
		return Subtonic
	case 7:
		return LeadingTone
	default:
		return Invalid
	}
}
