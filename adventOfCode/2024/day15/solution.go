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

func handlePart1(input string) int {
	lines := strings.Split(input, "\n")
	grid := make([][]string, 0)
	var moves string

	var gridCompleted bool

	robotX, robotY := 0, 0

	for i, line := range lines {
		if len(line) == 0 {
			gridCompleted = true
		}

		if !gridCompleted {
			if robotX == 0 && robotY == 0 {
				for k, v := range line {
					if v == '@' {
						robotX = i
						robotY = k
					}
				}
			}
			grid = append(grid, strings.Split(line, ""))
		} else {
			moves += line
		}
	}

	for _, m := range moves {
		res := move(robotX, robotY, grid, m)
		if res.set {
			robotX = res.newX
			robotY = res.newY
		}
		//printGrid(grid)
	}

	res := 0
	for i := range grid {
		for k := range grid[i] {
			if grid[i][k] == "O" {
				res += i*100 + k
			}
		}
	}

	return res
}

func printGrid(grid [][]string) {
	for i := range grid {
		for k := range grid[i] {
			fmt.Print(grid[i][k])
		}
		fmt.Println()
	}
}

type moveFuncRes struct {
	newX, newY int
	set        bool
}

func move(currentX, currentY int, grid [][]string, move rune) moveFuncRes {
	moveX, moveY := 0, 0
	switch move {
	case '^':
		moveX, moveY = -1, 0
	case '>':
		moveX, moveY = 0, 1
	case 'v':
		moveX, moveY = 1, 0
	case '<':
		moveX, moveY = 0, -1
	default:
		panic(move)
	}

	newX := currentX + moveX
	newY := currentY + moveY

	if newX >= len(grid) || newX < 0 {
		return moveFuncRes{}
	}
	if newY >= len(grid[0]) || newY < 0 {
		return moveFuncRes{}
	}

	if grid[newX][newY] == "#" {
		return moveFuncRes{}
	}

	if grid[newX][newY] == "." {
		grid[newX][newY] = "@"
		grid[currentX][currentY] = "."
		return moveFuncRes{newX, newY, true}
	}

	if grid[newX][newY] == "O" {
		newBoxPos := canPush(currentX, currentY, moveX, moveY, grid)
		if !newBoxPos.found {
			return moveFuncRes{}
		}
		grid[newX][newY] = "@"
		grid[currentX][currentY] = "."
		grid[newBoxPos.x][newBoxPos.y] = "O"

		return moveFuncRes{newX, newY, true}
	}
	return moveFuncRes{}
}

type newBoxPos struct {
	x, y  int
	found bool
}

func canPush(currX, currY, moveX, moveY int, grid [][]string) newBoxPos {
	newX := currX
	newY := currY
	for {
		newX += moveX
		newY += moveY
		if grid[newX][newY] == "#" {
			return newBoxPos{}
		}
		if grid[newX][newY] == "O" {
			continue
		}
		if grid[newX][newY] == "." {
			return newBoxPos{x: newX, y: newY, found: true}
		}
	}
}
