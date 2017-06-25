package permutation

// Sign computes sign of permutation
func Sign(p []int) int {
	sign := 1
	visited := make([]bool, len(p))
	for i, v := range visited {
		if !v {
			for j := p[i] - 1; j != i; j = p[j] - 1 {
				visited[j] = true
				sign = -sign
			}
		}
	}
	return sign
}
