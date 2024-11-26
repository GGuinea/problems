package solution

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type point struct {
	x, y int
}

func readFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(1)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	lines := make([]string, 0)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func part1(filename string) int {
	return handlePart1(readFile(filename), 1)
}

func part2(filename string) int {
	return handlePart1(readFile(filename), 999999)
}

func handlePart2(input []string, spaceIndicator int) int {
	grid := make([][]string, len(input))
	for i := range len(input) {
		split := strings.Split(input[i], "")
		grid[i] = split
	}

	rows, cols, points := getEmptyRowsAndCols(input)
	fmt.Println(rows, cols)

	slices.SortFunc(points, func(one, two point) int {
		if one.y > two.y {
			return 1
		}
		if one.y < two.y {
			return -1
		}
		if one.x > two.x {
			return 1
		}

		if one.x < two.x {
			return -1
		}

		return 0
	})

	res := 0

	for i := 0; i < len(points); i++ {
		for k := i + 1; k < len(points); k++ {
			res += spaceDistance(points[i], points[k], rows, cols, spaceIndicator)
		}
	}

	return res
}

func handlePart1(input []string, spaceIndicator int) int {
	grid := make([][]string, len(input))
	for i := range len(input) {
		split := strings.Split(input[i], "")
		grid[i] = split
	}

	rows, cols, points := getEmptyRowsAndCols(input)
	fmt.Println(rows, cols)

	slices.SortFunc(points, func(one, two point) int {
		if one.y > two.y {
			return 1
		}
		if one.y < two.y {
			return -1
		}
		if one.x > two.x {
			return 1
		}

		if one.x < two.x {
			return -1
		}

		return 0
	})

	res := 0

	for i := 0; i < len(points); i++ {
		for k := i + 1; k < len(points); k++ {
			res += spaceDistance(points[i], points[k], rows, cols, spaceIndicator)
		}
	}

	return res
}

func spaceDistance(one, two point, rows, cols []int, spaceIndicator int) int {
	res := distBetween(one, two)

	if one.x > two.x {
		tmp := one
		one = two
		two = tmp
	}

	for _, row := range rows {
		if one.x < row && two.x > row {
			res += spaceIndicator
		}
	}

	if one.y > two.y {
		tmp := one
		one = two
		two = tmp
	}

	for _, col := range cols {
		if one.y < col && two.y > col {
			res += spaceIndicator
		}
	}

	return res
}

func distBetween(one, two point) int {
	return int(math.Abs(float64(one.x-two.x))) + int(math.Abs(float64(one.y-two.y)))
}

func getEmptyRowsAndCols(input []string) ([]int, []int, []point) {
	var rows []int
	var cols []int
	var points []point

	for i := range len(input) {
		skip := false
		for k := range len(input[i]) {
			if input[i][k] == '#' {
				points = append(points, point{i, k})
				skip = true
			}
			if k == len(input[i])-1 && !skip {
				rows = append(rows, i)
			}
		}
	}

	for i := range len(input) {
		for k := range len(input[i]) {
			if input[k][i] == '#' {
				break
			}
			if k == len(input[i])-1 {
				cols = append(cols, i)
			}
		}
	}

	return rows, cols, points
}
