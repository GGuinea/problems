package day1

import (
	"os"
	"slices"
	"strconv"
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
	return handlePart1(input)
}

func part2() int {
	input := readFile("input")
	return handlePart2(input)
}

func handlePart1(input string) int {
	disk := make([]int, 0, len(input))
	input = strings.TrimSpace(input)
	for i := range len(input) {
		parsed, err := strconv.Atoi(string(input[i]))
		if err != nil {
			panic(err)
		}
		disk = append(disk, parsed)
	}

	resDisk := make([]int, 0, 1000)
	fileId := 0

	for i, v := range disk {
		isFile := false
		for range v {
			if i%2 == 1 {
				resDisk = append(resDisk, -1)
			} else {
				resDisk = append(resDisk, fileId)
				isFile = true
			}
		}
		if isFile {
			fileId++
		}
	}

	lastFilePointer := len(resDisk) - 1
	firstFreeSpace := 0
	for {
		if firstFreeSpace > lastFilePointer {
			break
		}
		if firstFreeSpace == len(resDisk) {
			break
		}

		if lastFilePointer <= 0 {
			continue
		}

		if resDisk[lastFilePointer] == -1 {
			lastFilePointer--
			continue
		}

		if resDisk[firstFreeSpace] != -1 {
			firstFreeSpace++
			continue
		}

		resDisk[firstFreeSpace] = resDisk[lastFilePointer]
		resDisk[lastFilePointer] = -1
		firstFreeSpace++
		lastFilePointer--
	}

	res := 0

	for i, v := range resDisk {
		if v == -1 {
			break
		}
		res += i * v
	}
	return res
}

func handlePart2(input string) int {
	disk := make([]int, 0, len(input))
	input = strings.TrimSpace(input)
	for i := range len(input) {
		parsed, err := strconv.Atoi(string(input[i]))
		if err != nil {
			panic(err)
		}
		disk = append(disk, parsed)
	}

	resDisk := make([]int, 0, 100000)
	fileId := 0

	freeSpace := make(map[int]int, 10000)
	files := make(map[int]int, 10000)
	filesToPosition := make(map[int]int, 10000)
	resDiskPointer := 0

	keys := []int{}
	maxFileId := 0

	for i, v := range disk {
		if fileId > maxFileId {
			maxFileId = fileId
		}
		isFile := false
		for range v {
			if i%2 == 1 {
				resDisk = append(resDisk, -1)
			} else {
				resDisk = append(resDisk, fileId)
				isFile = true
			}
		}
		if isFile {
			filesToPosition[fileId] = resDiskPointer
			files[resDiskPointer] = v
			fileId++
		} else {
			keys = append(keys, resDiskPointer)
			freeSpace[resDiskPointer] = v
		}
		resDiskPointer += v
	}

	for {
		if maxFileId <= 0 {
			break
		}

		filePosition := filesToPosition[maxFileId]
		fileSize := files[filePosition]

		var found bool
		newFilePosition := 0
		var foundSpace int
		for _, v := range keys {
			if v > filePosition {
				break
			}
			k := freeSpace[v]
			if k >= fileSize {
				newFilePosition = v
				foundSpace = k
				found = true
				break
			}
		}

		if !found {
			maxFileId--
			continue
		}

		for i := range fileSize {
			resDisk[newFilePosition+i] = maxFileId
			resDisk[filePosition+i] = -1
		}

		if foundSpace != fileSize {
			newFree := foundSpace - fileSize
			freeSpace[newFilePosition+fileSize] = newFree
			keys = append(keys, newFilePosition+fileSize)
			slices.Sort(keys)
		}

		freeSpace[newFilePosition] = -1
	}

	res := 0
	for i, v := range resDisk {
		if v == -1 {
			continue
		}
		res += i * v
	}
	return res
}
