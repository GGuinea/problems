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

type File struct {
	Name string
	Size int
}

type Directory struct {
	Name   string
	Parent *Directory
	Child  []*Directory
	Files  []*File
	Size   int
}

func newRoot() *Directory {
	directoryList := make([]*Directory, 0)
	filesList := make([]*File, 0)
	return &Directory{Name: "/", Parent: nil, Child: directoryList, Files: filesList}
}

func part1() {
	lines := readFile("input")
	tree := buildTree(lines)
	calculateDirSizesWithDfs(tree)
	res := findDirsWithSizeIsLessThan(tree, 100000)
	fmt.Println(res)
}

func buildTree(lines []string) *Directory {
	rootDirectory := newRoot()
	currentDir := rootDirectory
	for _, line := range lines {
		if strings.Contains(line, "$ cd") {
			if strings.Contains(line, "$ cd ..") {
				currentDir = currentDir.Parent
			} else if !strings.Contains(line, "cd /") {
				splitted := strings.Split(line, " ")
				newDirectory := Directory{
					Name:   splitted[2],
					Parent: currentDir,
					Child:  make([]*Directory, 0),
					Files:  make([]*File, 0),
					Size:   0,
				}
				currentDir.Child = append(currentDir.Child, &newDirectory)
				currentDir = &newDirectory
			}
		} else if strings.Contains(line, "$ ls") || strings.HasPrefix(line, "dir") {
			continue
		} else {
			splittedFile := strings.Split(line, " ")
			size, err := strconv.Atoi(splittedFile[0])
			if err != nil {
				panic(2)
			}
			file := File{
				Name: splittedFile[1],
				Size: size,
			}
			currentDir.Files = append(currentDir.Files, &file)
		}
	}
	return rootDirectory
}

func findDirsWithSizeIsLessThan(root *Directory, threshold int) int {
	counter := 0
	if (root.Size) < threshold {
		counter += root.Size
	}
	for _, dir := range root.Child {
		counter += findDirsWithSizeIsLessThan(dir, threshold)
	}
	return counter
}

func calculateDirSizesWithDfs(root *Directory) int {
	files := root.Files
	fileSize := 0
	for _, entry := range files {
		fileSize += entry.Size
	}
	dirs := root.Child

	dirsSize := 0
	for _, entry := range dirs {
		dirsSize += calculateDirSizesWithDfs(entry)
	}
	root.Size = fileSize + dirsSize
	return root.Size
}

func main() {
	part1()
}
