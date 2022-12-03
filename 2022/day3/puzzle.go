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

	// res, err = partOne(lines)
	res, err = partTwo(lines)

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
		var count map[rune]int = make(map[rune]int, l) // TODO should have used strings.ContainsRune to check!
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

func partTwo(lines []string) (string, error) {
	total := 0
	for group := 0; group < len(lines); group += 3 {
		l := len(lines[group]) + len(lines[group+1]) + len(lines[group+2])
		candidates := make(map[rune]bool, l)
		for _, c := range lines[group] {
			candidates[c] = false
		}
		for _, c := range lines[group+1] {
			if _, ok := candidates[c]; ok {
				candidates[c] = true
			}
		}
		for _, c := range lines[group+2] {
			if candidates[c] {
				total += (strings.Index(alphabet, string(c)) + 1)
				break
			}
		}

	}
	return strconv.Itoa(total), nil
}
