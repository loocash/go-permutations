package permutation

import "testing"

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

func TestReverse(t *testing.T) {
	var tests = []struct{ in, out []int }{
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
	}

	for _, test := range tests {
		cpy := make([]int, len(test.in))
		copy(cpy, test.in)
		reverse(cpy, len(test.in))
		if !equals(cpy, test.out) {
			t.Errorf("%v reversed should be %v instead of %v\n", test.in, test.out, cpy)
		}
	}
}

func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

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
