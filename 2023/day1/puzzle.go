package day1

import (
	"advent-of-code-go/util"
	"fmt"
	"regexp"
	"strconv"
)

const day = "day1"

var inputFile string

func Solve(easy bool) (name string, res int, err error) {
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

var validDigit = regexp.MustCompile("\\d")

func partOne(lines []string) (int, error) {
	sum := 0
	for _, line := range lines {
		digits := validDigit.FindAllString(line, -1)
		var result string
		if len(digits) == 0 {
			continue
		}
		if len(digits) == 1 {
			result = digits[0] + digits[0]
		} else {
			last := len(digits) - 1
			result = digits[0] + digits[last]
		}
		value, err := strconv.Atoi(result)
		if err != nil {
			return -1, fmt.Errorf("failed to convert to int %s", result)
		}
		sum += value
	}
	return sum, nil
}