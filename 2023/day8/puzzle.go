package day8

import (
	"advent-of-code-go/util"
	"fmt"
	"strings"
	"time"
)

const day = "day8"

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

type node struct {
	name  string
	left  string
	right string
}

func partOne(lines []string) (int, error) {
	t1 := time.Now()
	fmt.Println("started ", t1)
	instructions := lines[0]

	nodes := map[string]*node{}
	for _, line := range lines[2:] {
		s := strings.Split(line, " = (")
		v := strings.Split(s[1], ", ")
		node := &node{
			name:  s[0],
			left:  v[0],
			right: strings.Split(v[1], ")")[0],
		}
		nodes[node.name] = node
	}

	goal := "ZZZ"
	next := ""
	steps := 0
	nextToVisit := nodes["AAA"]
	resets := 0

outer:
	for t := 0; t < len(instructions); t++ {
		turn := instructions[t]
		steps++

		if nextToVisit.name == "TVL" {
			fmt.Println("close to goal")
		}

		if turn == 'R' {
			next = nextToVisit.right

		} else if turn == 'L' {
			next = nextToVisit.left

		}

		//fmt.Printf("next node is %s, with steps num %v\n", next, steps)

		if next == goal {
			break outer
		} else {
			nextToVisit = nodes[next]
			if t == len(instructions)-1 {
				resets++
				//fmt.Println("resetting instructions ", resets)
				t = -1
			}
		}
	}

	fmt.Println("ended, took ", time.Now().Sub(t1))
	return steps, nil
}
