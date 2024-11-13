package solution

import (
	"os"
	"strings"
	"sync"
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

func handlePart1(input string) int {
	lines := strings.Split(input, "\n")
	res := 0
	for _, line := range lines {
		if isValid(line) {
			res++
		}
	}
	return res
}

func isValid(input string) bool {
	var wg sync.WaitGroup
	wg.Add(3)
	var (
		vowelMatched         bool
		twiceInRowMatched    bool
		forbbidenRuleMatched bool
	)

	go func() {
		defer wg.Done()
		vowelMatched = isVowelRuleMatched(input)
	}()
	go func() {
		defer wg.Done()
		twiceInRowMatched = isTwiceInRowRuleMatched(input)
	}()
	go func() {
		defer wg.Done()
		forbbidenRuleMatched = isForbiddenStringsRuleMatched(input)
	}()

	wg.Wait()

	return vowelMatched && twiceInRowMatched && forbbidenRuleMatched

}

func isVowelRuleMatched(input string) bool {
	vovels := []string{"a", "e", "i", "o", "u"}
	countOfVowels := 0
	for _, vovel := range vovels {
		countOfVowels += strings.Count(input, vovel)
	}
	return countOfVowels >= 3
}

func isTwiceInRowRuleMatched(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			return true
		}
	}
	return false
}

func isForbiddenStringsRuleMatched(input string) bool {
	forbiddens := []string{"ab", "cd", "pq", "xy"}
	for _, forbidden := range forbiddens {
		if strings.Contains(input, forbidden) {
			return false
		}
	}
	return true
}

func part2() int {
	input := readFile("input")
	return handlePart2(input)
}

func handlePart2(input string) int {
	lines := strings.Split(input, "\n")
	res := 0
	for _, line := range lines {
		if isValidPart2(line) {
			res++
		}
	}
	return res
}

func isValidPart2(input string) bool {
	var wg sync.WaitGroup
	wg.Add(2)
	var (
		pairAtLeastTwice  bool
		pairWithSeparator bool
	)

	go func() {
		defer wg.Done()
		pairAtLeastTwice = isPairAtLeastTwice(input)
	}()
	go func() {
		defer wg.Done()
		pairWithSeparator = isPairWithSeparatorRuleMatching(input)
	}()

	wg.Wait()
	//fmt.Printf("input: %v; pairAtLeastTwice: %v, pairWithSeparator: %v", input, pairAtLeastTwice, pairWithSeparator)
	return pairAtLeastTwice && pairWithSeparator

}

func isPairAtLeastTwice(input string) bool {
	for i := 0; i < len(input)-2; i++ {
		if strings.Count(input, input[i:i+2]) >= 2 {
			return true
		}
	}
	return false
}

func isPairWithSeparatorRuleMatching(input string) bool {
	for i := 0; i < len(input)-2; i++ {
		if input[i] == input[i+2] {
			return true
		}
	}
	return false

}
