package scale

// Number return scale number
func (s Type) Number() int {
	return [...]int{
		0,

		// 3 notes, reference: reference https://allthescales.org/scales.php?n=3
		2184,

		// 4 notes, reference: reference https://allthescales.org/scales.php?n=4
		2193, 2328, 2244, 3144,
		2196, 2376, 2628, 2322,
		2185, 2200, 2440, 3140,
		2188, 2248, 3208, 2321,
		2194, 2344, 2372, 2596,
		2212, 2632, 2338, 2324,
		2186, 2216, 2696, 2594,
		2210, 2600,
		2340,
	}[s]
}
