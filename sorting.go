// Functions for sorting Solutions array by different factors
package main

type BySize SolutionsA

func (a BySize) Len() int {
	return len(a)
}
func (a BySize) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a BySize) Less(i, j int) bool {
	if a[i].left == a[j].left {
		return a[i].sol.Eval() > a[j].sol.Eval()
	}
	return a[i].left < a[j].left
}

type BySum SolutionsA

func (a BySum) Len() int {
	return len(a)
}
func (a BySum) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a BySum) Less(i, j int) bool {
	return a[i].sol.Eval() > a[j].sol.Eval()
}

type ByNums SolutionsA

func (a ByNums) Len() int {
	return len(a)
}
func (a ByNums) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByNums) Less(i, j int) bool {
	if a[i].left == a[j].left {
		return len(a[i].sol) < len(a[j].sol)
	}
	return a[i].left < a[j].left
}
