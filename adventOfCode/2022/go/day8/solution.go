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
	visibleTrees, _ := findNoOfTreesVisibleFromOutside(treeMap)
	fmt.Println(visibleTrees)
}

func part2() {
	lines := readFile("input")
	treeMap := parseMap(lines)
	_, maxVision := findNoOfTreesVisibleFromOutside(treeMap)
	fmt.Println(maxVision)
}

func findNoOfTreesVisibleFromOutside(treeMap [][]int) (int, int) {
	treesVisible := 0
	maxViewFactor := 0

	for i, row := range treeMap {
		for j, treeHeight := range row {
			visible, viewFactor := isVisible(treeHeight, i, j, treeMap)
			if visible {
				treesVisible++
			}
			if maxViewFactor < viewFactor {
				maxViewFactor = viewFactor
			}
		}
	}

	return treesVisible, maxViewFactor
}

func isVisible(treeHeight, i, j int, grid [][]int) (bool, int) {
	transponed := transpone(grid)

	leftPart := grid[i][:j]
	rightPart := grid[i][j+1:]

	topPart := transponed[j][:i]
	bottomPart := transponed[j][i+1:]

	visibleFromRight, rightView := isHigherThanSlice(treeHeight, rightPart, false)
	visibleFromLeft, leftView := isHigherThanSlice(treeHeight, leftPart, true)
	visibleFromTop, topView := isHigherThanSlice(treeHeight, topPart, true)
	visibleFromBottom, bottomView := isHigherThanSlice(treeHeight, bottomPart, false)
	res := visibleFromLeft || visibleFromRight || visibleFromTop || visibleFromBottom
	return res, rightView * leftView * topView * bottomView
}

func isHigherThanSlice(height int, other1 []int, invert bool) (bool, int) {
	other := make([]int, len(other1))
	copy(other, other1)

	if invert {
		for i, j := 0, len(other)-1; i < j; i, j = i+1, j-1 {
			other[i], other[j] = other[j], other[i]
		}
	}

	distance := 1
	for _, otherTreeHeight := range other {
		if height <= otherTreeHeight {
			return false, distance
		}
		distance++
	}
	return true, distance - 1
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
	//part1()
	part2()
}
