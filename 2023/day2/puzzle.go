package day1

import (
	"advent-of-code-go/util"
)

const day = "day2"

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


func partOne(lines []string) (int, error) {
	sum := 0
	for _, line := range lines {
		
	}
	return sum, nil
}

