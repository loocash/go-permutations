package permutation

import "testing"

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
