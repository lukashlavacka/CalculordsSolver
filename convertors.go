// Functions to convert between typeDefs and interfaces
package main

func (a Numbers) ToS() S {
	r := make(S, len(a))
	for i, v := range a {
		r[i] = v
	}
	return r
}

func (a Solutions) ToS() S {
	r := make(S, len(a))
	for i, v := range a {
		r[i] = v
	}
	return r
}

func (a Targets) ToS() S {
	r := make(S, len(a))
	for i, v := range a {
		r[i] = v
	}
	return r
}

func (a OperationsA) ToS() S {
	r := make(S, len(a))
	for i, v := range a {
		r[i] = v
	}
	return r
}

func (a SolutionsA) ToS() S {
	r := make(S, len(a))
	for i, v := range a {
		r[i] = v
	}
	return r
}

func (s S) ToNumbers() Numbers {
	r := make(Numbers, len(s))
	for i, v := range s {
		r[i] = v.(cUint)
	}
	return r
}

func (s S) ToTargets() Targets {
	r := make(Targets, len(s))
	for i, v := range s {
		r[i] = v.(cInt)
	}
	return r
}

func (s S) ToSolutions() Solutions {
	r := make(Solutions, len(s))
	for i, v := range s {
		r[i] = v.(Solution)
	}
	return r
}

func (s S) ToOperations() Operations {
	r := make(Operations, len(s))
	for i, v := range s {
		r[i] = v.(cByte)
	}
	return r
}

func (u U) ToOperationsA() OperationsA {
	r := make(OperationsA, len(u))
	for i, v := range u {
		r[i] = v.ToOperations()
	}
	return r
}

func (a Operations) ToS() S {
	r := make(S, len(a))
	for i, v := range a {
		r[i] = v
	}
	return r
}

func (s S) ToS() S {
	return s
}
