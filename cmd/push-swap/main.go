package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var lesOnze = []string{"sa", "sb", "ss", "pa", "pb", "ra", "rb", "rr", "rra", "rrb", "rrr"}

type etat struct {
	a, b   []int
	chemin []string
}

func parseArgs(args []string) ([]int, error) {
	var nums []int
	seen := make(map[int]bool)
	for _, arg := range args {
		tokens := strings.Fields(arg)
		for _, token := range tokens {
			n64, err := strconv.ParseInt(token, 10, 64)
			if err != nil || n64 > math.MaxInt32 || n64 < math.MinInt32 {
				return nil, fmt.Errorf("error")
			}
			val := int(n64)
			if seen[val] {
				return nil, fmt.Errorf("error")
			}
			seen[val] = true
			nums = append(nums, val)
		}
	}
	return nums, nil
}

func copier(s []int) []int {
	res := make([]int, len(s))
	copy(res, s)
	return res
}

func estTriee(a, b []int) bool {
	if len(b) > 0 {
		return false
	}
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

func appliquer(op string, aIn, bIn []int) ([]int, []int) {
	a := copier(aIn)
	b := copier(bIn)

	switch op {
	case "sa":
		if len(a) > 1 {
			a[0], a[1] = a[1], a[0]
		}
	case "sb":
		if len(b) > 1 {
			b[0], b[1] = b[1], b[0]
		}
	case "ss":
		if len(a) > 1 {
			a[0], a[1] = a[1], a[0]
		}
		if len(b) > 1 {
			b[0], b[1] = b[1], b[0]
		}
	case "pa":
		if len(b) > 0 {
			a = append([]int{b[0]}, a...)
			b = b[1:]
		}
	case "pb":
		if len(a) > 0 {
			b = append([]int{a[0]}, b...)
			a = a[1:]
		}
	case "ra":
		if len(a) > 1 {
			a = append(a[1:], a[0])
		}
	case "rb":
		if len(b) > 1 {
			b = append(b[1:], b[0])
		}
	case "rr":
		if len(a) > 1 {
			a = append(a[1:], a[0])
		}
		if len(b) > 1 {
			b = append(b[1:], b[0])
		}
	case "rra":
		if len(a) > 1 {
			a = append([]int{a[len(a)-1]}, a[:len(a)-1]...)
		}
	case "rrb":
		if len(b) > 1 {
			b = append([]int{b[len(b)-1]}, b[:len(b)-1]...)
		}
	case "rrr":
		if len(a) > 1 {
			a = append([]int{a[len(a)-1]}, a[:len(a)-1]...)
		}
		if len(b) > 1 {
			b = append([]int{b[len(b)-1]}, b[:len(b)-1]...)
		}
	}
	return a, b
}

func forceBrute(depart []int) []string {
	file := []etat{{depart, []int{}, nil}}
	vu := map[string]bool{fmt.Sprint(depart, []int{}): true}

	for len(file) > 0 {
		e := file[0]
		file = file[1:]

		if estTriee(e.a, e.b) {
			return e.chemin
		}

		for _, op := range lesOnze {
			a, b := appliquer(op, e.a, e.b)
			cle := fmt.Sprint(a, b)
			if !vu[cle] {
				vu[cle] = true
				chemin := append(append([]string{}, e.chemin...), op)
				file = append(file, etat{a, b, chemin})
			}
		}
	}
	return nil
}

func calculerRangs(depart []int) []int {
	copie := copier(depart)
	sort.Ints(copie)

	rangs := make([]int, len(depart))
	for i, val := range depart {
		rangs[i] = sort.SearchInts(copie, val)
	}
	return rangs
}

func triRadix(depart []int) []string {
	a := copier(depart)
	var b []int
	var ops []string
	n := len(a)

	nbBits := 0
	for (n-1)>>nbBits != 0 {
		nbBits++
	}

	for bit := 0; bit < nbBits; bit++ {
		for i := 0; i < n; i++ {
			if (a[0]>>bit)&1 == 0 {
				a, b = appliquer("pb", a, b)
				ops = append(ops, "pb")
			} else {
				a, b = appliquer("ra", a, b)
				ops = append(ops, "ra")
			}
		}
		for len(b) > 0 {
			a, b = appliquer("pa", a, b)
			ops = append(ops, "pa")
		}
	}
	return ops
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	nombres, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	if estTriee(nombres, []int{}) {
		return
	}

	var ops []string
	if len(nombres) <= 6 {
		ops = forceBrute(nombres)
	} else {
		rangs := calculerRangs(nombres)
		ops = triRadix(rangs)
	}

	for _, op := range ops {
		fmt.Println(op)
	}
}
