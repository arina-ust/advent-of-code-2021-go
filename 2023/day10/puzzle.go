package day10

import (
	"advent-of-code-go/util"
	"fmt"
)

const day = "day10"

var inputFile string

func Solve(easy bool) (name string, res int, err error) {
	name = day
	setInput(easy)
	lines, err := util.ReadMatrixString(inputFile)
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

func partOne(lines [][]string) (int, error) {
	res := 0

	// [1 2 3]
	// [1 2 3]
	// [1 2 3]

	colNum := len(lines[0]) // x
	rowNum := len(lines)    // y

	colNumOfS := 0
	rowNumofS := 0

outer:
	for i := 0; i < rowNum; i++ {
		for j := 0; j < colNum; j++ {
			if lines[i][j] == "S" {
				colNumOfS = j
				rowNumofS = i
				break outer
			}
		}
	}

	nextColNum := colNumOfS
	nextRowNum := rowNumofS
	
	previousCol := nextColNum
	previousRow := nextRowNum

	loop := map[node]bool{}
	for {
		loop[node{value: lines[nextRowNum][nextColNum], col: nextColNum, row: nextRowNum}] = true

		nextRowNum, nextColNum, previousRow, previousCol = nextMove(lines, nextRowNum, nextColNum, previousRow, previousCol)
		
		res++
		
		if lines[nextRowNum][nextColNum] == "S" {
			break
		}
	}

	countInsideLoopTiles := 0

	for i := 0; i < rowNum; i++ {
		for j := 0; j < colNum; j++ {
			nd := node{value: lines[i][j], col: j, row: i}
			_, ok := loop[nd]
			if !ok {
				//                           up  down  right left
				var neighbours = []int{i + 1, i - 1, j + 1, j - 1}
				countLoopNeighbours := 0

				for k, n := range neighbours {
					if n >= colNum || n >= rowNum || n < 0{
						continue
					}
					var value node
					if k == 0 || k == 1 { // up or down
						value = node{value: lines[n][j], col: j, row: n,}
					} else { // left or right
						value = node{value: lines[i][n], col: n, row: i,}
					}
					if _, ok := loop[value]; ok {
						countLoopNeighbours++
					}
				}
				if countLoopNeighbours == 4 { // TODO incorrect, there could be neighbors that are not in the loop but still enclosed
					countInsideLoopTiles++
				}
			}
		}
	}


//	return res / 2, nil // part one
	return countInsideLoopTiles, nil // part two
}

func nextMove(lines [][]string, i int, j int, prevI, prevJ int) (int, int, int, int) {
	colNum := len(lines[0]) // x
	rowNum := len(lines)    // y

	currValue := &node{value: lines[i][j], col: j, row:i,}

	//                           up  down  right left
	var neighbours = []int{i + 1, i - 1, j + 1, j - 1}

	for k, n := range neighbours {
		if n >= colNum || n >= rowNum || n < 0{
			continue
		}
		var value *node
		if k == 0 || k == 1 { // up or down
			
			if n == prevI && j == prevJ {
				continue
			}
			
			value = &node{value: lines[n][j], col: j, row: n,}
			if currValue.canMoveTo(value) {
				return n, j, i, j
			}
		} else { // left or right
			
			if n == prevJ && i == prevI {
				continue
			}
			
			value = &node{value: lines[i][n], col: n, row: i,}
			if currValue.canMoveTo(value) {
				return i, n, i, j
			}
		}
	}
	fmt.Println("shouldn't be here in next move")
	return -1, -1, i, j
}

type node struct {
	value string
	col   int
	row   int
}

func (n *node) canMoveTo(other *node) bool {
	if n.value == "|" {
		if n.col == other.col {
			if n.row > other.row { // up
				return other.value == "|" || other.value == "F" || other.value == "7" || other.value == "S"
			}
			if n.row < other.row { // down
				return other.value == "|" || other.value == "L" || other.value == "J" || other.value == "S"
			}
		}
	}
	if n.value == "-" {
		if n.row == other.row {
			if n.col < other.col { // to the right
				return other.value == "-" || other.value == "7" || other.value == "J" || other.value == "S"
			}
			if n.col > other.col { // to the left
				return other.value == "-" || other.value == "L" || other.value == "F" || other.value == "S"
			}
		}
	}
	if n.value == "L" {
		if n.col == other.col {
			if n.row > other.row { // up
				return other.value == "|" || other.value == "F" || other.value == "7" || other.value == "S"
			}
		}
		if n.row == other.row {
			if n.col < other.col { // to the right
				return other.value == "-" || other.value == "7" || other.value == "J" || other.value == "S"
			}
		}
	}
	if n.value == "J" {
		if n.col == other.col {
			if n.row > other.row { // up
				return other.value == "|" || other.value == "F" || other.value == "7" || other.value == "S"
			}
		}
		if n.row == other.row {
			if n.col > other.col { // to the left
				return other.value == "-" || other.value == "L" || other.value == "F" || other.value == "S"
			}
		}
	}
	if n.value == "7" {
		if n.col == other.col {
			if n.row < other.row { // down
				return other.value == "|" || other.value == "L" || other.value == "J" || other.value == "S"
			}
		}
		if n.row == other.row {
			if n.col > other.col { // to the left
				return other.value == "-" || other.value == "L" || other.value == "F" || other.value == "S"
			}
		}
	}
	if n.value == "F" {
		if n.col == other.col {
			if n.row < other.row { // down
				return other.value == "|" || other.value == "L" || other.value == "J" || other.value == "S"
			}
		}
		if n.row == other.row {
			if n.col < other.col { // to the right
				return other.value == "-" || other.value == "7" || other.value == "J" || other.value == "S"
			}
		}
	}
	if other.value == "." {
		return false
	}
	if n.value == "S" {
		if n.col == other.col {
			if n.row > other.row { // up
				return other.value == "|" || other.value == "F" || other.value == "7"
			}
			if n.row < other.row { // down
				return other.value == "|" || other.value == "L" || other.value == "J"
			}
		}
		if n.row == other.row {
			if n.col < other.col { // to the right
				return other.value == "-" || other.value == "7" || other.value == "J"
			}
			if n.col > other.col { // to the left
				return other.value == "-" || other.value == "L" || other.value == "F"
			}
		}
	}
	fmt.Println("shouldn't be here in can move")
	return false
}
