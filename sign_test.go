package permutation

import "testing"
import "fmt"

func TestSign(t *testing.T) {
	var tests = []struct {
		perm []int
		sign int
	}{
		{[]int{1, 2, 3, 4}, 1},
		{[]int{1, 3, 2, 4}, -1},
		{[]int{4, 3, 2, 1}, 1},
	}

	for _, test := range tests {
		if sign := Sign(test.perm); sign != test.sign {
			t.Errorf("Sign %v wanted: %d, got: %d\n", test.perm, test.sign, sign)
		}
	}
}

func ExampleSign() {
	fmt.Println(Sign([]int{1, 2, 3, 4}))
	fmt.Println(Sign([]int{1, 3, 2, 4}))
	// Output:
	// 1
	// -1
}
