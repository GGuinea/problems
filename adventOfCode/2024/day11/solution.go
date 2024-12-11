package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
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
	return handlePart1(input, 25)
}

func part2() int {
	input := readFile("input")
	return handlePart1(input, 75)
}

type cache struct {
	mu   sync.RWMutex
	data map[int]map[int]int
}

func (c *cache) get(val int, blink int) (int, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	res, found := c.data[val][blink]
	return res, found
}

func (c *cache) add(x int, blink int, val int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, found := c.data[x]
	if !found {
		c.data[x] = make(map[int]int)
	}

	c.data[x][blink] = val
	return
}

func handlePart1(input string, blinks int) int {
	input = strings.TrimSpace(input)
	strs := strings.Split(input, " ")
	nums := make([]int, 0, len(strs))
	for _, str := range strs {
		if len(str) == 0 {
			continue
		}

		val, err := strconv.Atoi(str)
		check(err)
		nums = append(nums, val)
	}

	cacheMap := cache{mu: sync.RWMutex{}, data: make(map[int]map[int]int)}
	res := 0
	for _, v := range nums {
		res += calculate(v, blinks, &cacheMap)
	}

	return res
}

func calculate(v int, blinks int, c *cache) int {
	fromCache, found := c.get(v, blinks)
	if found {
		return fromCache
	}
	if blinks == 0 {
		return 1
	}

	str := fmt.Sprint(v)
	var res int
	switch {
	case v == 0:
		res = calculate(1, blinks-1, c)
	case len(str)%2 == 0:
		first := str[:len(str)/2]
		second := str[len(str)/2:]
		firstVal, _ := strconv.Atoi(first)
		secondVal, _ := strconv.Atoi(second)
		res = calculate(firstVal, blinks-1, c) + calculate(secondVal, blinks-1, c)
	default:
		res = calculate(v*2024, blinks-1, c)
	}
	c.add(v, blinks, res)
	return res
}
