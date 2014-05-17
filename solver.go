// solver
package main

import "errors"

func (s Solution) Eval() cInt {
	a := s.nums
	o := s.ops
	r := int(a[0])
	n := len(a)
	for i := 1; i < n; i++ {
		switch o[i-1] {
		case '+':
			r += int(a[i])
		case '-':
			r -= int(a[i])
		case '*':
			r *= int(a[i])
		}
	}
	return cInt(r)
}

func (s Solutions) Eval() cInt {
	r := cInt(0)
	for _, v := range s {
		r += v.Eval()
	}
	return cInt(r)
}

func (problem Problem) Solve(one bool) Solutions {
	var r S
	for powerset := range powerSetGenerator(problem.nums) {
		for permutation := range permutationsGenerator(powerset) {
			l := len(powerset)
			for operations := range multiSetGenerator(problem.ops, l-1) {
				possibleSolution := Solution{permutation.ToNumbers(), operations.ToOperations()}
				if possibleSolution.Eval() == problem.target {
					r, _ = appendIfUnique(r, Solution{permutation.ToNumbers(), operations.ToOperations()})
					if one {
						return r.ToSolutions()
					}
				}
			}
		}
	}
	return r.ToSolutions()
}

func (p Problem) SolveOne() (Solution, error) {
	s := p.Solve(true)
	if len(s) > 0 {
		return s[0], nil
	}
	return Solution{}, errors.New("No solution found")
}

func (a Assignment) SolveSpecificTargets() SolutionsLeft {
	leftNums := append(Numbers{}, a.nums...)
	var t S
	for _, target := range a.targets {
		problem := Problem{target, leftNums, a.ops}
		solution, err := problem.SolveOne()
		if err == nil {
			t = append(t, solution)
		} else {
			break
		}
		leftNums = without(leftNums, solution.nums).ToNumbers()
	}
	return SolutionsLeft{t.ToSolutions(), len(leftNums)}
}

func (a Assignment) SolveSpecificTargetsAsync(c chan SolutionsLeft) {
	c <- a.SolveSpecificTargets()
}

func AssignmentClone(targets Targets, nums Numbers, ops Operations) Assignment {
	return Assignment{
		append(Targets{}, targets...),
		append(Numbers{}, nums...),
		append(Operations{}, ops...),
	}
}

func (a Assignment) SolveAsync() SolutionsA {
	var r SolutionsA
	c := make(chan SolutionsLeft)
	j := 0
	for targets := range permutationsGenerator(a.targets) {
		aClone := AssignmentClone(targets.ToTargets(), a.nums, a.ops)
		go aClone.SolveSpecificTargetsAsync(c)
		j++
	}
	for i := 0; i < j; i++ {
		r = append(r, <-c)
	}
	return r
}
