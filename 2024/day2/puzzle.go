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

	var numbers [][]int
	// Parse each line into a slice of integers
	for _, line := range lines {
		var lineNums []int
		fields := strings.Fields(line)
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				return name, 0, err
			}
			lineNums = append(lineNums, num)
		}
		numbers = append(numbers, lineNums)
	}

	//res, err = partOne(numbers)
	// Uncomment for part two
	res, err = partTwo(numbers)

	return
}

func setInput(easy bool) {
	if easy {
		inputFile = day + "/" + util.InputFileEasy
	} else {
		inputFile = day + "/" + util.InputFileFull
	}
}

func partOne(numbers [][]int) (int, error) {
	validCount := 0
	for _, line := range numbers {
		if isValidSequence(line) {
			validCount++
		}
	}
	return validCount, nil
}

func isValidSequence(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]

		if diff < -3 || diff > 3 || diff == 0 {
			isIncreasing = false
			isDecreasing = false
			break
		}

		if diff < 0 {
			isIncreasing = false
		}
		if diff > 0 {
			isDecreasing = false
		}
	}

	return isIncreasing || isDecreasing
}

func partTwo(numbers [][]int) (int, error) {
	validCount := 0

	for _, line := range numbers {
		// First check if original line is valid
		if isValidSequence(line) {
			validCount++
			continue
		}

		// Try removing each number
		isValid := false
		for i := 0; i < len(line); i++ {
			// Create new slice without number at index i
			modified := make([]int, 0, len(line)-1)
			modified = append(modified, line[:i]...)
			modified = append(modified, line[i+1:]...)

			if isValidSequence(modified) {
				isValid = true
				break
			}
		}

		if isValid {
			validCount++
		}
	}

	return validCount, nil
}
