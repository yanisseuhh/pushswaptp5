package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseArgs(args []string) ([]int, error) {
	var nums []int
	seen := make(map[int]bool)
	for _, arg := range args {
		for _, token := range strings.Fields(arg) {
			n, err := strconv.ParseInt(token, 10, 64)
			if err != nil || n > math.MaxInt32 || n < math.MinInt32 {
				return nil, fmt.Errorf("error")
			}
			val := int(n)
			if seen[val] {
				return nil, fmt.Errorf("error")
			}
			seen[val] = true
			nums = append(nums, val)
		}
	}
	return nums, nil
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

func main() {
	if len(os.Args) < 2 {
		return
	}

	a, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	var b []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		op := strings.TrimRight(scanner.Text(), "\r ")
		if op == "" {
			continue
		}
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
		default:
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}

	if estTriee(a, b) {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
