package scale

import "github.com/edipermadi/music-db/pkg/theory/pitch"

// Type is a type for scale
type Type int

// Scale name enumerations
const (
	InvalidScale Type = iota
	Minoric      Type = iota
	Thaptic      Type = iota
	Lothic       Type = iota
	Phratic      Type = iota
	Aerathic     Type = iota
	Epathic      Type = iota
	Mynic        Type = iota
	Rothic       Type = iota
	Eporic       Type = iota
	Zyphic       Type = iota
	Epogic       Type = iota
	Lanic        Type = iota
	Pyrric       Type = iota
	Aeoloric     Type = iota
	Gonic        Type = iota
	Dalic        Type = iota
	Dygic        Type = iota
	Daric        Type = iota
	Lonic        Type = iota
	Phradic      Type = iota
	Bolic        Type = iota
	Saric        Type = iota
	Zoptic       Type = iota
	Aeraphic     Type = iota
	Byptic       Type = iota
	Aeolic       Type = iota
	Koptic       Type = iota
	Mixolyric    Type = iota
	Lydic        Type = iota
	Stathic      Type = iota
	Dadic        Type = iota
	Phrynic      Type = iota
)

// AllScales return all scales
func AllScales() []Type {
	return []Type{
		Minoric,
		Thaptic,
		Lothic,
		Phratic,
		Aerathic,
		Epathic,
		Mynic,
		Rothic,
		Eporic,
		Zyphic,
		Epogic,
		Lanic,
		Pyrric,
		Aeoloric,
		Gonic,
		Dalic,
		Dygic,
		Daric,
		Lonic,
		Phradic,
		Bolic,
		Saric,
		Zoptic,
		Aeraphic,
		Byptic,
		Aeolic,
		Koptic,
		Mixolyric,
		Lydic,
		Stathic,
		Dadic,
		Phrynic,
	}
}

// Pitches returns set of pitches of the scale with given tonic
func (s Type) Pitches(tonic pitch.Type) []pitch.Type {
	pitches := make([]pitch.Type, 0)
	for _, v := range pitch.AllPitches() {
		if s.Number()&v.Number() == v.Number() {
			pitches = append(pitches, v.Transpose(int(tonic-pitch.CNatural)))
		}
	}

	return pitches
}
