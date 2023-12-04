package day4

import (
	"advent-of-code-go/util"
	"strings"
)

const day = "day4"

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
	winningMap := map[string]int{}

	for _, line := range lines {
		input := strings.Split(strings.Split(line, ": ")[1], " | ")
		winning := strings.Split(input[0], " ")
		have := strings.Split(input[1], " ")

		for _, c := range winning {
			if len(c) == 0 {
				continue
			}
			winningMap[c] = 1
		}

		score := 0

		for _, c := range have {
			if winningMap[c] == 1 {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
		
		sum += score
		clear(winningMap)
	}

	return sum, nil
}
