package day2

import (
	"advent-of-code-go/util"
	"strconv"
	"strings"
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


var constraints = map[string]int{"red": 12, "green": 13, "blue": 14}

func partOne(lines []string) (int, error) {
	sum := 0
	for _, line := range lines {
		wasPossble := true
		
		instruction := strings.Split(line, ":")
		game := strings.Split(instruction[0], " ")
		gameNum, err := strconv.Atoi(game[1])
		if err != nil {
			return 0, err
		}
		
		shows := strings.Split(instruction[1], ";")
		for _, s := range shows {
			draws := strings.Split(s, ",")
			for _, d := range draws {
				res := util.RemoveWhiteSpaces(strings.Split(d, " "))
				constraint := constraints[res[1]]
				num, err := strconv.Atoi(res[0])
				if err != nil {
					return 0, err
				}
				if num > constraint {
					wasPossble = false
				}
			}
		}
		if wasPossble {
			sum += gameNum
		}
	}
	return sum, nil
}

