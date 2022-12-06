package day6

import (
	"advent-of-code-go/util"
	"strconv"
)

const day = "day6"

var inputFile string

func Solve(easy bool) (name string, res string, err error) {
	name = day
	setInput(easy)
	lines, err := util.ReadStringList(inputFile)
	if err != nil {
		return
	}

	res, err = partOne(lines)

	return
}

func setInput(easy bool) {
	if easy {
		inputFile = day + "/" + util.InputFileEasy
	} else {
		inputFile = day + "/" + util.InputFileFull
	}
}

type marker string

func (m marker) hasDuplicate() bool {
	seen := make(map[rune]bool, len(m))
	for i := 0; i < len(m); i++ {
		ch := rune(m[i])
		if seen[ch] {
			return true
		} else {
			seen[ch] = true
		}
	}
	return false
}

func partOne(lines []string) (string, error) {
	line := lines[0]
	charCount := len(line)

	for i, j := 0, 4; i < charCount-3 && j < charCount+1; {
		if marker(line[i:j]).hasDuplicate() {
			i++
			j++
		} else {
			return strconv.Itoa(j), nil
		}
	}

	return strconv.Itoa(0), nil
}
