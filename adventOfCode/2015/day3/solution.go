package solution

import (
	"os"
)

func readFile(filename string) string {
	res, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(res)
}

func part1() int {
	input := readFile("input")
	return handlePart1(input)
}

type visited struct {
	x, y int
}

func handlePart1(input string) int {
	res := make([]visited, 0, 100)
	x, y := 0, 0
	res = append(res, visited{x, y})
	for _, sign := range input {
		switch sign {
		case '^':
			y++
		case 'v':
			y--
		case '>':
			x++
		case '<':
			x--
		}
		res = tryToAdd(res, visited{x, y})
	}
	return len(res)
}

func tryToAdd(res []visited, newOne visited) []visited {
	for _, house := range res {
		if house.x == newOne.x && house.y == newOne.y {
			return res
		}
	}

	return append(res, newOne)
}

func part2() int {
	input := readFile("input")
	return handlePart2(input)
}

func handlePart2(input string) int {
	res1 := make([]visited, 0, 100)
	x1, y1 := 0, 0
	x2, y2 := 0, 0

	res1 = append(res1, visited{x1, y1})

	for idx, sign := range input {
		addToX, addToY := 0, 0
		switch sign {
		case '^':
			addToY = 1
		case 'v':
			addToY = -1
		case '>':
			addToX = 1
		case '<':
			addToX = -1
		}

		if idx%2 == 1 {
			x1 += addToX
			y1 += addToY
			res1 = tryToAdd(res1, visited{x1, y1})
		} else {
			x2 += addToX
			y2 += addToY
			res1 = tryToAdd(res1, visited{x2, y2})
		}

	}
	return len(res1)
}
