package day5

import (
	"advent-of-code-go/util"
	"strconv"
	"strings"
)

const day = "day5"

var inputFile string

func Solve(easy bool) (name string, res string, err error) {
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

func partOne(lines []string) (string, error) {
	var numStacks int = 9 + 1 // for full task; 3 + 1 for easy

	stacks := make([][]rune, numStacks)

	working := false
	for _, line := range lines {
		if line != "" && !working && line[1] != '1' { // reading initial stacks setup
			i := 0
			for _, r := range line {
				if r == ' ' {
					i++
				} else if r == '[' || r == ']' {
					i++
					continue
				} else {
					i++
					var st int
					if i < 4 {
						st = 1
					} else {
						st = i/4 + 1
					}
					stacks[st] = append(stacks[st], r)
				}
			}
		} else if line == "" {
			working = true
			for _, stack := range stacks {
				for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
					stack[i], stack[j] = stack[j], stack[i]
				}
			}
			continue
		} else if line[1] == '1' {
			continue
		}

		if working { // reading and implemeting instructions
			instructions := strings.Split(line, " ")
			numCrates, _ := strconv.Atoi(instructions[1])
			fromStack, _ := strconv.Atoi(instructions[3])
			toStack, _ := strconv.Atoi(strings.Trim(instructions[5], "\n"))

			for j := 0; j < numCrates; j++ {
				crate := stacks[fromStack][len(stacks[fromStack])-1]
				stacks[toStack] = append(stacks[toStack], crate)
				stacks[fromStack] = stacks[fromStack][:len(stacks[fromStack])-1]
			}
		}

	}
	var result string = ""
	for _, stack := range stacks[1:] {
		lastStack := len(stack) - 1
		result += string(stack[lastStack])
	}

	return result, nil
}
