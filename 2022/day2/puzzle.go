package day2

import (
	"advent-of-code-go/util"
	"strconv"
)

const day = "day2"

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

const lost, draw, won int = 0, 3, 6

type round struct {
	name         string
	yourResponse bool
	score        int
}

func (r round) play(opponent round) int {
	if r.name == opponent.name {
		return draw
	}
	if r.name == "rock" {
		if opponent.name == "scissors" {
			return won
		}
		return lost
	}
	if r.name == "paper" {
		if opponent.name == "rock" {
			return won
		}
		return lost
	}
	if r.name == "scissors" {
		if opponent.name == "paper" {
			return won
		}
		return lost
	}
	return lost
}

var scores = map[rune]round{
	'A': round{
		name:         "rock",
		yourResponse: false,
		score:        1,
	},
	'B': round{
		name:         "paper",
		yourResponse: false,
		score:        2,
	},
	'C': round{
		name:         "scissors",
		yourResponse: false,
		score:        3,
	},
	'X': round{
		name:         "rock",
		yourResponse: true,
		score:        1,
	},
	'Y': round{
		name:         "paper",
		yourResponse: true,
		score:        2,
	},
	'Z': round{
		name:         "scissors",
		yourResponse: true,
		score:        3,
	},
}

func partOne(lines []string) (string, error) {
	total := 0
	for _, line := range lines {
		opponent, response := scores[rune(line[0])], scores[rune(line[2])]
		score := response.play(opponent) + response.score
		total += score
	}
	return strconv.Itoa(total), nil
}
