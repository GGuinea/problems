package day1

import (
	"fmt"
	"os"
	"slices"
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
	dir  direction
}

type road struct {
	points []point
}

type direction struct {
	dirName, x, y int
}

const (
	Up = iota
	Right
	Down
	Left
)

func nextDirection(dir direction) direction {
	switch dir.dirName {
	case Up:
		return direction{Right, 0, 1}
	case Right:
		return direction{Down, 1, 0}
	case Down:
		return direction{Left, 0, -1}
	case Left:
		return direction{Up, -1, 0}
	default:
		panic("oh oh")
	}
}

func (r *road) append(p point) bool {
	for _, pk := range r.points {
		if pk.x == p.x && pk.y == p.y {
			return false
		}
	}
	r.points = append(r.points, p)
	return true
}

func (r *road) append2(p point) bool {
	for _, pk := range r.points {
		if pk.x == p.x && pk.y == p.y && pk.dir.dirName == p.dir.dirName {
			return false
		}
	}
	r.points = append(r.points, p)
	return true
}

func handlePart1(input string) int {
	startX, startY := 0, 0
	grid := make([][]string, 0)
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		chars := strings.Split(line, "")
		idx := slices.Index(chars, "^")
		if idx != -1 {
			startX = i
			startY = idx
		}
		grid = append(grid, chars)
	}

	road := road{
		points: make([]point, 0),
	}
	currDir := direction{Up, -1, 0}

	for {
		fmt.Println("doing", startX, startY)
		if grid[startX][startY] == "#" {
			//fmt.Println("switching;", startX, startY, grid[startX][startY])
			startX -= currDir.x
			startY -= currDir.y
			currDir = nextDirection(currDir)
			continue
		}
		road.append(point{x: startX, y: startY})
		startX += currDir.x
		startY += currDir.y
		if startX >= len(grid) || startX < 0 {
			fmt.Println("breaking x")
			break
		}
		if startY >= len(grid[startX]) || startY < 0 {
			fmt.Println("breaking y")
			break
		}
	}
	return len(road.points)
}

func handlePart2(input string) int {
	startX, startY := 0, 0
	grid := make([][]string, 0)
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		chars := strings.Split(line, "")
		idx := slices.Index(chars, "^")
		if idx != -1 {
			startX = i
			startY = idx
		}
		grid = append(grid, chars)
	}

	road := road{
		points: make([]point, 0),
	}
	currDir := direction{Up, -1, 0}
	startX2, startY2 := startX, startY

	for {
		if grid[startX][startY] == "#" {
			//fmt.Println("switching;", startX, startY, grid[startX][startY])
			startX -= currDir.x
			startY -= currDir.y
			currDir = nextDirection(currDir)
			continue
		}
		road.append(point{x: startX, y: startY, dir: currDir})
		startX += currDir.x
		startY += currDir.y
		if startX >= len(grid) || startX < 0 {
			fmt.Println("breaking x")
			break
		}
		if startY >= len(grid[startX]) || startY < 0 {
			fmt.Println("breaking y")
			break
		}
	}

	fmt.Println("start searching")

	res := 0
	for _, v := range road.points {
		grid[v.x][v.y] = "#"
		if isLoop(grid, startX2, startY2) {
			res++
		}
		grid[v.x][v.y] = "."
	}
	return res
}

func isLoop(grid [][]string, startX, startY int) bool {
	road := road{
		points: make([]point, 0),
	}
	currDir := direction{Up, -1, 0}

	for {
		if grid[startX][startY] == "#" {
			startX -= currDir.x
			startY -= currDir.y
			currDir = nextDirection(currDir)
			continue
		}
		ok := road.append2(point{x: startX, y: startY, dir: currDir})
		if !ok {
			return true
		}
		startX += currDir.x
		startY += currDir.y
		if startX >= len(grid) || startX < 0 {
			break
		}
		if startY >= len(grid[startX]) || startY < 0 {
			break
		}
	}

	return false
}
