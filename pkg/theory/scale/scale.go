package scale

import "github.com/edipermadi/music-db/pkg/theory/pitch"

// Type is a type for scale
type Type int

// Scale name enumerations
const (
	Invalid        Type = iota
	Minoric        Type = iota
	Thaptic        Type = iota
	Lothic         Type = iota
	Phratic        Type = iota
	Aerathic       Type = iota
	Epathic        Type = iota
	Mynic          Type = iota
	Rothic         Type = iota
	Eporic         Type = iota
	Zyphic         Type = iota
	Epogic         Type = iota
	Lanic          Type = iota
	Pyrric         Type = iota
	Aeoloric       Type = iota
	Gonic          Type = iota
	Dalic          Type = iota
	Dygic          Type = iota
	Daric          Type = iota
	Lonic          Type = iota
	Phradic        Type = iota
	Bolic          Type = iota
	Saric          Type = iota
	Zoptic         Type = iota
	Aeraphic       Type = iota
	Byptic         Type = iota
	Aeolic         Type = iota
	Koptic         Type = iota
	Mixolyric      Type = iota
	Lydic          Type = iota
	Stathic        Type = iota
	Dadic          Type = iota
	Phrynic        Type = iota
	Epathitonic    Type = iota
	Mynitonic      Type = iota
	Rocritonic     Type = iota
	Pentatonic     Type = iota
	Thaptitonic    Type = iota
	Magitonic      Type = iota
	Daditonic      Type = iota
	Aeolyphritonic Type = iota
	Gycritonic     Type = iota
	Pyritonic      Type = iota
	Gathitonic     Type = iota
	Ionitonic      Type = iota
	Phrynitonic    Type = iota
	Stathitonic    Type = iota
	Thalitonic     Type = iota
	Zolitonic      Type = iota
	Epogitonic     Type = iota
	Lanitonic      Type = iota
	Paptitonic     Type = iota
	Ionacritonic   Type = iota
	Phraditonic    Type = iota
	Aeoloritonic   Type = iota
	Gonitonic      Type = iota
	Dalitonic      Type = iota
	Dygitonic      Type = iota
	Aeracritonic   Type = iota
	Byptitonic     Type = iota
	Daritonic      Type = iota
	Lonitonic      Type = iota
	Ionycritonic   Type = iota
	Lothitonic     Type = iota
	Phratonic      Type = iota
	Aerathitonic   Type = iota
	Saritonic      Type = iota
	Zoptitonic     Type = iota
	Dolitonic      Type = iota
	Poritonic      Type = iota
	Aerylitonic    Type = iota
	Zagitonic      Type = iota
	Lagitonic      Type = iota
	Molitonic      Type = iota
	Staptitonic    Type = iota
	Mothitonic     Type = iota
	Aeritonic      Type = iota
	Ragitonic      Type = iota
	Ionaditonic    Type = iota
	Bocritonic     Type = iota
	Gythitonic     Type = iota
	Pagitonic      Type = iota
	Aeolythitonic  Type = iota
	Zacritonic     Type = iota
	Laritonic      Type = iota
	Thacritonic    Type = iota
	Styditonic     Type = iota
	Loritonic      Type = iota
	Aeolyritonic   Type = iota
	Goritonic      Type = iota
	Aeoloditonic   Type = iota
	Doptitonic     Type = iota
	Aeraphitonic   Type = iota
	Zathitonic     Type = iota
	Raditonic      Type = iota
	Stonitonic     Type = iota
	Syptitonic     Type = iota
	Ionythitonic   Type = iota
	Aeolanitonic   Type = iota
	Danitonic      Type = iota
	Ionaritonic    Type = iota
	Dynitonic      Type = iota
	Zyditonic      Type = iota
	Aeolacritonic  Type = iota
	Zythitonic     Type = iota
	Dyritonic      Type = iota
	Koptitonic     Type = iota
	Thocritonic    Type = iota
	Lycritonic     Type = iota
	Daptitonic     Type = iota
	Kygitonic      Type = iota
	Mocritonic     Type = iota
	Zynitonic      Type = iota
	Epygitonic     Type = iota
	Zaptitonic     Type = iota
	Kagitonic      Type = iota
	Zogitonic      Type = iota
	Epyritonic     Type = iota
	Zothitonic     Type = iota
	Phrolitonic    Type = iota
	Ionagitonic    Type = iota
	Aeolapritonic  Type = iota
	Kyritonic      Type = iota
	Ionyptitonic   Type = iota
	Gyritonic      Type = iota
	Zalitonic      Type = iota
	Stolitonic     Type = iota
	Bylitonic      Type = iota
	Thoditonic     Type = iota
	Dogitonic      Type = iota
	Phralitonic    Type = iota
	Garitonic      Type = iota
	Soptitonic     Type = iota
	Kataritonic    Type = iota
	Sylitonic      Type = iota
	Thonitonic     Type = iota
	Phropitonic    Type = iota
	Staditonic     Type = iota
	Lyditonic      Type = iota
	Mythitonic     Type = iota
	Sogitonic      Type = iota
	Gothitonic     Type = iota
	Rothitonic     Type = iota
	Zylitonic      Type = iota
	Zoditonic      Type = iota
	Zaritonic      Type = iota
	Phrythitonic   Type = iota
	Rolitonic      Type = iota
	Ranitonic      Type = iota
	Laditonic      Type = iota
	Poditonic      Type = iota
	Ionothitonic   Type = iota
	Kanitonic      Type = iota
	Ryphitonic     Type = iota
	Gylitonic      Type = iota
	Aeolycritonic  Type = iota
	Pynitonic      Type = iota
	Zanitonic      Type = iota
	Phronitonic    Type = iota
	Banitonic      Type = iota
	Aeronitonic    Type = iota
	Golitonic      Type = iota
	Dyptitonic     Type = iota
	Aerynitonic    Type = iota
	Palitonic      Type = iota
	Stothitonic    Type = iota
	Aerophitonic   Type = iota
	Katagitonic    Type = iota
	Ionoditonic    Type = iota
	Bogitonic      Type = iota
	Mogitonic      Type = iota
	Docritonic     Type = iota
	Epaditonic     Type = iota
	Mixitonic      Type = iota
	Phrothitonic   Type = iota
	Katycritonic   Type = iota
	Ionalitonic    Type = iota
	Loptitonic     Type = iota
	Thyritonic     Type = iota
	Thoptitonic    Type = iota
	Bycritonic     Type = iota
	Pathitonic     Type = iota
	Myditonic      Type = iota
	Bolitonic      Type = iota
	Bothitonic     Type = iota
	Kataditonic    Type = iota
	Koditonic      Type = iota
	Tholitonic     Type = iota
)

