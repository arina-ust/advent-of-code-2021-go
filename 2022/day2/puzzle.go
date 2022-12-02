package day2

import (
	"advent-of-code-go/util"
	"strconv"
)

const day = "day2"

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

func partOne(lines []string) (string, error) {
	total := 0
	for _, line := range lines {
		total += scores[rune(line[0])][rune(line[2])]
	}
	return strconv.Itoa(total), nil
}

// part 1
//          1   2   3
// Op/Resp  R   P   S
//        R 1+3 2+6 3+0
//        P 1+0 2+3 3+6
//        S 1+6 2+0 3+3
//

var scores = map[rune]map[rune]int{
	'A': {
		'X': 4,
		'Y': 8,
		'Z': 3,
	},
	'B': {
		'X': 1,
		'Y': 5,
		'Z': 9,
	},
	'C': {
		'X': 7,
		'Y': 2,
		'Z': 6,
	},
}

// part 2
// Score  1   2   3
//        R   P   S
//
//          0 3 6
// Op/Res   X Y Z
//        R S R P
//        P R P S
//        S P S R
//
//          0 3 6
// Op/Res   X Y Z
//        R 0+3 3+1 6+2
//        P 0+1 3+2 6+3
//        S 0+2 3+3 6+1

var scores2 = map[rune]map[rune]int{
	'A': {
		'X': 3,
		'Y': 4,
		'Z': 8,
	},
	'B': {
		'X': 1,
		'Y': 5,
		'Z': 9,
	},
	'C': {
		'X': 2,
		'Y': 6,
		'Z': 7,
	},
}

func partTwo(lines []string) (string, error) {
	total := 0
	for _, line := range lines {
		total += scores2[rune(line[0])][rune(line[2])]
	}
	return strconv.Itoa(total), nil
}
