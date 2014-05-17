// Functions for generating different arrays from inputs
// All methods return channels (generator pattern)

package main

import (
	"math"
	"strconv"
)

func powerSetGenerator(u Array) chan S {
	c := make(chan S)
	go func() {
		a := u.ToS()
		n := len(a)
		for i := 0; i < 1<<uint(n); i++ {
			var v []T
			for j := 0; j < n; j++ {
				if (uint(i) & (1 << uint(j))) > 0 {
					v = append(v, a[j])
				}
			}
			if len(v) > 0 {
				c <- v
			}
		}
		close(c)
	}()
	return c
}

func permutationsGenerator(u Array) chan S {
	c := make(chan S)
	go func() {
		var r []S
		a := u.ToS()
		var permutation func(i S, pos int)
		permutation = func(i S, pos int) {
			n := len(i)
			if n-pos == 1 {
				var err error
				j := append(S{}, i...)
				if r, err = appendIfUniqueArray(r, j); err == nil {
					c <- j
				}
				return
			} else {
				for j := pos; j < n; j++ {
					i[pos], i[j] = i[j], i[pos]
					permutation(i, pos+1)
					i[pos], i[j] = i[j], i[pos]
				}
			}
		}
		permutation(a, 0)
		close(c)
	}()
	return c
}

func multiSetGenerator(u Array, c int) chan S {
	ch := make(chan S)
	go func() {
		a := u.ToS()
		var r []S
		n := len(a)
		l := int(math.Pow(float64(c), float64(n)))
		for i := 0; i < l+1; i++ {
			s := strconv.FormatInt(int64(i), n)
			s = Reverse(s)
			if len(s) < c {
				s = (s + "000000000000000")[0:c]
			}
			var t []T
			for j := 0; j < c; j++ {
				t = append(t, a[s[j]-48])
			}
			var err error
			if r, err = appendIfUniqueArray(r, t[:c]); err == nil {
				ch <- t[:c]
			}
		}
		close(ch)
	}()
	return ch
}
