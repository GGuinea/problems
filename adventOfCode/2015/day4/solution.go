package solution

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

func part1(hash string) int {
	return handlePart1(hash)
}

func handlePart1(input string) int {
	i := 1
	for {
		hsh := md5.New()
		io.WriteString(hsh, fmt.Sprintf("%v%v", input, i))
		str := fmt.Sprintf("%x", hsh.Sum(nil))
		if strings.HasPrefix(str, "00000") {
			fmt.Println(str)
			return i
		}
		i++
	}
}

func part2(hash string) int {
	return handlePart2(hash)
}

func handlePart2(input string) int {
	i := 1
	for {
		hsh := md5.New()
		io.WriteString(hsh, fmt.Sprintf("%v%v", input, i))
		str := fmt.Sprintf("%x", hsh.Sum(nil))
		if strings.HasPrefix(str, "000000") {
			fmt.Println(str)
			return i
		}
		i++
	}
}
