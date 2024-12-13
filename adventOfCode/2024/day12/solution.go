package day1

import (
	"fmt"
	"math"
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

type region struct {
	symbol     string
	plants     []*plant
	symCounter int
	fenceSize  int
}

type plant struct {
	x, y int
}

func (r *region) addNewPoint(x, y int) {
	for _, v := range r.plants {
		if v.x == x && v.y == y {
			return
		}

		diffX := v.x - x
		diffY := v.y - y
		if y == v.y && (diffX == 1 || diffX == -1) {
			r.fenceSize -= 2
		}
		if x == v.x && (diffY == 1 || diffY == -1) {
			r.fenceSize -= 2
		}
	}

	r.symCounter++
	r.fenceSize += 4
	r.plants = append(r.plants, &plant{x, y})
}

func handlePart1(input string) int {
	uniqueSymbols := make(map[string]struct{})
	grid := make([][]string, 0)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		symbols := strings.Split(line, "")
		newRow := make([]string, 0, len(symbols))
		for _, symbol := range symbols {
			newRow = append(newRow, symbol)
			uniqueSymbols[symbol] = struct{}{}
		}
		grid = append(grid, newRow)
	}

	regions := make([]*region, 0)
	for i := range grid {
		for k := range grid {
			if grid[i][k] == "-1" {
				continue
			}
			reg := &region{symbol: grid[i][k]}
			findAll(grid[i][k], grid, i, k, reg)
			regions = append(regions, reg)
		}
	}

	res := 0
	for _, v := range regions {
		res += v.fenceSize * v.symCounter
	}
	return res
}

func handlePart2(input string) int {
	return 0
}

func findAll(symbol string, grid [][]string, x int, y int, reg *region) {
	if x < 0 || x >= len(grid) {
		return
	}

	if y < 0 || y >= len(grid[x]) {
		return
	}

	if grid[x][y] == "-1" {
		return
	}

	if grid[x][y] != symbol {
		return
	}

	grid[x][y] = "-1"
	reg.addNewPoint(x, y)
	findAll(symbol, grid, x-1, y, reg)
	findAll(symbol, grid, x+1, y, reg)
	findAll(symbol, grid, x, y-1, reg)
	findAll(symbol, grid, x, y+1, reg)
}
