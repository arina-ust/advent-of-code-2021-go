package day16

import (
	"advent-of-code-2021/util"
	"fmt"
	"strings"
)

const name = "day16"

var inputFile string

func Solve(easy bool) (name string, res string, err error) {
	setInput(easy)
	lines, err := util.ReadStringList(inputFile)

	if easy {
		res = partOne(lines)
	} else {
		return name, "", fmt.Errorf("not solved yet")
	}

	return
}

func setInput(easy bool) {
	if easy {
		inputFile = name + "/" + util.InputFileEasy
	} else {
		inputFile = name + "/" + util.InputFileFull
	}
}

func partOne(lines []string) string {
	return strings.Join(lines, "\n") // TODO: solve
}
