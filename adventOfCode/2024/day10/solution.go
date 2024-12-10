package day1

import (
	"os"
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

type point struct {
	x, y int
}

type peaks struct {
	points []point
	size   int
}

func (p *peaks) append(newPoint point) {
	for _, v := range p.points {
		if v.x == newPoint.x && v.y == newPoint.y {
			return
		}
	}
	p.size++
	p.points = append(p.points, newPoint)
}

func (p *peaks) append2(newPoint point) {
	p.size++
	p.points = append(p.points, newPoint)
}

func handlePart1(input string) int {
	grid := make([][]int, 0, len(input))
	lines := strings.Split(input, "\n")
	zeros := make([]point, 0)
	for i, row := range lines {
		if len(row) == 0 {
			continue
		}
		parsedRow := make([]int, 0, len(row))
		for k, sym := range row {
			num, err := strconv.Atoi(string(sym))
			check(err)
			if num == 0 {
				zeros = append(zeros, point{i, k})
			}
			parsedRow = append(parsedRow, num)
		}
		grid = append(grid, parsedRow)
	}

	allRes := make([]*peaks, 0)
	for _, v := range zeros {
		p := peaks{points: make([]point, 0)}
		allRes = append(allRes, &p)
		findPath(v.x, v.y, grid, p.append)
	}
	res := 0
	for _, v := range allRes {
		//fmt.Println(v)
		res += v.size
	}
	return res
}

func findPath(i, k int, grid [][]int, appender func(point)) int {
	if grid[i][k] == 9 {
		appender(point{i, k})
		return 1
	}

	res := 0
	if i+1 < len(grid) && canStep(grid, grid[i][k], i+1, k) {
		res += findPath(i+1, k, grid, appender)
	}

	if i-1 >= 0 && canStep(grid, grid[i][k], i-1, k) {
		res += findPath(i-1, k, grid, appender)
	}

	if k-1 >= 0 && canStep(grid, grid[i][k], i, k-1) {
		res += findPath(i, k-1, grid, appender)
	}

	if k+1 < len(grid[i]) && canStep(grid, grid[i][k], i, k+1) {
		res += findPath(i, k+1, grid, appender)
	}

	return res
}

func canStep(grid [][]int, currValue, i, k int) bool {
	if diff := grid[i][k] - currValue; diff == 1 {
		return true
	}
	return false
}

func handlePart2(input string) int {
	grid := make([][]int, 0, len(input))
	lines := strings.Split(input, "\n")
	zeros := make([]point, 0)
	for i, row := range lines {
		if len(row) == 0 {
			continue
		}
		parsedRow := make([]int, 0, len(row))
		for k, sym := range row {
			num, err := strconv.Atoi(string(sym))
			check(err)
			if num == 0 {
				zeros = append(zeros, point{i, k})
			}
			parsedRow = append(parsedRow, num)
		}
		grid = append(grid, parsedRow)
	}

	allRes := make([]*peaks, 0)
	for _, v := range zeros {
		p := peaks{points: make([]point, 0)}
		allRes = append(allRes, &p)
		findPath(v.x, v.y, grid, p.append2)
	}
	res := 0
	for _, v := range allRes {
		//fmt.Println(v)
		res += v.size
	}
	return res
}
