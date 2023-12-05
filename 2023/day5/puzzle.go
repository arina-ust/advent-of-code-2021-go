package day5

import (
	"advent-of-code-go/util"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const day = "day5"

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

var title = regexp.MustCompile("[A-Za-z]")


// TODO for Part 2
// 1. Optimize import:
//   -put sources into map with source as key and list of dest, range, dx as value
//   -put sources into slice of sources
//   -sort the slice
//   -put into struct both slice and map?
//   -add struct to the list of structs (each struct = each map)
// 2. Optimize solution:
//   -do binary search in the slice of a given struct
//   -calculate mapped value based on the found value and corresponding values from the map 

func partOne(lines []string) (int, error) {
	minLoc := math.MaxInt

	mapIndex := -1
	for _, line := range lines[1:] {
		if len(line) == 0 {
			continue
		}
		if title.MatchString(line) {
			mapIndex++
			var workingMap []string
			input[mapIndex] = workingMap
		} else {
			input[mapIndex] = append(input[mapIndex], line)
		}
	}
	
	seeds := getSeeds(strings.Split(strings.Split(lines[0], ": ")[1], " "))

	// for each seed
	for _, seed := range seeds {
		res, err := findLocation(seed)
		if err != nil {
			return -1, err
		}
		if res < minLoc {
			minLoc = res
		}
	}
	return minLoc, nil
}

func getSeeds(inputSeeds []string) []string {
	var res []string
	for i, j := 0, 1; j < len(inputSeeds); j += 2 {
		start := inputSeeds[i]
		rng := inputSeeds[j]

		s, err := strconv.Atoi(start)
		if err != nil {
			return res
		}

		r, err := strconv.Atoi(rng)
		if err != nil {
			return res
		}

		for k := s; k < s + r; k++ {
			res = append(res, strconv.Itoa(k))
		}

		i += 2
	}
	return res
}

var input [7][]string
var mappings [7][][]int

func findLocation(seed string) (int, error) {
	s, err := strconv.Atoi(seed)
	if err != nil {
		return -1, err
	}

	// for each map which is rather an row in array[len 7] of arrays
	res := s
	for i, _ := range mappings {
		if len(mappings[i]) == 0 {
			err := fillMappings(i, input[i])
			if err != nil {
				return -1, err
			}
		}
		res, err = findMappedValue(res, i)
		if err != nil {
			return -1, err
		}
	}
	return res, nil
}

func fillMappings(index int, oneMap []string) error {
	for _, line := range oneMap {
		processedInput := strings.Split(line, " ")
		destSourceRange := make([]int, len(processedInput))

		for i, v := range processedInput {
			atoi, err := strconv.Atoi(v)
			if err != nil {
				return err
			}
			destSourceRange[i] = atoi
		}

		mappings[index] = append(mappings[index], destSourceRange)
	}

	return nil
}

func findMappedValue(from int, index int) (int, error) {
	dest, source, err := findDestAndSource(from, index)
	if err != nil {
		return -1, err
	}

	// if no interval, return with (from, nil)
	if dest == -1 || source == -1 {
		return from, nil
	}
	// calculate dx = dest - source
	dx := dest - source

	// return from + dx
	return from + dx, nil
}

func findDestAndSource(from int, index int) (int, int, error) {
	dest, source := -1, -1

	for _, destSourceRange := range mappings[index] {

		// find corresponding interval of dest and source
		d := destSourceRange[0]
		s := destSourceRange[1]
		r := destSourceRange[2]

		if from >= s && from < (s+r) {
			source = s
			dest = d
		}
	}

	return dest, source, nil
}
