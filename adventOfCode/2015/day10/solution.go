package solution

import (
	"fmt"
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

func part1() int {
	return len(handlePart1("1321131112", 40))
}

func part2() int {
	return len(handlePart1("1321131112", 50))
}

func handlePart1(input string, iterations int) string {
	res := input
	for i := 0; i < iterations; i++ {
		var sb strings.Builder
		idx := 0

		for {
			if idx >= len(res) {
				break
			}

			currentChar := string(res[idx])
			occr := countOccurrences(res, currentChar, idx)
			idx += occr
			sb.WriteString(fmt.Sprintf("%v%v", occr, currentChar))
		}
		res = sb.String()
		//fmt.Printf("iter: %v, newRes: %v\n", i, newRes)
	}
	return res
}

func countOccurrences(str string, c string, start int) int {
	res := 0
	for i := start; i < len(str); i++ {
		if string(str[i]) == c {
			res++
		} else {
			break
		}
	}
	return res
}
