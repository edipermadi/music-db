package theory

import (
	"encoding/json"
	"errors"
)

// SimplifiedPitch is simplified pitch object
type SimplifiedPitch struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// DetailedPitch is detailed pitch object
type DetailedPitch struct {
	ID            int64   `json:"id" db:"id"`
	Name          string  `json:"name" db:"name"`
	ZeitlerNumber int     `json:"zeitler_number" db:"zeitler_number"`
	RingNumber    int     `json:"ring_number" db:"ring_number"`
	Frequency     float64 `json:"frequency" db:"frequency"`
}

// SimplifiedChordQuality is simplified chord quality object
type SimplifiedChordQuality struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// DetailedChordQuality is detailed chord quality object
type DetailedChordQuality struct {
	ID              int64    `json:"id" db:"id"`
	Name            string   `json:"name" db:"name"`
	ZeitlerNumber   int      `json:"zeitler_number" db:"zeitler_number"`
	RingNumber      int      `json:"ring_number" db:"ring_number"`
	Cardinality     int      `json:"cardinality" db:"cardinality"`
	PitchClass      SliceInt `json:"pitch_class" db:"pitch_class"`
	IntervalPattern SliceInt `json:"interval_pattern" db:"interval_pattern"`
}

// SimplifiedChord is simplified chord object
type SimplifiedChord struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// DetailedChord is detailed chord object
type DetailedChord struct {
	ID            int64                  `json:"id" db:"id"`
	Quality       SimplifiedChordQuality `json:"quality" db:"quality"`
	Root          SimplifiedPitch        `json:"root" db:"root"`
	Name          string                 `json:"name" db:"name"`
	ZeitlerNumber int                    `json:"zeitler_number" db:"zeitler_number"`
	RingNumber    int                    `json:"ring_number" db:"ring_number"`
}

// ChordFilter represents chord filter
type ChordFilter struct {
	ChordQualityID int64 `form:"chord_quality_id"`
	RootID         int64 `form:"root_id"`
	ZeitlerNumber  int   `form:"zeitler_number"`
	RingNumber     int   `form:"ring_number"`
	Cardinality    int   `form:"cardinality"`
}

// DetailedScale is detailed scale object
type DetailedScale struct {
	ID                       int64    `json:"id" db:"id"`
	Name                     string   `json:"name" db:"name"`
	Cardinality              int      `json:"cardinality" db:"cardinality"`
	ZeitlerNumber            int      `json:"zeitler_number" db:"zeitler_number"`
	RingNumber               int      `json:"ring_number" db:"ring_number"`
	Perfection               int      `json:"perfection" db:"perfection"`
	Imperfection             int      `json:"imperfection" db:"imperfection"`
	PitchClass               SliceInt `json:"pitch_class" db:"pitch_class"`
	IntervalPattern          SliceInt `json:"interval_pattern" db:"interval_pattern"`
	RotationalSymmetric      bool     `json:"rotational_symmetric" db:"rotational_symmetric"`
	RotationalSymmetryLevel  int      `json:"rotational_symmetry_level" db:"rotational_symmetry_level"`
	Palindromic              bool     `json:"palindromic" db:"palindromic"`
	ReflectionalSymmetric    bool     `json:"reflectional_symmetric" db:"reflectional_symmetric"`
	ReflectionalSymmetryAxes SliceInt `json:"reflectional_symmetry_axes" db:"reflectional_symmetry_axes"`
	Balanced                 bool     `json:"balanced" db:"balanced"`
	FifthGeneratorRootDegree int      `json:"fifth_generator_root_degree" db:"fifth_generator_root_degree"`
}

// SimplifiedScale is simplified scale object
type SimplifiedScale struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// ScaleFilter is scale filter
type ScaleFilter struct {
	TonicID                 int64 `form:"tonic_id"`
	ZeitlerNumber           int   `form:"zeitler_number"`
	RingNumber              int   `form:"ring_number"`
	Perfection              *int  `form:"perfection"`
	Imperfection            *int  `form:"imperfection"`
	Balanced                *bool `form:"balanced"`
	RotationalSymmetric     *bool `form:"rotational_symmetric"`
	RotationalSymmetryLevel int   `form:"rotational_symmetry_level"`
	ReflectionalSymmetric   *bool `form:"reflectional_symmetric"`
	Palindromic             *bool `form:"palindromic"`
	Cardinality             int   `form:"cardinality"`
}

// DetailedKey is detailed key object
type DetailedKey struct {
	ID            int             `json:"id" db:"id"`
	Scale         SimplifiedScale `json:"scale" db:"scale"`
	Tonic         SimplifiedPitch `json:"tonic" db:"tonic"`
	Name          string          `json:"name" db:"name"`
	ZeitlerNumber int             `json:"zeitler_number" db:"zeitler_number"`
	RingNumber    int             `json:"ring_number" db:"ring_number"`
	Balanced      bool            `json:"balanced" db:"balanced"`
	CenterX       float64         `json:"center_x" db:"center_x"`
	CenterY       float64         `json:"center_y" db:"center_y"`
}

// SimplifiedKey is simplified key object
type SimplifiedKey struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// KeyFilter is key filter
type KeyFilter struct {
	ScaleID                 int64 `form:"scale_id"`
	TonicID                 int64 `form:"tonic_id"`
	ZeitlerNumber           int   `form:"zeitler_number"`
	RingNumber              int   `form:"ring_number"`
	Perfection              *int  `form:"perfection"`
	Imperfection            *int  `form:"imperfection"`
	Balanced                *bool `form:"balanced"`
	RotationalSymmetric     *bool `form:"rotational_symmetric"`
	RotationalSymmetryLevel int   `form:"rotational_symmetry_level"`
	ReflectionalSymmetric   *bool `form:"reflectional_symmetric"`
	Palindromic             *bool `form:"palindromic"`
	Cardinality             int   `form:"cardinality"`
}

// SliceInt implements array of int jsonb
type SliceInt []int

// Scan scan JSONB array of int
func (s *SliceInt) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, s)
}
