package pitch

// Slice represents pitch slice
type Slice []Type

// Unique returns unique pitch slice
func (s Slice) Unique() Slice {
	pitchMap := make(map[Type]struct{})

	filtered := make([]Type, 0)
	for _, v := range s {
		if _, found := pitchMap[v]; !found {
			pitchMap[v] = struct{}{}
			filtered = append(filtered, v)
		}
	}

	return filtered
}

// Signature return pitch slice signature
func (s Slice) Signature() int {
	pitchMap := make(map[Type]struct{})

	var amount int
	for _, v := range s {
		if _, found := pitchMap[v]; !found {
			pitchMap[v] = struct{}{}
			amount += v.Number()
		}
	}

	return amount
}

// Transpose return transposed pitch slice
func (s Slice) Transpose(amount int) Slice {
	if amount == 0 {
		return s
	}

	transposed := make([]Type, 0)
	for _, v := range s {
		transposed = append(transposed, v.Transpose(amount))
	}

	return transposed
}

// Equal return true when two pitches are equal
func (s Slice) Equal(v Slice) bool {
	return s.Signature() == v.Signature()
}
