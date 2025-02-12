package day3

import (
	"advent-of-code-go/util"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

const day = "day3"

var inputFile string
var mulRegex = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
var doRegex = regexp.MustCompile(`do\(\)`)
var dontRegex = regexp.MustCompile(`don't\(\)`)
var enabled = true // Moving this to package level

func Solve(easy bool) (name string, res int, err error) {
	name = day
	setInput(easy)
	lines, err := util.ReadStringList(inputFile)
	if err != nil {
		return
	}
	if len(lines) == 0 {
		return name, 0, nil
	}

	enabled = true // Reset at the start of solving, not for each line
	res = 0
	for i, line := range lines {
		fmt.Printf("Line %d starting with enabled=%v\n", i+1, enabled)
		//lineRes, err := partOne(line)
		lineRes, err := partTwo(line)
		if err != nil {
			return name, 0, err
		}
		res += lineRes
	}

	return
}

func setInput(easy bool) {
	if easy {
		inputFile = day + "/" + util.InputFileEasy
	} else {
		inputFile = day + "/" + util.InputFileFull
	}
}

func partOne(line string) (int, error) {
	sum := 0

	// Find all matches of mul(X,Y)
	matches := mulRegex.FindAllStringSubmatch(line, -1)
	fmt.Printf("Found matches: %v\n", matches)

	for _, match := range matches {
		// match[0] is the full match, match[1] is X, match[2] is Y
		x, err := strconv.Atoi(match[1])
		if err != nil {
			return 0, err
		}

		y, err := strconv.Atoi(match[2])
		if err != nil {
			return 0, err
		}

		fmt.Printf("Processing: mul(%d,%d) = %d\n", x, y, x*y)
		sum += x * y
	}

	return sum, nil
}

func partTwo(line string) (int, error) {
	sum := 0

	// Find all matches upfront
	doMatches := doRegex.FindAllStringIndex(line, -1)
	dontMatches := dontRegex.FindAllStringIndex(line, -1)
	mulMatches := mulRegex.FindAllStringSubmatchIndex(line, -1)

	//fmt.Printf("do() positions: %v\n", doMatches)
	//fmt.Printf("don't() positions: %v\n", dontMatches)
	//fmt.Printf("mul() positions: %v\n", mulMatches)

	// Create a sorted list of control points
	type controlPoint struct {
		pos  int
		isDo bool
	}
	controlPoints := make([]controlPoint, 0, len(doMatches)+len(dontMatches))
	for _, do := range doMatches {
		controlPoints = append(controlPoints, controlPoint{pos: do[0], isDo: true})
	}
	for _, dont := range dontMatches {
		controlPoints = append(controlPoints, controlPoint{pos: dont[0], isDo: false})
	}
	// Sort by position
	sort.Slice(controlPoints, func(i, j int) bool {
		return controlPoints[i].pos < controlPoints[j].pos
	})

	// Process each multiplication
	for _, match := range mulMatches {
		pos := match[0]

		// Find the last control point before this position
		lastControlPoint := -1
		for i, cp := range controlPoints {
			if cp.pos > pos {
				break
			}
			lastControlPoint = i
		}

		// If we found a control point, check if we're in a disabled state
		if lastControlPoint >= 0 {
			cp := controlPoints[lastControlPoint]
			if !cp.isDo {
				enabled = false
			} else {
				enabled = true
			}
		}

		// Extract the numbers
		x, err := strconv.Atoi(line[match[2]:match[3]])
		if err != nil {
			return 0, err
		}

		y, err := strconv.Atoi(line[match[4]:match[5]])
		if err != nil {
			return 0, err
		}

		result := x * y
		if enabled {
			sum += result
		}
	}

	return sum, nil
}
