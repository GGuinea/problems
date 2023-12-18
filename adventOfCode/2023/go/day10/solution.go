package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func parseLines(lines []string) [][]string {
	signs := make([][]string, 0)

	for _, line := range lines {
		split := strings.Split(line, "")
		signs = append(signs, split)
	}

	return signs
}

func findStartingPoint(lines [][]string) (int, int) {
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == "S" {
				return i, j
			}
		}
	}
	return -1, -1
}

func getNextMove(matrix [][]string, x, y, prevX, prevY int) (int, int) {
	if matrix[x][y] == "S" {
		if matrix[x+1][y] == "|" || matrix[x+1][y] == "L" || matrix[x+1][y] == "J" {
			return x + 1, y
		}

		if matrix[x-1][y] == "|" || matrix[x-1][y] == "F" || matrix[x-1][y] == "7" {
			return x - 1, y
		}

		if matrix[x][y+1] == "-" || matrix[x][y+1] == "J" || matrix[x][y+1] == "7" {
			return x, y + 1
		}

		if matrix[x][y-1] == "-" || matrix[x][y+1] == "F" || matrix[x][y+1] == "L" {
			return x, y - 1
		}
	}

	if matrix[x][y] == "-" {
		if prevY == y-1 {
			return x, y + 1
		} else if prevY == y+1 {
			return x, y - 1
		}
	}

	if matrix[x][y] == "|" {
		if prevX == x-1 {
			return x + 1, y
		} else if prevX == x+1 {
			return x - 1, y
		}
	}

	if matrix[x][y] == "F" {
		if prevX == x+1 {
			return x, y + 1
		} else if prevY == y+1 {
			return x + 1, y
		}
	}

	if matrix[x][y] == "L" {
		if prevX == x-1 {
			return x, y + 1
		} else if prevY == y+1 {
			return x - 1, y
		}
	}

	if matrix[x][y] == "J" {
		if prevX == x-1 {
			return x, y - 1
		} else if prevY == y-1 {
			return x - 1, y
		}
	}

	if matrix[x][y] == "7" {
		if prevX == x+1 {
			return x, y - 1
		} else if prevY == y-1 {
			return x + 1, y
		}
	}

	return -1, -1
}

func countDistance(matrix [][]string, startingX int, startingY int) (result int) {
	currentX := startingX
	currentY := startingY
	prevX := currentX
	prevY := currentX

	for {
		nextX, nextY := getNextMove(matrix, currentX, currentY, prevX, prevY)
		prevX = currentX
		prevY = currentY

		currentX = nextX
		currentY = nextY

		if currentX == startingX && currentY == startingY {
			break
		}
		result++
	}

	return (result / 2) + 1
}

func part1(filename string) int {
	lines := readFile(filename)
	parsed := parseLines(lines)
	x, y := findStartingPoint(parsed)
	dist := countDistance(parsed, x, y)
	return dist
}

func part2(filename string) int {
	return -1

}

func main() {
	fileName := os.Args[1]
	fmt.Println(part1(fileName))
	//fmt.Println(part2(fileName))
}