// AllScales return all scales
func AllScales() []Type {
	return []Type{
		// 3 notes
		Minoric,

		// 4 notes
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

		// 5 notes
		Epathitonic,
		Mynitonic,
		Rocritonic,
		Pentatonic,
		Thaptitonic,
		Magitonic,
		Daditonic,
		Aeolyphritonic,
		Gycritonic,
		Pyritonic,
		Gathitonic,
		Ionitonic,
		Phrynitonic,
		Stathitonic,
		Thalitonic,
		Zolitonic,
		Epogitonic,
		Lanitonic,
		Paptitonic,
		Ionacritonic,
		Phraditonic,
		Aeoloritonic,
		Gonitonic,
		Dalitonic,
		Dygitonic,
		Aeracritonic,
		Byptitonic,
		Daritonic,
		Lonitonic,
		Ionycritonic,
		Lothitonic,
		Phratonic,
		Aerathitonic,
		Saritonic,
		Zoptitonic,
		Dolitonic,
		Poritonic,
		Aerylitonic,
		Zagitonic,
		Lagitonic,
		Molitonic,
		Staptitonic,
		Mothitonic,
		Aeritonic,
		Ragitonic,
		Ionaditonic,
		Bocritonic,
		Gythitonic,
		Pagitonic,
		Aeolythitonic,
		Zacritonic,
		Laritonic,
		Thacritonic,
		Styditonic,
		Loritonic,
		Aeolyritonic,
		Goritonic,
		Aeoloditonic,
		Doptitonic,
		Aeraphitonic,
		Zathitonic,
		Raditonic,
		Stonitonic,
		Syptitonic,
		Ionythitonic,
		Aeolanitonic,
		Danitonic,
		Ionaritonic,
		Dynitonic,
		Zyditonic,
		Aeolacritonic,
		Zythitonic,
		Dyritonic,
		Koptitonic,
		Thocritonic,
		Lycritonic,
		Daptitonic,
		Kygitonic,
		Mocritonic,
		Zynitonic,
		Epygitonic,
		Zaptitonic,
		Kagitonic,
		Zogitonic,
		Epyritonic,
		Zothitonic,
		Phrolitonic,
		Ionagitonic,
		Aeolapritonic,
		Kyritonic,
		Ionyptitonic,
		Gyritonic,
		Zalitonic,
		Stolitonic,
		Bylitonic,
		Thoditonic,
		Dogitonic,
		Phralitonic,
		Garitonic,
		Soptitonic,
		Kataritonic,
		Sylitonic,
		Thonitonic,
		Phropitonic,
		Staditonic,
		Lyditonic,
		Mythitonic,
		Sogitonic,
		Gothitonic,
		Rothitonic,

		Zylitonic,
		Zoditonic,
		Zaritonic,
		Phrythitonic,
		Rolitonic,

		Ranitonic,
		Laditonic,
		Poditonic,
		Ionothitonic,
		Kanitonic,

		Ryphitonic,
		Gylitonic,
		Aeolycritonic,
		Pynitonic,
		Zanitonic,

		Phronitonic,
		Banitonic,
		Aeronitonic,
		Golitonic,
		Dyptitonic,

		Aerynitonic,
		Palitonic,
		Stothitonic,
		Aerophitonic,
		Katagitonic,

		Ionoditonic,
		Bogitonic,
		Mogitonic,
		Docritonic,
		Epaditonic,

		Mixitonic,
		Phrothitonic,
		Katycritonic,
		Ionalitonic,
		Loptitonic,

		Thyritonic,
		Thoptitonic,
		Bycritonic,
		Pathitonic,
		Myditonic,

		Bolitonic,
		Bothitonic,
		Kataditonic,
		Koditonic,
		Tholitonic,
	}
}

// Pitches returns set of pitches of the scale with given tonic
func (s Type) Pitches(tonic pitch.Type) []pitch.Type {
	amount := int(tonic - pitch.CNatural)
	pitches := make([]pitch.Type, 0)
	for _, v := range pitch.AllPitches() {
		if s.Number()&v.Number() == v.Number() {
			pitches = append(pitches, v.Transpose(amount))
		}
	}

	return pitches
}

// Perfection stores perfection profile
type Perfection struct {
	Perfection   int
	Imperfection int
}

func (s Type) Perfection() Perfection {
	pitchMap := make(map[pitch.Type]struct{})
	pitches := s.Pitches(pitch.CNatural)

	for _, v := range pitches {
		pitchMap[v] = struct{}{}
	}

	var result Perfection
	for _, v := range pitches {
		_, found := pitchMap[v.NextFifth()]
		if found {
			result.Perfection++
		} else {
			result.Imperfection++
		}
	}

	return result
}
