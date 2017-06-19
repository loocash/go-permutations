// Package permutation generates permutations
package permutation

import (
	"fmt"
	"os"
)

const (
	// MaxLen maximum length of permutation
	MaxLen = 15
)

func reverse(p []int, n int) {
	for i := 0; i < n/2; i++ {
		p[i], p[n-i-1] = p[n-i-1], p[i]
	}
}

// Antilex prints out all permutations in anti-lexicographical order
func Antilex(n int) chan []int {
	if n < 1 || n > MaxLen {
		err := fmt.Errorf("Antilex: invalid argument: %d", n)
		fmt.Println(err)
		os.Exit(1)
	}
	ch := make(chan []int)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i + 1
	}
	go func() {
		perm(p, n, ch)
		close(ch)
	}()
	return ch
}

func perm(p []int, n int, ch chan []int) {
	if n == 1 {
		cpy := make([]int, len(p))
		copy(cpy, p)
		ch <- cpy
	} else {
		perm(p, n-1, ch)
		for i := n - 2; i >= 0; i-- {
			reverse(p, n-1)
			p[i], p[n-1] = p[n-1], p[i]
			perm(p, n-1, ch)
		}
	}
}
