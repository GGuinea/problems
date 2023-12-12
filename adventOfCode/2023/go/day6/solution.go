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

type Race struct {
	Time     int
	Distacne int
}

func getNumbers(string string) []int {
	numbers := make([]int, 0)

	split := strings.Split(string, " ")
	for _, sign := range split {
		num, err := strconv.Atoi(sign)
		if err != nil {
			continue
		}
		numbers = append(numbers, num)
	}

	return numbers
}

func getRaces(lines []string) []Race {
	races := make([]Race, 0)
	times := getNumbers(lines[0])
	dists := getNumbers(lines[1])

	for i := range times {
		races = append(races, Race{
			Time:     times[i],
			Distacne: dists[i],
		})
	}

	return races
}

func calculatePossibleWinsNo(race Race) int {
	res := 0

	for i := 0; i < race.Time; i++ {
		timeRemaining := race.Time - i
		neededSpeed := race.Distacne / timeRemaining
		if i > neededSpeed {
			res++
		}
	}

	return res
}

func part1(filename string) int {
	lines := readFile(filename)
	races := getRaces(lines)
	resp := 1

	for _, race := range races {
		resp *= calculatePossibleWinsNo(race)
	}

	return resp
}

func part2(filename string) int {
	lines := readFile(filename)
	races := getRaces(lines)
	time := ""
	dist := ""

	for _, race := range races {
		time += strconv.Itoa(race.Time)
		dist += strconv.Itoa(race.Distacne)
	}

	timeConv, _ := strconv.Atoi(time)
	distConv, _ := strconv.Atoi(dist)

	return calculatePossibleWinsNo(Race{Time: timeConv, Distacne: distConv})
}

func main() {
	fileName := os.Args[1]
	//fmt.Println(part1(fileName))
	fmt.Println(part2(fileName))
}
