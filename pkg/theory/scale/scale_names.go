package scale

// String returns scale name
func (s Type) String() string {
	return [...]string{
		// invalid
		"InvalidScale",

		// 3 notes, reference https://allthescales.org/scales.php?n=3
		"Minoric",

		// 4 notes, reference https://allthescales.org/scales.php?n=4
		"Thaptic", "Lothic", "Phratic", "Aerathic",
		"Epathic", "Mynic", "Rothic", "Eporic",
		"Zyphic", "Epogic", "Lanic", "Pyrric",
		"Aeoloric", "Gonic", "Dalic", "Dygic",
		"Daric", "Lonic", "Phradic", "Bolic",
		"Saric", "Zoptic", "Aeraphic", "Byptic",
		"Aeolic", "Koptic", "Mixolyric", "Lydic",
		"Stathic", "Dadic",
		"Phrynic",

		// 5 notes, reference https://allthescales.org/scales.php?n=5
		"Epathitonic", "Mynitonic", "Rocritonic", "Pentatonic", "Thaptitonic",
		"Magitonic", "Daditonic", "Aeolyphritonic", "Gycritonic", "Pyritonic",
		"Gathitonic", "Ionitonic", "Phrynitonic", "Stathitonic", "Thalitonic",
		"Zolitonic", "Epogitonic", "Lanitonic", "Paptitonic", "Ionacritonic",
		"Phraditonic", "Aeoloritonic", "Gonitonic", "Dalitonic", "Dygitonic",
		"Aeracritonic", "Byptitonic", "Daritonic", "Lonitonic", "Ionycritonic",
		"Lothitonic", "Phratonic", "Aerathitonic", "Saritonic", "Zoptitonic",
		"Dolitonic", "Poritonic", "Aerylitonic", "Zagitonic", "Lagitonic",
		"Molitonic", "Staptitonic", "Mothitonic", "Aeritonic", "Ragitonic",
		"Ionaditonic", "Bocritonic", "Gythitonic", "Pagitonic", "Aeolythitonic",
		"Zacritonic", "Laritonic", "Thacritonic", "Styditonic", "Loritonic",
		"Aeolyritonic", "Goritonic", "Aeoloditonic", "Doptitonic", "Aeraphitonic",
		"Zathitonic", "Raditonic", "Stonitonic", "Syptitonic", "Ionythitonic",
		"Aeolanitonic", "Danitonic", "Ionaritonic", "Dynitonic", "Zyditonic",
		"Aeolacritonic", "Zythitonic", "Dyritonic", "Koptitonic", "Thocritonic",
		"Lycritonic", "Daptitonic", "Kygitonic", "Mocritonic", "Zynitonic",
		"Epygitonic", "Zaptitonic", "Kagitonic", "Zogitonic", "Epyritonic",
		"Zothitonic", "Phrolitonic", "Ionagitonic", "Aeolapritonic", "Kyritonic",
		"Ionyptitonic", "Gyritonic", "Zalitonic", "Stolitonic", "Bylitonic",
		"Thoditonic", "Dogitonic", "Phralitonic", "Garitonic", "Soptitonic",
		"Kataritonic", "Sylitonic", "Thonitonic", "Phropitonic", "Staditonic",
		"Lyditonic", "Mythitonic", "Sogitonic", "Gothitonic", "Rothitonic",
		"Zylitonic", "Zoditonic", "Zaritonic", "Phrythitonic", "Rolitonic",
		"Ranitonic", "Laditonic", "Poditonic", "Ionothitonic", "Kanitonic",
		"Ryphitonic", "Gylitonic", "Aeolycritonic", "Pynitonic", "Zanitonic",
		"Phronitonic", "Banitonic", "Aeronitonic", "Golitonic", "Dyptitonic",
		"Aerynitonic", "Palitonic", "Stothitonic", "Aerophitonic", "Katagitonic",
		"Ionoditonic", "Bogitonic", "Mogitonic", "Docritonic", "Epaditonic",
		"Mixitonic", "Phrothitonic", "Katycritonic", "Ionalitonic", "Loptitonic",
		"Thyritonic", "Thoptitonic", "Bycritonic", "Pathitonic", "Myditonic",
		"Bolitonic", "Bothitonic", "Kataditonic", "Koditonic", "Tholitonic",
	}[s]
}
