package day3

import (
	"advent-of-code-go/util"
	"fmt"
	"regexp"
	"strconv"
)

const day = "day3"

var inputFile string

func Solve(easy bool) (name string, res int, err error) {
	name = day
	setInput(easy)
	lines, err := util.ReadStringList(inputFile)
	if err != nil {
		return
	}

	res, err = partOne(lines)
	//			res, err = partTwo(lines)

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

	d := len(lines)
	symbolsMatrix := util.GetEmptyMatrix[string](d, d)

	for y, line := range lines {
		for x, ch := range line {
			s := string(ch)
			if !isDigit(s) && !isDot(s) {
				symbolsMatrix[y][x] = s
			}
		}
	}

	for y, line := range lines {
		num := ""
		isAdj := false
		for x, ch := range line {
			s := string(ch)
			if isDigit(s) {
				num += s

				if isAdj {
					continue
				}

				// check for adjacent symbols
				xm1, xp1, ym1, yp1 := x-1, x+1, y-1, y+1
				if xm1 < d && ym1 < d && xm1 >= 0 && ym1 >= 0 && len(symbolsMatrix[ym1][xm1]) != 0 {
					isAdj = true
				}
				if ym1 < d && ym1 >= 0 && len(symbolsMatrix[ym1][x]) != 0 {
					isAdj = true
				}
				if xp1 < d && ym1 < d && xp1 >= 0 && ym1 >= 0 && len(symbolsMatrix[ym1][xp1]) != 0 {
					isAdj = true
				}
				if xm1 < d && xm1 >= 0 && len(symbolsMatrix[y][xm1]) != 0 {
					isAdj = true
				}
				if xp1 < d && xp1 >= 0 && len(symbolsMatrix[y][xp1]) != 0 {
					isAdj = true
				}
				if xm1 < d && yp1 < d && xm1 >= 0 && yp1 >= 0 && len(symbolsMatrix[yp1][xm1]) != 0 {
					isAdj = true
				}
				if yp1 < d && yp1 >= 0 && len(symbolsMatrix[yp1][x]) != 0 {
					isAdj = true
				}
				if xp1 < d && yp1 < d && xp1 >= 0 && yp1 >= 0 && len(symbolsMatrix[yp1][xp1]) != 0 {
					isAdj = true
				}
			} else {
				if isAdj {
					n, err := strconv.Atoi(num)
					if err != nil {
						return -1, fmt.Errorf("failed to convert to int %s", num)
					}
					sum += n
				}

				// start over
				num = ""
				isAdj = false
			}
		}
		if isAdj {
			n, err := strconv.Atoi(num)
			if err != nil {
				return -1, fmt.Errorf("failed to convert to int %s", num)
			}
			sum += n
		}

		// start over
		num = ""
		isAdj = false
	}
	return sum, nil
}

var validDigit = regexp.MustCompile("\\d")
var validDot = regexp.MustCompile("\\.")

func isDigit(s string) bool {
	return validDigit.MatchString(s)
}

func isDot(s string) bool {
	return validDot.MatchString(s)
}

func partTwo(lines []string) (int, error) {
	sum := 0
	//	for _, line := range lines {
	//
	//	}
	return sum, nil
}
