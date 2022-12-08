package day8

import (
	"advent-of-code-go/util"
	"strconv"
)

const day = "day8"

var inputFile string

func Solve(easy bool) (name string, res string, err error) {
	name = day
	setInput(easy)
	matrix, err := util.ReadMatrix(inputFile)
	if err != nil {
		return
	}

	res, err = partOne(matrix)
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

func partOne(matrix [][]int) (string, error) {
	count := len(matrix)*2 + len(matrix[0])*2 - 4

	right := len(matrix[0]) - 1
	left := 0
	top := 0
	bottom := len(matrix) - 1

	for i := 1; i < right; i++ {
		for j := 1; j < bottom; j++ {
			treeHeight := matrix[i][j]

			// look right
			isVisible := true
			for c := j + 1; c <= right; c++ {
				if matrix[i][c] >= treeHeight {
					isVisible = false
					break
				}
			}
			if isVisible {
				count++
				continue
			}

			// look left
			isVisible = true
			for c := j - 1; c >= left; c-- {
				if matrix[i][c] >= treeHeight {
					isVisible = false
					break
				}
			}
			if isVisible {
				count++
				continue
			}

			// look up
			isVisible = true
			for c := i - 1; c >= top; c-- {
				if matrix[c][j] >= treeHeight {
					isVisible = false
					break
				}
			}
			if isVisible {
				count++
				continue
			}

			// look down
			isVisible = true
			for c := i + 1; c <= bottom; c++ {
				if matrix[c][j] >= treeHeight {
					isVisible = false
					break
				}
			}
			if isVisible {
				count++
				continue
			}
		}
	}

	return strconv.Itoa(count), nil
}
