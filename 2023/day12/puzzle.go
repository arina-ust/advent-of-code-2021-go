package day12

import (
	"advent-of-code-go/util"
	"fmt"
	"strconv"
	"strings"
)

const day = "day12"

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

func partOne(lines []string) (int, error) {
	sum := 0

	for _, line := range lines {
		fmt.Println(line)

		input := strings.Split(line, " ")
		damagedStr := strings.Split(input[1], ",")
		var damaged []int
		for _, s := range damagedStr {
			v, _ := strconv.Atoi(s)
			damaged = append(damaged, v)
		}
		sum += countArrangements(input[0], damaged)
	}

	return sum, nil
}

type broken struct {
	startIndex int
	length     int
}

func countArrangements(s string, damaged []int) int {
	var knownBrokens []broken
	b := broken{}
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "#" {
			b.startIndex = i
			b.length = 1

			for j := i + 1; j < len(s); j++ {
				if string(s[j]) == "#" {
					b.length += 1
				} else {
					break
				}
			}
			i += b.length

			knownBrokens = append(knownBrokens, b)
			b = broken{}
		}
	}
	fmt.Println(knownBrokens)

	var expectedBroken []broken
	for _, d := range damaged {
		expectedBroken = append(expectedBroken, broken{length: d})
	}

	for _, eb := range expectedBroken {
		countBroken := 0
		countUnknown := 0
		for j := 0; j < len(s) && eb.length > countBroken; j++ {
			ch := string(s[j])
			if ch == "." {
				countBroken = 0
				countUnknown = 0
				continue
			} else if ch == "#" {
				countBroken++
			} else {
				countUnknown++
			}
		}
	}

	restored := ""

	fmt.Println(restored)
	return 1
}
