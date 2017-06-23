package permutation

import "testing"



func TestPermutationsLength(t *testing.T) {
	n := 3
	plen := factorial(n)
	result := 0
	for range Antilex(n) {
		result++
	}
	if plen != result {
		t.Errorf("permutation of %d elements should have %d variations instead of %d", n, plen, result)
	}
}

func TestDistribution(t *testing.T) {
	const n = 4
	times := factorial(n - 1)
	var distribution [n][n]int

	for p := range Antilex(n) {
		for i, v := range p {
			distribution[v-1][i]++
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if distribution[i][j] != times {
				t.Errorf("element %d should appear exactly %d times in every place and appeared %d times in the index of %d\n", i+1, times, distribution[i][j], j)
			}
		}
	}
}
