package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/aclements/go-z3/z3"
	"github.com/kkevinchou/adventofcode/utils"
)

var file string = "input.txt"

func main() {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	parenthesis := regexp.MustCompile(`\((.*?)\)`)
	braces := regexp.MustCompile(`\{(.*?)\}`)
	lines := strings.Split(string(data), "\r\n")

	var result int
	for _, line := range lines {
		match := braces.FindString(line)
		match = strings.Trim(match, "{}")
		goal := utils.StringSliceToIntSlice(strings.Split(match, ","))

		ops := [][]int{}
		matches := parenthesis.FindAllString(line, -1)
		for _, match := range matches {
			match := strings.Trim(match, "()")
			ops = append(ops, utils.StringSliceToIntSlice(strings.Split(match, ",")))
		}

		best, _, ok := minPressesZ3(goal, ops)
		if !ok {
			fmt.Println("UNSAT: no solution")
			return
		}

		result += best
	}

	fmt.Println(result)
}

func minPressesZ3(target []int, buttons [][]int) (int, []int, bool) {
	m := len(target)  // counters
	n := len(buttons) // buttons

	cfg := z3.NewContextConfig()
	ctx := z3.NewContext(cfg)
	s := z3.NewSolver(ctx)

	intSort := ctx.IntSort()
	z0 := ctx.FromInt(0, intSort).(z3.Int)

	// Variables x0..x(n-1), all integers >= 0
	x := make([]z3.Int, n)
	for j := 0; j < n; j++ {
		x[j] = ctx.IntConst(fmt.Sprintf("x%d", j))
		s.Assert(x[j].GE(z0))
	}

	// Helper: does button j affect counter i?
	affects := func(j, i int) bool {
		for _, idx := range buttons[j] {
			if idx == i {
				return true
			}
		}
		return false
	}

	// Ax = b constraints: for each counter i, sum_j A[i,j]*xj == target[i]
	for i := 0; i < m; i++ {
		sum := ctx.FromInt(0, intSort).(z3.Int)
		for j := 0; j < n; j++ {
			if affects(j, i) {
				sum = sum.Add(x[j])
			}
		}
		ti := ctx.FromInt(int64(target[i]), intSort).(z3.Int)
		s.Assert(sum.Eq(ti))
	}

	// total = sum(x)
	total := ctx.FromInt(0, intSort).(z3.Int)
	for j := 0; j < n; j++ {
		total = total.Add(x[j])
	}

	// Upper bound: sum(target) always bounds total presses from above (loose but safe)
	hi := 0
	for _, v := range target {
		hi += v
	}
	lo := 0

	var best int
	bestFound := false
	var bestModel *z3.Model

	for lo <= hi {
		mid := (lo + hi) / 2

		s.Push()
		s.Assert(total.LE(ctx.FromInt(int64(mid), intSort).(z3.Int)))

		sat, err := s.Check()
		if err != nil {
			return 0, nil, false
		}

		if sat {
			bestFound = true
			best = mid
			bestModel = s.Model()
			hi = mid - 1
		} else {
			lo = mid + 1
		}

		s.Pop()
	}

	if !bestFound {
		return 0, nil, false
	}

	// Extract solution
	sol := make([]int, n)
	for j := 0; j < n; j++ {
		v := bestModel.Eval(x[j], true).(z3.Int)
		val, _, ok2 := v.AsInt64()
		if !ok2 {
			return 0, nil, false
		}
		sol[j] = int(val)
	}
	return best, sol, true
}
