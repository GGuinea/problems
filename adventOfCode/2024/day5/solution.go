package day1

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(filename string) string {
	dat, err := os.ReadFile(filename)
	check(err)
	return string(dat)
}

func part1() int {
	input := readFile("input")
	return handlePart1(input)
}

func part2() int {
	input := readFile("input")
	return handlePart2(input)
}

func handlePart1(input string) int {
	split := strings.Split(input, "\n")
	rules := make(map[int][]int)
	updates := make([][]int, 0)
	areUpdates := false
	for _, rule := range split {
		if !areUpdates {
			if len(rule) == 0 {
				areUpdates = true
				continue
			}
			var one, two int
			n, err := fmt.Sscanf(rule, "%d|%d", &one, &two)
			check(err)
			if n != 2 {
				panic("not expect no of inputs")
			}
			rules[one] = append(rules[one], two)
			continue
		}
		updateSplit := strings.Split(rule, ",")
		if len(updateSplit) < 2 {
			break
		}
		updates = append(updates, getIntArr(updateSplit))
	}

	res := 0

	for _, update := range updates {
		if isProperOrder(update, rules) {
			res += update[len(update)/2]
		}
	}

	return res
}

func handlePart2(input string) int {
	split := strings.Split(input, "\n")
	rules := make(map[int][]int)
	updates := make([][]int, 0)
	areUpdates := false
	for _, rule := range split {
		if !areUpdates {
			if len(rule) == 0 {
				areUpdates = true
				continue
			}
			var one, two int
			n, err := fmt.Sscanf(rule, "%d|%d", &one, &two)
			check(err)
			if n != 2 {
				panic("not expect no of inputs")
			}
			rules[one] = append(rules[one], two)
			continue
		}
		updateSplit := strings.Split(rule, ",")
		if len(updateSplit) < 2 {
			break
		}
		updates = append(updates, getIntArr(updateSplit))
	}

	res := 0

	for _, update := range updates {
		if isProperOrder(update, rules) {
			continue
		}

		cpy := make([]int, len(update))
		copy(cpy, update)
		slices.SortFunc(cpy, func(a, b int) int {
			rulesA := rules[a]
			rulesB := rules[b]

			if slices.Contains(rulesA, b) {
				return -1
			} else if slices.Contains(rulesB, a) {
				return 1
			}
			return 0
		})

		res += cpy[len(cpy)/2]
	}

	return res

}

func isProperOrder(update []int, rules map[int][]int) bool {
	for i := 0; i < len(update); i++ {
		rule := rules[update[i]]
		beforePart := update[0:i]
		for _, v := range beforePart {
			if slices.Contains(rule, v) {
				return false
			}
		}
	}
	return true
}

func getIntArr(arr []string) []int {
	res := make([]int, 0, len(arr))
	for _, v := range arr {
		parsed, err := strconv.Atoi(v)
		check(err)
		res = append(res, parsed)
	}
	return res
}
