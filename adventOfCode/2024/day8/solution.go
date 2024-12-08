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

type point struct {
	x, y int
}

type antinodes struct {
	size           int
	limitX, limitY int
	nodes          []point
}

func (a *antinodes) append(p point) {
	if p.x < 0 || p.y < 0 {
		return
	}
	if p.x >= a.limitX || p.y >= a.limitY {
		return
	}
	for _, node := range a.nodes {
		if node.x == p.x && node.y == p.y {
			return
		}
	}
	a.size++
	a.nodes = append(a.nodes, p)
}

func handlePart1(input string) int {
	grid := make([][]string, 0)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		lineSplit := strings.Split(line, "")
		if len(lineSplit) == 0 {
			continue
		}
		grid = append(grid, lineSplit)
	}

	uniquePoints := make(map[string][]point, 0)

	for i := range grid {
		for k := range grid[i] {
			if sym := grid[i][k]; sym != "." {
				_, ok := uniquePoints[sym]
				if !ok {
					uniquePoints[sym] = make([]point, 0)
				}
				uniquePoints[sym] = append(uniquePoints[sym], point{i, k})
			}
		}
	}

	antinodes := antinodes{nodes: make([]point, 0), limitX: len(grid), limitY: len(grid[0])}

	for _, v := range uniquePoints {
		an := findAntinodes(v)
		for _, n := range an {
			antinodes.append(n)
		}
	}

	return antinodes.size
}

func findAntinodes(nodes []point) []point {
	res := make([]point, 0)
	for _, node1 := range nodes {
		for _, node2 := range nodes {
			if node1.x == node2.x && node1.y == node2.y {
				continue
			}
			if node1.x > node2.x {
				diffX := node1.x - node2.x
				diffY := node1.y - node2.y
				newX := node1.x + diffX
				newY := node1.y + diffY

				newX2 := node2.x - diffX
				newY2 := node2.y - diffY
				res = append(res, point{newX, newY})
				res = append(res, point{newX2, newY2})
			} else if node1.x <= node2.x {
				diffX := node2.x - node1.x
				diffY := node2.y - node1.y
				newX := node2.x + diffX
				newY := node2.y + diffY

				newX2 := node1.x - diffX
				newY2 := node1.y - diffY
				res = append(res, point{newX2, newY2})
				res = append(res, point{newX, newY})
			}

		}
	}
	return res
}

func handlePart2(input string) int {
	grid := make([][]string, 0)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		lineSplit := strings.Split(line, "")
		if len(lineSplit) == 0 {
			continue
		}
		grid = append(grid, lineSplit)
	}

	uniquePoints := make(map[string][]point, 0)

	for i := range grid {
		for k := range grid[i] {
			if sym := grid[i][k]; sym != "." {
				_, ok := uniquePoints[sym]
				if !ok {
					uniquePoints[sym] = make([]point, 0)
				}
				uniquePoints[sym] = append(uniquePoints[sym], point{i, k})
			}
		}
	}

	antinodes := antinodes{nodes: make([]point, 0), limitX: len(grid), limitY: len(grid[0])}

	for _, v := range uniquePoints {
		an := findAntinodes2(v, len(grid), len(grid[1]))
		for _, n := range an {
			antinodes.append(n)
		}
	}

	for _, v := range uniquePoints {
		for _, vv := range v {
			antinodes.append(vv)
		}
	}
	return antinodes.size
}

func findAntinodes2(nodes []point, limitX, limitY int) []point {
	res := make([]point, 0)
	for _, node1 := range nodes {
		for _, node2 := range nodes {
			if node1.x == node2.x && node1.y == node2.y {
				continue
			}

			if node1.x > node2.x {
				diffX := node1.x - node2.x
				diffY := node1.y - node2.y
				newX := node1.x + diffX
				newY := node1.y + diffY
				newX2 := node2.x - diffX
				newY2 := node2.y - diffY
				res = append(res, point{newX, newY})
				res = append(res, point{newX2, newY2})
				for {
					newX = newX + diffX
					newY = newY + diffY
					newX2 = newX2 - diffX
					newY2 = newY2 - diffY
					var stop1, stop2 bool

					if newX >= limitX || newX < 0 || newY >= limitY || newY < 0 {
						stop1 = true
					}
					if !stop1 {
						res = append(res, point{newX, newY})
					}

					if newX2 >= limitX || newX2 < 0 || newY2 >= limitY || newY2 < 0 {
						stop2 = true
					}
					if !stop2 {
						res = append(res, point{newX2, newY2})
					}
					if stop1 && stop2 {
						break
					}
				}
			} else if node1.x <= node2.x {
				diffX := node2.x - node1.x
				diffY := node2.y - node1.y
				newX := node2.x + diffX
				newY := node2.y + diffY
				newX2 := node1.x - diffX
				newY2 := node1.y - diffY
				res = append(res, point{newX, newY})
				res = append(res, point{newX2, newY2})
				for {
					var stop1, stop2 bool
					newX = newX + diffX
					newY = newY + diffY
					newX2 = newX2 - diffX
					newY2 = newY2 - diffY
					if newX >= limitX || newX < 0 || newY >= limitY || newY < 0 {
						stop1 = true
					}
					if !stop1 {
						res = append(res, point{newX, newY})
					}

					if newX2 >= limitX || newX2 < 0 || newY2 >= limitY || newY2 < 0 {
						stop2 = true
					}
					if !stop2 {
						res = append(res, point{newX2, newY2})
					}
					if stop1 && stop2 {
						break
					}
				}
			}

		}
	}
	return res
}
