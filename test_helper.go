package permutation

// equals tests if two slices contain the same values in the same order
func equals(a, b []int) bool {
	lena, lenb := len(a), len(b)
	if lena != lenb {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// factorial computes n!
func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
