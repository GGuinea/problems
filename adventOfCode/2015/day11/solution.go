package solution

import (
	"os"
	"strings"
)

func readFile(filename string) string {
	res, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(res)
}

func part1(input string) string {
	return handlePart1(input)
}

func handlePart1(input string) string {
	ptr := len(input) - 1
	inputRune := []rune(input)
	for {
		res, over := incrementCharacter(byte(inputRune[ptr]))
		inputRune[ptr] = rune(res)

		if over {
			inputRune = updateRest(inputRune, ptr-1)
		}

		if validateRules(string(inputRune), []func(string) bool{atLeastThreeLettersIncreasedInRow, containsTwoDifferentPairs, notContainsForbidden}) {
			return string(inputRune)
		}
	}
}

func updateRest(input []rune, ptr int) []rune {
	if ptr < 0 {
		return input
	}

	res, over := incrementCharacter(byte(input[ptr]))
	input[ptr] = rune(res)
	if over {
		return updateRest(input, ptr-1)
	}
	return input
}

func atLeastThreeLettersIncreasedInRow(input string) bool {
	for i := 0; i <= len(input)-3; i++ {
		if input[i]+1 == input[i+1] && input[i+1]+1 == input[i+2] {
			return true
		}
	}
	return false
}

func containsTwoDifferentPairs(input string) bool {
	pairs := 0
	var firstPair byte
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			if pairs == 0 {
				pairs++
				i++
				firstPair = input[i]
			} else if firstPair != input[i] {
				return true
			}
		}
	}
	return false
}

func notContainsForbidden(input string) bool {
	return !strings.ContainsAny(input, "iol")
}

func validateRules(input string, fns []func(string) bool) bool {
	for _, fn := range fns {
		res := fn(input)
		if !res {
			return false
		}
	}

	return true
}

func incrementCharacter(c byte) (byte, bool) {
	if c == 'z' {
		return 'a', true
	}
	return c + 1, false
}
