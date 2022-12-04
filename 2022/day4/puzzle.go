package day4

import (
	"advent-of-code-go/util"
	"strconv"
	"strings"
)

const day = "day4"

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

type elfPair struct {
	s1 []int
	s2 []int
}

func (ep elfPair) hasFullOverlap() bool {
	// if ep.s2[0] > ep.s1[1] || ep.s2[1] < ep.s1[0] { // no overlap
	// 	return false
	// }
	if (ep.s1[0] <= ep.s2[0] && ep.s1[1] >= ep.s2[1]) || (ep.s1[0] >= ep.s2[0] && ep.s1[1] <= ep.s2[1]) {
		return true
	}
	return false
}

func partOne(lines []string) (string, error) {
	count := 0
	for _, line := range lines {
		sections := strings.Split(line, ",")
		s1 := strings.Split(sections[0], "-")
		s2 := strings.Split(sections[1], "-")
		x, err := strconv.Atoi(s1[0])
		y, err := strconv.Atoi(s1[1])
		z, err := strconv.Atoi(s2[0])
		w, err := strconv.Atoi(s2[1])
		if err != nil {
			return "", err
		}

		ep := elfPair{
			s1: []int{x, y},
			s2: []int{z, w},
		}

		if ep.hasFullOverlap() {
			count++
		}

	}
	return strconv.Itoa(count), nil
}
