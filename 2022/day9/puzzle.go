package day9

import (
	"advent-of-code-go/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const day = "day9"

var inputFile string

func Solve(easy bool) (name string, res string, err error) {
	name = day
	setInput(easy)
	lines, err := util.ReadStringList(inputFile)
	if err != nil {
		return
	}

	// res, err = partOne(lines, 2)
	res, err = partOne(lines, 10)

	return
}

func setInput(easy bool) {
	if easy {
		inputFile = day + "/" + util.InputFileEasy
	} else {
		inputFile = day + "/" + util.InputFileFull
	}
}

type rope struct {
	knots         []point
	tailPositions map[point]bool
}

type point struct {
	x int
	y int
}

func (r *rope) areTouchingHeadAndTail(i, j int) bool {
	if r.knots[i].x == r.knots[j].x && r.knots[i].y == r.knots[j].y {
		return true // same position
	} else if r.knots[i].x == r.knots[j].x && math.Abs(float64(r.knots[i].y-r.knots[j].y)) == 1 {
		return true // touch vertically
	} else if r.knots[i].y == r.knots[j].y && math.Abs(float64(r.knots[i].x-r.knots[j].x)) == 1 {
		return true // touch horizontally
	} else if math.Abs(float64(r.knots[i].x-r.knots[j].x)) == 1 && math.Abs(float64(r.knots[i].y-r.knots[j].y)) == 1 {
		return true // touch diagonally
	}
	return false
}

func (r *rope) moveRight(i int) {
	r.knots[i].x = r.knots[i].x + 1
	r.moveTail(i, i+1)
}

func (r *rope) noteTailsPosition() {
	tail := len(r.knots) - 1
	r.tailPositions[r.knots[tail]] = true
}

func (r *rope) moveTail(i, j int) {
	if r.areTouchingHeadAndTail(i, j) {
		return
	}

	if r.knots[i].x-r.knots[j].x == 2 {
		if r.knots[i].y-r.knots[j].y == 0 {
			r.knots[j].x = r.knots[j].x + 1
		} else if r.knots[i].y-r.knots[j].y >= 1 {
			r.knots[j].x = r.knots[j].x + 1
			r.knots[j].y = r.knots[j].y + 1
		} else if r.knots[i].y-r.knots[j].y <= -1 {
			r.knots[j].x = r.knots[j].x + 1
			r.knots[j].y = r.knots[j].y - 1
		}
	} else if r.knots[i].x-r.knots[j].x == -2 {
		if r.knots[i].y-r.knots[j].y == 0 {
			r.knots[j].x = r.knots[j].x - 1
		} else if r.knots[i].y-r.knots[j].y >= 1 {
			r.knots[j].x = r.knots[j].x - 1
			r.knots[j].y = r.knots[j].y + 1
		} else if r.knots[i].y-r.knots[j].y <= -1 {
			r.knots[j].x = r.knots[j].x - 1
			r.knots[j].y = r.knots[j].y - 1
		}
	} else if r.knots[i].y-r.knots[j].y == 2 {
		if r.knots[i].x-r.knots[j].x == 0 {
			r.knots[j].y = r.knots[j].y + 1
		} else if r.knots[i].x-r.knots[j].x >= 1 {
			r.knots[j].y = r.knots[j].y + 1
			r.knots[j].x = r.knots[j].x + 1
		} else if r.knots[i].x-r.knots[j].x <= -1 {
			r.knots[j].y = r.knots[j].y + 1
			r.knots[j].x = r.knots[j].x - 1
		}
	} else if r.knots[i].y-r.knots[j].y == -2 {
		if r.knots[i].x-r.knots[j].x == 0 {
			r.knots[j].y = r.knots[j].y - 1
		} else if r.knots[i].x-r.knots[j].x >= 1 {
			r.knots[j].y = r.knots[j].y - 1
			r.knots[j].x = r.knots[j].x + 1
		} else if r.knots[i].x-r.knots[j].x <= -1 {
			r.knots[j].y = r.knots[j].y - 1
			r.knots[j].x = r.knots[j].x - 1
		}
	} else {
		fmt.Println("!!!!!Shouldn't go in here?????")
	}
}

func (r *rope) moveLeft(i int) {
	r.knots[i].x = r.knots[i].x - 1
	r.moveTail(i, i+1)
}

func (r *rope) moveUp(i int) {
	r.knots[i].y = r.knots[i].y - 1
	r.moveTail(i, i+1)
}

func (r *rope) moveDown(i int) {
	r.knots[i].y = r.knots[i].y + 1
	r.moveTail(i, i+1)
}

func (r *rope) calculateTailPositions() int {
	count := 0
	for _, v := range r.tailPositions {
		if v {
			count++
		}
	}
	return count
}

func partOne(lines []string, numKnots int) (string, error) {

	p := point{
		x: 0,
		y: 0,
	}
	m := map[point]bool{
		p: true,
	}

	knots := make([]point, numKnots)
	for k := 0; k < numKnots; k++ {
		knots[k] = p
	}

	rope := &rope{
		knots:         knots,
		tailPositions: m,
	}

	for _, line := range lines {
		motion := strings.Split(line, " ")

		steps, err := strconv.Atoi(motion[1])
		if err != nil {
			return "", err
		}

		var f func(int)
		switch motion[0] {
		case "R":
			f = rope.moveRight
		case "L":
			f = rope.moveLeft
		case "U":
			f = rope.moveUp
		case "D":
			f = rope.moveDown
		default:
			return "", fmt.Errorf("unknown motion")
		}

		for s := 0; s < steps; s++ {
			f(0) // head and a knot right after moves
			for i := 1; i < numKnots-1; i++ {
				rope.moveTail(i, i+1)
			}
			rope.noteTailsPosition()
		}
	}

	count := rope.calculateTailPositions()

	return strconv.Itoa(count), nil
}
