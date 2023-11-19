package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Direction string
	Count     int
}

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

type PointContainer struct {
	Container []Point
}

func NewPointContainer() *PointContainer {
	container := make([]Point, 0)
	return &PointContainer{Container: container}
}

func (pc *PointContainer) Save(newPoint Point) {
	for _, savedPoint := range pc.Container {
		if savedPoint.X == newPoint.X && savedPoint.Y == newPoint.Y {
			return
		}
	}
	pc.Container = append(pc.Container, newPoint)
}

type Snake struct {
	SnakeParts []*Point
	Length     int
}

func newSnake(length int) *Snake {
	var snake []*Point
	for i := 0; i < length; i++ {
		snake = append(snake, &Point{X: 0, Y: 0})
	}
	return &Snake{SnakeParts: snake, Length: length}
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

func part1() {
	lines := readFile("input")
	commands := parseCommands(lines)
	poinContainer := NewPointContainer()
	findSafePath(commands, poinContainer)
	fmt.Println(len(poinContainer.Container))
}

func findSafePath(commands []Command, pc *PointContainer) {
	head := &Point{X: 0, Y: 0}
	tail := &Point{X: 0, Y: 0}
	for _, command := range commands {
		makeMove(head, tail, command, pc)
	}
}

func makeMove(head, tail *Point, command Command, pc *PointContainer) {
	for i := 0; i < command.Count; i++ {
		fmt.Println(head, tail)
		moveHead(head, command.Direction)
		moveTail(head, tail)
		pc.Save(*tail)
	}
}

func part2() {
	lines := readFile("input")
	commands := parseCommands(lines)
	snake := newSnake(10)
	poinContainer := NewPointContainer()
	findSnakeSafePath(commands, poinContainer, snake)
	fmt.Println(len(poinContainer.Container))
}

func findSnakeSafePath(commands []Command, pc *PointContainer, snake *Snake) {
	for _, command := range commands {
		makeSnakeMove(snake.SnakeParts[0], command, pc, snake)
	}
}

func makeSnakeMove(head *Point, command Command, pc *PointContainer, snake *Snake) {
	for i := 0; i < command.Count; i++ {
		moveHead(head, command.Direction)
		for i := 0; i < snake.Length-1; i++ {
			moveTail(snake.SnakeParts[i], snake.SnakeParts[i+1])
		}
		pc.Save(*snake.SnakeParts[snake.Length-1])
	}
}

func moveTail(head, tail *Point) {
	xDistance := math.Abs(float64(calculateDistance(head.X, tail.X)))
	yDistance := math.Abs(float64(calculateDistance(head.Y, tail.Y)))
	if xDistance <= 1 && yDistance <= 1 {
		return
	} else if xDistance >= 2 && yDistance >= 2 {
		if head.X > tail.X {
			tail.X = head.X - 1
		} else if head.X < tail.X {
			tail.X = head.X + 1
		}
		if head.Y > tail.Y {
			tail.Y = head.Y - 1
		} else if head.Y < tail.Y {
			tail.Y = head.Y + 1
		}
	} else if xDistance > 1 {
		if head.X > tail.X {
			tail.X = head.X - 1
			tail.Y = head.Y
		} else if head.X < tail.X {
			tail.X = head.X + 1
			tail.Y = head.Y
		}
	} else if yDistance > 1 {
		if head.Y > tail.Y {
			tail.X = head.X
			tail.Y = head.Y - 1
		} else if head.Y < tail.Y {
			tail.X = head.X
			tail.Y = head.Y + 1
		}
	}
}

func calculateDistance(x1, x2 int) int {
	return x1 - x2
}

func moveHead(head *Point, direction string) {
	switch direction {
	case "R":
		head.X++
		break
	case "L":
		head.X--
		break
	case "U":
		head.Y++
		break
	case "D":
		head.Y--
		break
	default:
		panic(3)
	}
}

func parseCommands(lines []string) []Command {
	commands := make([]Command, 0)
	for _, line := range lines {
		splitted := strings.Split(line, " ")
		countConverted, err := strconv.Atoi(splitted[1])
		if err != nil {
			panic(2)
		}
		commands = append(commands, Command{
			Direction: splitted[0],
			Count:     countConverted,
		})
	}
	return commands
}

func main() {
	//part1()
	part2()
}
