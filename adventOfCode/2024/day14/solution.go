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
	return handlePart1(input, 100, 103, 101)
}

func part2() int {
	input := readFile("input")
	return handlePart2(input, 100, 103, 101)
}

type robot struct {
	x, y   int
	vX, vY int
}

func (r *robot) move(maxX, maxY int) {
	newX := r.x + r.vX
	newY := r.y + r.vY

	if newX < 0 {
		newX = maxX + newX
	}

	if newY < 0 {
		newY = maxY + newY
	}

	if newX >= maxX {
		newX = newX % maxX
	}

	if newY >= maxY {
		newY = newY % maxY
	}

	r.x = newX
	r.y = newY
}

func handlePart1(input string, seconds int, maxTilesX, maxTilesY int) int {
	lines := strings.Split(input, "\n")
	robots := make([]*robot, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		var x, y, vX, vY int
		n, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &y, &x, &vY, &vX)
		check(err)
		if n != 4 {
			panic("wrong nr of input: " + line)
		}

		robots = append(robots, &robot{x, y, vX, vY})
	}

	for i := 0; i < seconds; i++ {
		for _, r := range robots {
			r.move(maxTilesX, maxTilesY)
		}
	}

	middleY := maxTilesY / 2
	middleX := maxTilesX / 2

	quadrants := make(map[int]int, 0)
	for _, r := range robots {
		//fmt.Println(r.x, r.y)
		if r.x == middleX || r.y == middleY {
			fmt.Println("skipping: ", r.x, r.y)
			continue
		}

		if r.x < middleX && r.y < middleY {
			quadrants[0]++
		}
		if r.x < middleX && r.y > middleY {
			quadrants[1]++
		}

		if r.x > middleX && r.y < middleY {
			quadrants[2]++
		}

		if r.x > middleX && r.y > middleY {
			quadrants[3]++
		}
	}

	res := 1
	for _, v := range quadrants {
		//fmt.Println(k, v)
		res *= v
	}
	return res
}

func handlePart2(input string, seconds int, maxTilesX, maxTilesY int) int {
	lines := strings.Split(input, "\n")
	robots := make([]*robot, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		var x, y, vX, vY int
		n, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &y, &x, &vY, &vX)
		check(err)
		if n != 4 {
			panic("wrong nr of input: " + line)
		}

		robots = append(robots, &robot{x, y, vX, vY})
	}
	grid := make([][]string, 0, maxTilesY)
	for k := 0; k < maxTilesX; k++ {
		grid = append(grid, make([]string, maxTilesX))
	}
	for i := 0; i < 6700; i++ {
		for k := range grid {
			for z := range grid[k] {
				grid[z][k] = ""
			}
		}
		for _, r := range robots {
			r.move(maxTilesX, maxTilesY)
			grid[r.x][r.y] = "#"
		}
		if i > 6660 && i < 6670 {
		fmt.Println("--------------------------------", i+1, "------------------")
		fmt.Println("--------------------------------", i+1, "------------------")
		fmt.Println("--------------------------------", i+1, "------------------")
			for k := range grid {
				for _, z := range grid[k] {
					if z != "#" {
						fmt.Print(".")
					} else {
						fmt.Print("#")
					}
				}
				fmt.Println()
			}

		}
	}

	return 1
}
