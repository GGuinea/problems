package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func part1() {
	lines := readFile("input")
	treeMap := parseMap(lines)
	visibleTrees := findNoOfTreesVisibleFromOutside(treeMap)
	fmt.Println(visibleTrees)
}

func findNoOfTreesVisibleFromOutside(treeMap [][]int) int {
	treesVisible := 0

	for i, row := range treeMap {
		for j, treeHeight := range row {
			if i == 0 || j == 0 || i == len(treeMap)-1 || j == len(treeMap)-1 {
				treesVisible++
				continue
			}
			visible := isVisible(treeHeight, i, j, treeMap)
			if visible {
				treesVisible++
			}
		}
	}

	return treesVisible
}

func isVisible(treeHeight, i, j int, grid [][]int) bool {
	rightPart := grid[i][j+1:]
	leftPart := grid[i][:j]
	transoned := transpone(grid)

	col := make([]int, len(grid[0]))

	for c := range grid {
		col[c] = grid[c][j]
	}

	bottomPart := transoned[j][:i]
	topPart := transoned[j][i+1:]
	visibleFromRight := isHigherThanSlice(treeHeight, rightPart)
	visibleFromLeft := isHigherThanSlice(treeHeight, leftPart)
	visibleFromTop := isHigherThanSlice(treeHeight, topPart)
	visibleFromBottom := isHigherThanSlice(treeHeight, bottomPart)
	res := visibleFromLeft || visibleFromRight || visibleFromTop || visibleFromBottom
	return res
}

func isHigherThanSlice(height int, other []int) bool {
	if len(other) == 0 {
		return true
	}
	for _, otherTreeHeight := range other {
		if height <= otherTreeHeight {
			return false
		}
	}
	return true
}

func transpone(grid [][]int) [][]int {
	newGrid := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			newGrid[j] = append(newGrid[j], grid[i][j])
		}
	}
	return newGrid
}

func parseMap(lines []string) [][]int {
	treeMap := make([][]int, len(lines))
	for i := range treeMap {
		treeMap[i] = make([]int, len(lines))
	}

	for i, line := range lines {
		splitted := strings.Split(line, "")

		for j, lett := range splitted {
			converted, err := strconv.Atoi(lett)
			if err != nil {
				panic(2)
			}
			treeMap[i][j] = converted
		}
	}

	return treeMap
}

func main() {
	part1()
}
