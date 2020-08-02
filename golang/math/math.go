package math

// Min of a and b
func Min(a, b int) int {
	if a-b < 0 {
		return a
	}

	return b
}

// Max of a and b
func Max(a, b int) int {
	if a-b > 0 {
		return a
	}

	return b
}
