package day3

import (
	"advent-of-code-go/util"
	"strconv"
	"strings"
)

const day = "day3"

var inputFile string

func Solve(easy bool) (name string, res string, err error) {
	name = day
	setInput(easy)
	lines, err := util.ReadStringList(inputFile)
	if err != nil {
		return
	}

	res, err = partOne(lines)
	// res, err = partTwo(lines)

	return
}

func setInput(easy bool) {
	if easy {
		inputFile = day + "/" + util.InputFileEasy
	} else {
		inputFile = day + "/" + util.InputFileFull
	}
}

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func partOne(lines []string) (string, error) {
	total := 0
	for _, line := range lines {
		l := len(line)
		secondCompartmentIndex := l / 2
		var count map[rune]int = make(map[rune]int, l)
		for _, c := range line[:secondCompartmentIndex] {
			count[c] += 1
		}
		for _, c := range line[secondCompartmentIndex:] {
			if count[c] != 0 {
				total += (strings.Index(alphabet, string(c)) + 1)
				break
			}
		}

	}
	return strconv.Itoa(total), nil
}
