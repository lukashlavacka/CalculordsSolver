package main

import (
	"fmt"
	"testing"
)

func Test_AppendIfUnique(t *testing.T) {
	input := Numbers{cUint(1), cUint(2), cUint(3)}
	expected1 := input
	expected2 := Numbers{cUint(1), cUint(2), cUint(3), cUint(4)}

	par1 := cUint(3)
	par2 := cUint(4)

	actual1, err1 := appendIfUnique(input, par1)
	actual2, err2 := appendIfUnique(actual1, par2)

	if err1 == nil {
		t.Errorf("appendIfUnique - inserting existing didn't return error")
	}
	if err2 != nil {
		t.Errorf("appendIfUnique - inserting new returned error")
	}
	if fmt.Sprint(actual1) != fmt.Sprint(expected1) {
		t.Errorf("appendIfUnique(%+v, %+v) = %+v, want %+v", input, par1, actual1, expected1)
	}
	if fmt.Sprint(actual2) != fmt.Sprint(expected2) {
		t.Errorf("appendIfUnique(%+v, %+v) = %+v, want %+v", actual1, par2, actual2, expected2)
	}
}

func Test_AppendIfUnique_Long(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	input := Numbers{cUint(1), cUint(2), cUint(3), cUint(4), cUint(5), cUint(6), cUint(7), cUint(8), cUint(9), cUint(10), cUint(11), cUint(12), cUint(13), cUint(14), cUint(15), cUint(16), cUint(17), cUint(18), cUint(19), cUint(20), cUint(21), cUint(22), cUint(23), cUint(24), cUint(25), cUint(26), cUint(27), cUint(28), cUint(29), cUint(30), cUint(31), cUint(32), cUint(33), cUint(34), cUint(35), cUint(36), cUint(37), cUint(38), cUint(39), cUint(40), cUint(41), cUint(42), cUint(43), cUint(44), cUint(45), cUint(46), cUint(47), cUint(48), cUint(49), cUint(50), cUint(51), cUint(52), cUint(53), cUint(54), cUint(55), cUint(56), cUint(57), cUint(58), cUint(59), cUint(60), cUint(61), cUint(62), cUint(63), cUint(64), cUint(65), cUint(66), cUint(67), cUint(68), cUint(69), cUint(70)}
	expected1 := input
	expected2 := Numbers{cUint(1), cUint(2), cUint(3), cUint(4), cUint(5), cUint(6), cUint(7), cUint(8), cUint(9), cUint(10), cUint(11), cUint(12), cUint(13), cUint(14), cUint(15), cUint(16), cUint(17), cUint(18), cUint(19), cUint(20), cUint(21), cUint(22), cUint(23), cUint(24), cUint(25), cUint(26), cUint(27), cUint(28), cUint(29), cUint(30), cUint(31), cUint(32), cUint(33), cUint(34), cUint(35), cUint(36), cUint(37), cUint(38), cUint(39), cUint(40), cUint(41), cUint(42), cUint(43), cUint(44), cUint(45), cUint(46), cUint(47), cUint(48), cUint(49), cUint(50), cUint(51), cUint(52), cUint(53), cUint(54), cUint(55), cUint(56), cUint(57), cUint(58), cUint(59), cUint(60), cUint(61), cUint(62), cUint(63), cUint(64), cUint(65), cUint(66), cUint(67), cUint(68), cUint(69), cUint(70), cUint(71)}

	par1 := cUint(3)
	par2 := cUint(71)

	actual1, err1 := appendIfUnique(input, par1)
	actual2, err2 := appendIfUnique(actual1, par2)

	if err1 == nil {
		t.Errorf("appendIfUnique - inserting existing didn't return error")
	}
	if err2 != nil {
		t.Errorf("appendIfUnique - inserting new returned error")
	}
	if fmt.Sprint(actual1) != fmt.Sprint(expected1) {
		t.Errorf("appendIfUnique(%+v, %+v) = %+v, want %+v", input, par1, actual1, expected1)
	}
	if fmt.Sprint(actual2) != fmt.Sprint(expected2) {
		t.Errorf("appendIfUnique(%+v, %+v) = %+v, want %+v", actual1, par2, actual2, expected2)
	}
}

func Test_AppendIfUniqueArray(t *testing.T) {
	input := make([]S, 2)
	input[0] = Operations{cByte('+'), cByte('-'), cByte('*')}.ToS()
	input[1] = Operations{cByte('+'), cByte('*')}.ToS()

	expected1 := input
	expected2 := make([]S, 3)
	expected2[0] = Operations{cByte('+'), cByte('-'), cByte('*')}.ToS()
	expected2[1] = Operations{cByte('+'), cByte('*')}.ToS()
	expected2[2] = Operations{cByte('+'), cByte('/')}.ToS()

	par1 := Operations{cByte('+'), cByte('*')}.ToS()
	par2 := Operations{cByte('+'), cByte('/')}.ToS()

	actual1, err1 := appendIfUniqueArray(input, par1)
	actual2, err2 := appendIfUniqueArray(actual1, par2)

	if err1 == nil {
		t.Errorf("appendIfUniqueArray - inserting existing didn't return error")
	}
	if err2 != nil {
		t.Errorf("appendIfUniqueArray - inserting new returned error")
	}
	if fmt.Sprint(actual1) != fmt.Sprint(expected1) {
		t.Errorf("appendIfUniqueArray(%+v, %+v) = %+v, want %+v", input, par1, actual1, expected1)
	}
	if fmt.Sprint(actual2) != fmt.Sprint(expected2) {
		t.Errorf("appendIfUniqueArray(%+v, %+v) = %+v, want %+v", actual1, par2, actual2, expected2)
	}
}

func Test_Contains_True(t *testing.T) {
	input := Operations{'+', '-', '*'}
	expected := true

	actual := contains(input, cByte('-'))

	if actual != expected {
		t.Errorf("contains(%+v, '-') = %+v, want %+v", input, actual, expected)
	}
}

func Test_Contains_False(t *testing.T) {
	input := Operations{'+', '-', '*'}
	expected := false

	actual := contains(input, cByte('/'))

	if actual != expected {
		t.Errorf("contains(%+v, '-') = %+v, want %+v", input, actual, expected)
	}
}

func Test_Without(t *testing.T) {
	input1, input2 := Numbers{1, 2, 3, 4, 4, 4, 5}, Numbers{2, 4, 4}
	expected := Numbers{1, 3, 4, 5}

	actual := without(input1, input2)

	if fmt.Sprint(actual) != fmt.Sprint(expected) {
		t.Errorf("withoutOne(%+v, %+v) = %+v, want %+v", input1, input2, actual, expected)
	}
}

func Test_WithoutOneGeneric(t *testing.T) {
	input1 := Numbers{1, 2, 3, 4, 4, 4, 5}.ToS()
	input2 := cUint(4)
	expected := Numbers{1, 2, 3, 4, 4, 5}

	actual, removed := withoutOneGeneric(input1, input2)

	if !removed {
		t.Errorf("withoutOneGeneric(%+v, %+v) = %+v, expected removed to be true", input1, input2, actual)
	}
	if fmt.Sprint(actual) != fmt.Sprint(expected) {
		t.Errorf("withoutOneGeneric(%+v, %+v) = %+v, want %+v", input1, input2, actual, expected)
	}
}
