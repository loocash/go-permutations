package permutation

// reverse reverses the first n elements of a slice
func reverse(p []int, n int) {
	for i := 0; i < n/2; i++ {
		p[i], p[n-i-1] = p[n-i-1], p[i]
	}
}
