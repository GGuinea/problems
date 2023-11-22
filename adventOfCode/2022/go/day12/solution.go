package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
)

const S int = 83
const E int = 69
const VISITED_MARKER int = -1

type Maze struct {
	Labyrinth       [][]int
	Found           bool
	ShorthestLength int
	Length          int
	Visited         [][]int
	Queue           []Point
	Start           Point
	End             Point
}

type Point struct {
	X    int
	Y    int
	dist int
}

func enqueue(queue []Point, element Point) []Point {
	return append(queue, element)
}

func dequeue(queue []Point) (Point, []Point) {
	element := queue[0]
	if len(queue) == 1 {
		var tmp = []Point{}
		return element, tmp
	}
	return element, queue[1:]
}

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

func part1(fileName string) {
	lines := readFile(fileName)
	labyrinth, x, y, x2, y2 := parseInput(lines)

	queue := []Point{}
	queue = enqueue(queue, Point{X: x, Y: y, dist: 0})
	visited := make([][]int, len(labyrinth))

	end := Point{X: x2, Y: y2}
	start := Point{X: x, Y: y, dist: 0}

	for i := range visited {
		visited[i] = make([]int, len(labyrinth[0]))

	}

	maze := Maze{
		Labyrinth:       labyrinth,
		Found:           false,
		ShorthestLength: math.MaxInt,
		Visited:         visited,
		Queue:           queue,
		Start:           start,
		End:             end,
	}
	dist := traverseMaze(maze)
	fmt.Println(dist)
}

func traverseMaze(maze Maze) int {
	for len(maze.Queue) > 0 {
		point, queue := dequeue(maze.Queue)
		maze.Queue = queue

		if point.X == maze.End.X && point.Y == maze.End.Y {
			return point.dist
		}

		if maze.Visited[point.X][point.Y] == VISITED_MARKER {
			continue
		}

		maze.Visited[point.X][point.Y] = VISITED_MARKER

		if canVisit(&maze, point.X+1, point.Y, maze.Labyrinth[point.X][point.Y]) {
			maze.Queue = enqueue(maze.Queue, Point{X: point.X + 1, Y: point.Y, dist: point.dist + 1})
		}

		if canVisit(&maze, point.X, point.Y+1, maze.Labyrinth[point.X][point.Y]) {
			maze.Queue = enqueue(maze.Queue, Point{X: point.X, Y: point.Y + 1, dist: point.dist + 1})
		}

		if canVisit(&maze, point.X-1, point.Y, maze.Labyrinth[point.X][point.Y]) {
			maze.Queue = enqueue(maze.Queue, Point{X: point.X - 1, Y: point.Y, dist: point.dist + 1})
		}

		if canVisit(&maze, point.X, point.Y-1, maze.Labyrinth[point.X][point.Y]) {
			maze.Queue = enqueue(maze.Queue, Point{X: point.X, Y: point.Y - 1, dist: point.dist + 1})
		}
	}
	return -1
}

func part2(fileName string) {
	lines := readFile(fileName)
	labyrinth, startingPoints, end := parseInputPart2(lines)

	outputs := make([]int, 0)

	visited := make([][]int, len(labyrinth))
	for _, startingPoint := range startingPoints {

		for i := range visited {
			visited[i] = make([]int, len(labyrinth[0]))

		}
		queue := []Point{startingPoint}
		maze := Maze{
			Labyrinth:       labyrinth,
			Found:           false,
			ShorthestLength: math.MaxInt,
			Visited:         visited,
			Queue:           queue,
			End:             end,
		}
		dist := traverseMaze(maze)
		outputs = append(outputs, dist)
	}
	slices.Sort(outputs)
	for _, output := range outputs {
		if output != -1 {
			fmt.Println(output)
			break
		}
	}
}

func canVisit(maze *Maze, x, y int, currentValue int) bool {
	if x < 0 || y < 0 || x >= len(maze.Labyrinth) || y >= len(maze.Labyrinth[0]) {
		return false
	}

	if maze.Visited[x][y] == VISITED_MARKER {
		return false
	}

	if (maze.Labyrinth[x][y] - currentValue) > 1 {
		return false
	}

	return true
}

func getAbsDifference(a, b int) int {
	return int(math.Abs(float64(a) - float64(b)))
}

func parseInput(lines []string) ([][]int, int, int, int, int) {
	labyrinth := make([][]int, 0)
	startingI := 0
	startingJ := 0

	endI := 0
	endJ := 0

	for i, line := range lines {
		arr := make([]int, 0)
		for j, lett := range line {
			if int(lett) == S {
				startingI = i
				startingJ = j
				arr = append(arr, 97)
			} else if int(lett) == E {
				endI = i
				endJ = j
				arr = append(arr, 122)
			} else {
				arr = append(arr, int(lett))
			}
		}
		labyrinth = append(labyrinth, arr)
	}

	return labyrinth, startingI, startingJ, endI, endJ
}

func parseInputPart2(lines []string) ([][]int, []Point, Point) {
	labyrinth := make([][]int, 0)

	endI := 0
	endJ := 0

	startingPoints := make([]Point, 0)

	for i, line := range lines {
		arr := make([]int, 0)
		for j, lett := range line {
			if int(lett) == 97 && ((i == 0 || i == len(lines)) || (j == 0 || j == len(line))) {
				startingPoints = append(startingPoints, Point{X: i, Y: j})
				arr = append(arr, 97)
			} else if int(lett) == E {
				endI = i
				endJ = j
				arr = append(arr, 122)
			} else {
				arr = append(arr, int(lett))
			}
		}
		labyrinth = append(labyrinth, arr)
	}

	return labyrinth, startingPoints, Point{X: endI, Y: endJ}
}

func main() {
	fileName := os.Args[1]
	part2(fileName)
}
