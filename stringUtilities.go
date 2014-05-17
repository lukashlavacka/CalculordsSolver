// To string methods
package main

import (
	"fmt"
	"strconv"
)

func (s Solution) String() string {
	r := strconv.FormatUint(uint64(s.nums[0]), 10)
	for i := 0; i < len(s.ops); i++ {
		r += string(s.ops[i]) + strconv.FormatUint(uint64(s.nums[i+1]), 10)
	}
	return r
}

func (o Operations) String() string {
	var a []byte
	for _, v := range o {
		a = append(a, byte(v))
	}
	return string(a)
}

func (s SolutionsLeft) String() string {
	return "Numbers left: " + strconv.Itoa(s.left) + " - Solution: " + fmt.Sprint(s.sol)
}

func (t *Targets) String() string {
	return fmt.Sprint(*t)
}

func (n *Numbers) String() string {
	return fmt.Sprint(*n)
}

func Reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}
