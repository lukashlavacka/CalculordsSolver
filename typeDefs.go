// type definitions of structs on wrappers
package main

type S []T
type U []S

type cUint uint8
type cInt int8
type cByte byte
type Numbers []cUint
type Targets []cInt
type Operations []cByte
type OperationsA []Operations
type Solution struct {
	nums Numbers
	ops  Operations
}
type Solutions []Solution
type SolutionsLeft struct {
	sol  Solutions
	left int
}
type SolutionsA []SolutionsLeft
type Problem struct {
	target cInt
	nums   Numbers
	ops    Operations
}
type Assignment struct {
	targets Targets
	nums    Numbers
	ops     Operations
}
