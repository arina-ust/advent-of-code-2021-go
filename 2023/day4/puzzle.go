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

	played := make([]int, len(lines))

	for i, line := range lines {
		input := strings.Split(strings.Split(line, ": ")[1], " | ")
		winning := strings.Split(input[0], " ")
		have := strings.Split(input[1], " ")

		for _, c := range winning {
			if len(c) == 0 {
				continue
			}
			winningMap[c] = 1
		}

		played[i] += 1

		score := 0
		numOfWins := 0
		for _, c := range have {
			if winningMap[c] == 1 {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
				numOfWins++
			}
		}
		
		for j := 0; j < numOfWins; j++ {
			played[i+1+j] += played[i]
		}

		sum += score
		clear(winningMap)
	}

	//	return sum, nil // for part one

	numCards := 0
	for _, v := range played {
		numCards += v
	}
	return numCards, nil
}
