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

func countArrangements(s string, damaged []int) int {
	restored := "" 
	j := 0
	canNextBeDamaged := true
	for i := 0; i < len(s); i ++{
		ch := string(s[i])
		
		if ch == "." {
			canNextBeDamaged = true
		}
		
		if ch == "?" {
			if j < len(damaged) && damaged[j] > 0 && canNextBeDamaged{
				restored += "#"
				damaged[j] -= 1
			} else {
				restored += "."
				canNextBeDamaged = true
			}
		} else {
			if ch == "#" && j < len(damaged) {
				damaged[j] -= 1
			}
			restored += ch
		}
		if j < len(damaged)&& damaged[j] == 0 {
			j++
			canNextBeDamaged = false
		} 
	}
	fmt.Println(restored)
	return 1
}
