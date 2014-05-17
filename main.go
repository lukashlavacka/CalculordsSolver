// Solver project main.go
package main

import (
	"flag"
	"fmt"
	"sort"
)

var targets Targets
var numbers Numbers

func main() {
	flag.Parse()

	if len(targets) == 0 {
		targets = Targets{15, 14, 13, 12}
	}

	if len(numbers) == 0 {
		numbers = Numbers{2, 3, 2, 13, 7, 2, 5, 3}
	}

	operations := Operations{'+', '-', '*'}
	actual := Assignment{targets, numbers, operations}.SolveAsync()

	fmt.Println(len(actual))

	sort.Sort(ByNums(actual))
	fmt.Println(actual[0])

	sort.Sort(BySize(actual))
	fmt.Println(actual[0])

	sort.Sort(BySum(actual))
	fmt.Println(actual[0])
}
