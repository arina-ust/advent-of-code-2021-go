package day1

import (
	"advent-of-code-go/util"
	"fmt"
	"regexp"
	"strconv"
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

var validDigit = regexp.MustCompile("\\d")

func partOne(lines []string) (int, error) {
	sum := 0
	for _, line := range lines {
		digits := validDigit.FindAllString(line, -1)
		var result string
		if len(digits) == 0 {
			continue
		}
		if len(digits) == 1 {
			result = digits[0] + digits[0]
		} else {
			last := len(digits) - 1
			result = digits[0] + digits[last]
		}
		value, err := strconv.Atoi(result)
		if err != nil {
			return -1, fmt.Errorf("failed to convert to int %s", result)
		}
		sum += value
	}
	return sum, nil
}

var digitsMap = map[string]string{"one": "1", "two": "2", "three": "3", "four": "4",
	"five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9",
	"1": "1", "2": "2", "3": "3", "4": "4", "5": "5", "6": "6",
	"7": "7", "8": "8", "9": "9",
}

func partTwo(lines []string) (int, error) {
	sum := 0
	for _, line := range lines {
		var digits []string
		i := 0
		j := 1
		for i < len(line) { //4eightwo
			if i == j {
				j++
			}
			ch := line[i:j]
			v, ok := digitsMap[ch]
			if ok {
				digits = append(digits, v)
				i++
				j = i + 1
			} else {
				j++
			}
			if j == len(line) + 1 {
				i++
				j = i + 1
			}
		}
		value, err := getValue(digits)
		if err != nil {
			return -1, err
		}
		sum += value
	}
	return sum, nil
}

func getValue(digits []string) (int, error) {
	var result string
	if len(digits) == 1 {
		d := digits[0]
		if len(d) == 1 {
			result = d + d
		} else {
			result = digitsMap[d] + digitsMap[d]
		}
	} else {
		first := digits[0]
		lastI := len(digits) - 1
		last := digits[lastI]
		if len(first) == 1 {
			result = first + digitsMap[last]
		} else if len(last) == 1 {
			result = digitsMap[first] + last
		} else {
			result = digitsMap[first] + digitsMap[last]
		}
	}
	value, err := strconv.Atoi(result)
	if err != nil {
		return -1, fmt.Errorf("failed to convert to int %s", result)
	}
	return value, nil
}
