package day1

import (
	"fmt"
	"os"
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
	lines := strings.Split(input, "\n")
	var res int
	for i := 0; i < len(lines); {
		if len(lines[i]) == 0 {
			i++
			continue
		}

		var ax, ay, bx, by, px, py int
		n, err := fmt.Sscanf(lines[i], "Button A: X+%d, Y+%d", &ax, &ay)
		check(err)

		if n != 2 {
			panic(lines[i])
		}
		i++
		n, err = fmt.Sscanf(lines[i], "Button B: X+%d, Y+%d", &bx, &by)
		check(err)
		if n != 2 {
			panic(lines[i])
		}
		i++

		n, err = fmt.Sscanf(lines[i], "Prize: X=%d, Y=%d", &px, &py)
		check(err)
		i++
		//fmt.Println(ax, ay, bx, by, px, py)
		res += calculatePrice(ax, ay, bx, by, px, py, true)
	}

	return res
}

func calculatePrice(ax, ay, bx, by, px, py int, p1 bool) int {
	A := (px*by - py*bx) / (ax*by - ay*bx)
	B := (px - ax*A) / bx

	if p1 {
		if A > 100 || B > 100 {
			return 0
		}
	}

	intA := int(A)
	intB := int(B)

	if (intA*ax + intB*bx) != px {
		return 0
	}

	if (A*ay + B*by) != py {
		return 0
	}

	return 3*intA + intB
}

func handlePart2(input string) int {
	lines := strings.Split(input, "\n")
	var res int
	for i := 0; i < len(lines); {
		if len(lines[i]) == 0 {
			i++
			continue
		}

		var ax, ay, bx, by, px, py int
		n, err := fmt.Sscanf(lines[i], "Button A: X+%d, Y+%d", &ax, &ay)
		check(err)

		if n != 2 {
			panic(lines[i])
		}
		i++
		n, err = fmt.Sscanf(lines[i], "Button B: X+%d, Y+%d", &bx, &by)
		check(err)
		if n != 2 {
			panic(lines[i])
		}
		i++

		n, err = fmt.Sscanf(lines[i], "Prize: X=%d, Y=%d", &px, &py)
		check(err)
		i++
		res += calculatePrice(ax, ay, bx, by, px+10000000000000, py+10000000000000, false)
	}

	return res
}

