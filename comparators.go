// Functions implementing T interface Eq method for comparing two objects of the same type
package main

func (a Solution) Eq(b T) bool {
	t := b.(Solution)
	return a.nums.Eq(t.nums) && a.ops.Eq(t.ops)
}

func (a Numbers) Eq(b T) bool {
	t := b.(Numbers)
	if len(a) != len(t) {
		return false
	}
	for i := range t {
		if a[i] != t[i] {
			return false
		}
	}
	return true
}

func (a SolutionsLeft) Eq(b T) bool {
	t := b.(SolutionsLeft)
	return a.sol.Eq(t.sol) && a.left == t.left
}

func (a Operations) Eq(b T) bool {
	t := b.(Operations)
	if len(a) != len(t) {
		return false
	}
	for i := range t {
		if a[i] != t[i] {
			return false
		}
	}
	return true
}

func (a Solutions) Eq(b T) bool {
	t := b.(Solutions)
	if len(a) != len(t) {
		return false
	}
	for i := range t {
		if !a[i].Eq(t[i]) {
			return false
		}
	}
	return true
}

func (a cUint) Eq(b T) bool {
	return a == b.(cUint)
}

func (a cInt) Eq(b T) bool {
	return a == b.(cInt)
}

func (a cByte) Eq(b T) bool {
	return a == b.(cByte)
}

func (a S) Eq(b S) bool {
	for i := range b {
		if !a[i].Eq(b[i]) {
			return false
		}
	}
	return true
}
