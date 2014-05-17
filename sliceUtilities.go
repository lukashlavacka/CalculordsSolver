// Few utilities for appending if unique, if slice contains element and slice without element
package main

import "errors"

func appendIfUnique(dst Array, src T) (S, error) {
	a := dst.ToS()
	for _, v := range a {
		if v.Eq(src) {
			return a, errors.New("Already in array")
		}
	}
	return append(a, src), nil
}

func appendIfUniqueArray(dst []S, src S) (U, error) {
	a := dst
	for _, v := range a {
		if v.Eq(src) {
			return a, errors.New("Already contains")
		}
	}
	return append(a, src), nil
}

func contains(array Array, element T) bool {
	for _, a := range array.ToS() {
		if a.Eq(element) {
			return true
		}
	}
	return false
}

func withoutOneGeneric(array S, element T) (S, bool) {
	var r S
	removed := false
	for _, a := range array {
		if !a.Eq(element) || removed {
			r = append(r, a)
		} else {
			removed = true
		}
	}
	return r, removed
}

func without(array Array, elements Array) S {
	aT := append(S{}, array.ToS()...)
	eT := append(S{}, elements.ToS()...)
	var removed bool
	for _, e := range eT {
		aT, removed = withoutOneGeneric(aT, e)
		if removed {
			eT, _ = withoutOneGeneric(eT, e)
		}
	}
	return aT
}
