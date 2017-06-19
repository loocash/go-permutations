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
	result := len(Antilex(n))
	if plen != result {
		t.Errorf("permutation of %d elements should have %d variations instead of %d", n, plen, result)
	}
}
