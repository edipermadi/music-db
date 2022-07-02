package theory

// String returns scale name
func (s ScaleType) String() string {
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
	}[s]
}
