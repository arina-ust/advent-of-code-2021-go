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
	//	res, err = partTwo(lines)

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

					number := number{
						xStart: x - len(num),
						xEnd:   x - 1,
						y:      y,
						value:  n,
					}
					c1 := coord{x: x - 1, y: y}
					coordsToNum[c1] = number
					if len(num) >= 2 {
						c2 := coord{x: x - 2, y: y}
						coordsToNum[c2] = number
					}
					if len(num) >= 3 {
						c3 := coord{x: x - 3, y: y}
						coordsToNum[c3] = number
					}
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

			number := number{
				xStart: len(line) - len(num),
				xEnd:   len(line) - 1,
				y:      y,
				value:  n,
			}
			c1 := coord{x: len(line) - 1, y: y}
			coordsToNum[c1] = number
			if len(num) >= 2 {
				c2 := coord{x: len(line) - 2, y: y}
				coordsToNum[c2] = number
			}
			if len(num) >= 3 {
				c3 := coord{x: len(line) - 3, y: y}
				coordsToNum[c3] = number
			}
		}

		// start over
		num = ""
		isAdj = false
	}

	return partTwo(d, symbolsMatrix)
	//	return sum, nil // part one
}

var validDigit = regexp.MustCompile("\\d")
var validDot = regexp.MustCompile("\\.")

func isDigit(s string) bool {
	return validDigit.MatchString(s)
}

func isDot(s string) bool {
	return validDot.MatchString(s)
}

type coord struct {
	x int
	y int
}

type number struct {
	xStart int
	xEnd   int
	y      int
	value  int
}

var coordsToNum = map[coord]number{}

func partTwo(d int, symbolsMatrix [][]string) (int, error) {
	sum := 0

	for y := 0; y < d; y++ {
		for x := 0; x < d; x++ {
			if symbolsMatrix[y][x] == "*" {
				// check for adjacent numbers
				var nums []number

				xm1, xp1, ym1, yp1 := x-1, x+1, y-1, y+1
				n, ok := coordsToNum[coord{x: xm1, y: ym1}]
				if xm1 < d && ym1 < d && xm1 >= 0 && ym1 >= 0 && ok {
					nums = append(nums, n)
				}
				n, ok = coordsToNum[coord{x: x, y: ym1}]
				if ym1 < d && ym1 >= 0 && ok {
					nums = append(nums, n)
				}
				n, ok = coordsToNum[coord{x: xp1, y: ym1}]
				if xp1 < d && ym1 < d && xp1 >= 0 && ym1 >= 0 && ok {
					nums = append(nums, n)
				}

				n, ok = coordsToNum[coord{x: xm1, y: yp1}]
				if xm1 < d && yp1 < d && xm1 >= 0 && yp1 >= 0 && ok {
					nums = append(nums, n)
				}
				n, ok = coordsToNum[coord{x: x, y: yp1}]
				if yp1 < d && yp1 >= 0 && ok {
					nums = append(nums, n)
				}

				n, ok = coordsToNum[coord{x: xp1, y: yp1}]
				if xp1 < d && yp1 < d && xp1 >= 0 && yp1 >= 0 && ok {
					nums = append(nums, n)
				}

				n, ok = coordsToNum[coord{x: xm1, y: y}]
				if xm1 < d && xm1 >= 0 && ok {
					nums = append(nums, n)
				}

				n, ok = coordsToNum[coord{x: xp1, y: y}]
				if xp1 < d && xp1 >= 0 && ok {
					nums = append(nums, n)
				}

				seen := map[number]bool{}
				var res []number
				for _, n := range nums {
					if _, ok := seen[n]; !ok {
						seen[n] = true
						res = append(res, n)
					}
				}

				if len(res) == 2 {
					sum += res[0].value * res[1].value
				}
			}
		}
	}

	return sum, nil
}
