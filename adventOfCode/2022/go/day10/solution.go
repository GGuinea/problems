package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Command string
	Value   int
}

type Register struct {
	Value                          int
	CurrentCycle                   int
	ImportantCyclesCalculatedValue int
	CrtPointer                     int
}

var IMPORTANT_CYCLES = []int{20, 60, 100, 140, 180, 220}

func (r *Register) execute(command Command) {
	switch command.Command {
	case "noop":
		r.CurrentCycle++
		r.drawCurrentState()
		break
	case "addx":
		for i := 2; i > 0; i-- {
			r.CurrentCycle++
			for _, important := range IMPORTANT_CYCLES {
				if important == r.CurrentCycle {
					r.ImportantCyclesCalculatedValue += important * r.Value
				}
			}
			r.drawCurrentState()
		}
		r.Value += command.Value
		break
	}
}

func (r *Register) drawCurrentState() {
	if r.CrtPointer == 40 {
		fmt.Print("\n")
		r.CrtPointer = 0
	}
	if r.CrtPointer == r.Value-1 || r.CrtPointer == r.Value || r.CrtPointer == r.Value+1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	r.CrtPointer++
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
	commands := parseCommands(lines)
	register := &Register{Value: 1, CurrentCycle: 0}
	for _, command := range commands {
		register.execute(command)
	}
	fmt.Println(register.ImportantCyclesCalculatedValue)
}

func part2(fileName string) {
	lines := readFile(fileName)
	commands := parseCommands(lines)
	register := &Register{Value: 1, CurrentCycle: 0, CrtPointer: 0}
	for _, command := range commands {
		register.execute(command)
	}
}

func parseCommands(lines []string) []Command {
	commandList := make([]Command, 0)
	for _, line := range lines {
		splitted := strings.Split(line, " ")
		if splitted[0] == "noop" {
			commandList = append(commandList, Command{Command: splitted[0]})
			continue
		}

		value, err := strconv.Atoi(splitted[1])
		if err != nil {
			panic(2)
		}
		commandList = append(commandList, Command{Command: splitted[0], Value: value})
	}
	return commandList
}

func main() {
	fileName := os.Args[1]
	//part1(fileName)
	part2(fileName)
}
