package day1

import (
	"advent-of-code-go/util"
	"sort"
	"strconv"
	"strings"
)

const day = "day1"

var inputFile string

func Solve(easy bool) (name string, res int, err error) {
	name = day
	setInput(easy)
	lines, err := util.ReadStringList(inputFile)
	if err != nil {
		return
	}

	// Parse the two columns of numbers
	var firstNums, secondNums []int
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) != 2 {
			continue
		}

		first, err := strconv.Atoi(fields[0])
		if err != nil {
			return name, 0, err
		}

		second, err := strconv.Atoi(fields[1])
		if err != nil {
			return name, 0, err
		}

		firstNums = append(firstNums, first)
		secondNums = append(secondNums, second)
	}

	//res, err = partOne(firstNums, secondNums)
	// Uncomment for part two
	res, err = partTwo(firstNums, secondNums)

	return
}

func setInput(easy bool) {
	if easy {
		inputFile = day + "/" + util.InputFileEasy
	} else {
		inputFile = day + "/" + util.InputFileFull
	}
}

func partOne(firstNums, secondNums []int) (int, error) {
	// Sort both slices
	sort.Ints(firstNums)
	sort.Ints(secondNums)

	// Calculate sum of differences
	sum := 0
	for i := 0; i < len(firstNums); i++ {
		diff := firstNums[i] - secondNums[i]
		if diff < 0 {
			diff = -diff // Get absolute value
		}
		sum += diff
	}

	return sum, nil
}

func partTwo(firstNums, secondNums []int) (int, error) {
	// Create frequency map of second slice
	freqMap := make(map[int]int)
	for _, num := range secondNums {
		freqMap[num]++
	}

	// Calculate sum using the frequency map
	sum := 0
	for _, num := range firstNums {
		sum += num * freqMap[num]
	}
	return sum, nil
}
