package day9

import (
	"advent-of-code-go/util"
	"strconv"
	"strings"
)

const day = "day9"

var inputFile string

func Solve(easy bool) (name string, res int, err error) {
	name = day
	setInput(easy)
	lines, err := util.ReadStringList(inputFile)
	if err != nil {
		return
	}

//	res, err = partOne(lines)
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


func partOne(lines []string) (int, error) {
	res := 0
	
	for _, line := range lines {
		values := strings.Split(line, " ")
		
		// 6 12 17
		var placeholderIndices []int
		// 0 3 6 9 12 15 0(p) 3 3 3 3 3 0(p) 0 0 0 0 0(p)
		var calculations []int
		
		placeholderIndices = append(placeholderIndices, len(values))
		
		for _, v := range values {
			num, _ := strconv.Atoi(v)
			calculations = append(calculations, num)
		}
		calculations = append(calculations, 0) // append placeholder
		
		j := 0
		k := 0
		for {
			if j+1 == placeholderIndices[k] {
				k++
				placeholderIndices = append(placeholderIndices, len(calculations))
				calculations = append(calculations, 0) // append placeholder
				
				sumDiffs := 0
				for m := placeholderIndices[len(placeholderIndices)-1] - 1; m > placeholderIndices[len(placeholderIndices)-2]; m-- {
					sumDiffs += calculations[m]
				}
				if sumDiffs == 0 {
					break
				}
				
				j += 2
			}
			if j == len(calculations) {
				break
			}
			
			diff := calculations[j+1] - calculations[j]
			calculations = append(calculations, diff)
			
			j++
		}
		
		for i := len(placeholderIndices) - 1; i > 0; i-- { // 3 2 1
			nextPI := i-1 // 2 1 0
			currI := placeholderIndices[i] // 17 12 6
			nextI := placeholderIndices[nextPI] // 12 6
			calculations[nextI] = calculations[nextI-1] + calculations[currI]
		}
		
		res += calculations[placeholderIndices[0]]
	}

	return res, nil
}

func partTwo(lines []string) (int, error) {
	res := 0

	for _, line := range lines {
		values := strings.Split(line, " ")

		// 0 7 13 18
		var placeholderIndices []int
		// 0(p) 0 3 6 9 12 15 0(p) 3 3 3 3 3 0(p) 0 0 0 0 0(p)
		var calculations []int

		calculations = append(calculations, 0) // append placeholder
		placeholderIndices = append(placeholderIndices, 0)

		for _, v := range values {
			num, _ := strconv.Atoi(v)
			calculations = append(calculations, num)
		}
		calculations = append(calculations, 0) // append placeholder
		placeholderIndices = append(placeholderIndices, len(calculations)-1)

		j := 1
		k := 1
		for {
			if j+1 == placeholderIndices[k] {
				k++
				placeholderIndices = append(placeholderIndices, len(calculations))
				calculations = append(calculations, 0) // append placeholder

				sumDiffs := 0
				for m := placeholderIndices[len(placeholderIndices)-1] - 1; m > placeholderIndices[len(placeholderIndices)-2]; m-- {
					sumDiffs += calculations[m]
				}
				if sumDiffs == 0 {
					break
				}

				j += 2
			}
			if j == len(calculations) {
				break
			}

			diff := calculations[j+1] - calculations[j]
			calculations = append(calculations, diff)

			j++
		}

		for i := len(placeholderIndices) - 2; i > 0; i-- { // 2 1
			nextPI := i-1 // 1 0
			currI := placeholderIndices[i] // 13 7 0
			nextI := placeholderIndices[nextPI] // 7 0
			calculations[nextI] = calculations[nextI+1] - calculations[currI]
		}
		res += calculations[placeholderIndices[0]]
	}

	return res, nil
}
