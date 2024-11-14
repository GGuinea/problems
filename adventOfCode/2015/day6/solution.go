package solution

import (
	"os"
	"strconv"
	"strings"
)

const (
	SIZE_X = 1000
	SIZE_Y = 1000
)

func readFile(filename string) string {
	res, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(res)
}

func part1() int {
	input := readFile("input")
	return handlePart1(input)
}

func part2() int {
	input := readFile("input")
	return handlePart2(input)
}

func handlePart1(input string) int {
	lights := make([][]bool, SIZE_X)
	for i := range lights {
		lights[i] = make([]bool, SIZE_Y)
	}

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		ops := strings.Split(line, " ")
		if len(ops) < 3 {
			continue
		}
		var startX, startY, endX, endY int
		switch ops[0] {
		case "turn":
			startX, startY = getXY(ops[2])
			endX, endY = getXY(ops[4])
		case "toggle":
			startX, startY = getXY(ops[1])
			endX, endY = getXY(ops[3])
		}

		xDirection, yDirection := 1, 1
		if endX-startX < 0 {
			xDirection = -1
		}
		if endY-startY < 0 {
			yDirection = -1
		}

		currentX, currentY := startX, startY
		for {
			currentY = startY
			for {
				switch ops[1] {
				case "on":
					lights[currentX][currentY] = true
				case "off":
					lights[currentX][currentY] = false
				default: // toggle
					lights[currentX][currentY] = !lights[currentX][currentY]
				}
				if currentY == endY {
					break
				}
				currentY += 1 * yDirection
			}
			if currentX == endX {
				break
			}
			currentX += 1 * xDirection
		}
	}

	res := 0
	for i := range SIZE_X {
		for j := range SIZE_Y {
			if lights[i][j] {
				res++
			}
		}
	}
	return res
}

func handlePart2(input string) int {
	lights := make([][]int, SIZE_X)
	for i := range lights {
		lights[i] = make([]int, SIZE_Y)
	}

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		ops := strings.Split(line, " ")
		if len(ops) < 3 {
			continue
		}
		var startX, startY, endX, endY int
		switch ops[0] {
		case "turn":
			startX, startY = getXY(ops[2])
			endX, endY = getXY(ops[4])
		case "toggle":
			startX, startY = getXY(ops[1])
			endX, endY = getXY(ops[3])
		}

		xDirection, yDirection := 1, 1
		currentX, currentY := startX, startY

		if endX-startX < 0 {
			xDirection = -1
		}

		if endY-startY < 0 {
			yDirection = -1
		}
		for {
			currentY = startY
			for {
				switch ops[1] {
				case "on":
					lights[currentX][currentY]++
				case "off":
					if lights[currentX][currentY] > 0 {
						lights[currentX][currentY]--
					}
				default: // toggle
					lights[currentX][currentY] += 2
				}
				if currentY == endY {
					break
				}
				currentY += 1 * yDirection
			}
			if currentX == endX {
				break
			}
			currentX += 1 * xDirection
		}
	}

	res := 0
	for i := range SIZE_X {
		for j := range SIZE_Y {
			res += lights[i][j]
		}
	}
	return res
}

func getXY(input string) (int, int) {
	split := strings.Split(input, ",")
	x, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}
	return x, y
}
