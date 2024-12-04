package day1

import (
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

type direction struct {
	x, y int
}

func handlePart1(input string) int {
	directions := []direction{
		{
			x: 0,
			y: 1,
		},
		{
			x: 1,
			y: 0,
		},
		{
			x: 1,
			y: 1,
		},
		{
			x: 0,
			y: -1,
		},
		{
			x: -1,
			y: 0,
		},
		{
			x: -1,
			y: -1,
		},
		{
			x: -1,
			y: 1,
		},
		{
			x: 1,
			y: -1,
		},
	}

	splitGrid := strings.Split(input, "\n")
	grid := make([][]string, 0, len(splitGrid))
	for _, line := range splitGrid {
		if len(line) == 0 {
			continue
		}
		chars := strings.Split(line, "")
		grid = append(grid, chars)
	}
	res := 0

	for i := range len(grid) {
		for k := range len(grid[0]) {
			if grid[i][k] == "X" {
				for _, dir := range directions {
					res += search(i+dir.x, k+dir.y, grid, "MAS", dir)
				}
			}
		}
	}

	return res
}

func handlePart2(input string) int {
	splitGrid := strings.Split(input, "\n")
	grid := make([][]string, 0, len(splitGrid))
	for _, line := range splitGrid {
		if len(line) == 0 {
			continue
		}
		chars := strings.Split(line, "")
		grid = append(grid, chars)
	}
	res := 0
	// M . S
	// . A .
	// M . S

	for i := range len(grid) {
		for k := range len(grid[0]) {
			if i <= 0 || i > len(grid)-2 {
				continue
			}

			if k <= 0 || k > len(grid[i])-2 {
				continue
			}
			// M . S
			// . A .
			// M . S
			if grid[i][k] == "A" && grid[i-1][k-1] == "M" && grid[i-1][k+1] == "S" && grid[i+1][k+1] == "S" && grid[i+1][k-1] == "M" {
				res++
			}

			// S . M
			// . A .
			// S . M
			if grid[i][k] == "A" && grid[i-1][k-1] == "S" && grid[i-1][k+1] == "M" && grid[i+1][k+1] == "M" && grid[i+1][k-1] == "S" {
				res++
			}

			// S . S
			// . A .
			// M . M
			if grid[i][k] == "A" && grid[i-1][k-1] == "S" && grid[i-1][k+1] == "S" && grid[i+1][k+1] == "M" && grid[i+1][k-1] == "M" {
				res++
			}
			// M . M
			// . A .
			// S . S
			if grid[i][k] == "A" && grid[i-1][k-1] == "M" && grid[i-1][k+1] == "M" && grid[i+1][k+1] == "S" && grid[i+1][k-1] == "S" {
				res++
			}
		}
	}

	return res
}

func search(startX, startY int, grid [][]string, searchedString string, dir direction) int {
	if len(searchedString) == 0 {
		return 1
	}

	if startX < 0 || startX >= len(grid) {
		return 0
	}

	if startY < 0 || startY >= len(grid[startX]) {
		return 0
	}

	if grid[startX][startY] == string(searchedString[0]) {
		return search(startX+dir.x, startY+dir.y, grid, searchedString[1:], dir)
	}
	return 0
}
