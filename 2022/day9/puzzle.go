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

	res, err = partOne(lines)
	// res, err = partTwo(matrix)

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
	head          head
	tail          tail
	tailPositions map[point]bool
}

type head struct {
	coord point
}

type tail struct {
	coord point
}

type point struct {
	x int
	y int
}

func (r *rope) areTouchingHeadAndTail() bool {
	if r.head.coord.x == r.tail.coord.x && r.head.coord.y == r.tail.coord.y {
		return true // same position
	} else if r.head.coord.x == r.tail.coord.x && math.Abs(float64(r.head.coord.y-r.tail.coord.y)) == 1 {
		return true // touch vertically
	} else if r.head.coord.y == r.tail.coord.y && math.Abs(float64(r.head.coord.x-r.tail.coord.x)) == 1 {
		return true // touch horizontally
	} else if math.Abs(float64(r.head.coord.x-r.tail.coord.x)) == 1 && math.Abs(float64(r.head.coord.y-r.tail.coord.y)) == 1 {
		return true // touch diagonally
	}
	return false
}

func (r *rope) moveRight() {
	r.head.coord.x = r.head.coord.x + 1
	r.moveTail()
}

func (r *rope) noteTailsPosition() {
	r.tailPositions[r.tail.coord] = true
}

func (r *rope) moveTail() {
	if r.areTouchingHeadAndTail() {
		return
	}

	if r.head.coord.x-r.tail.coord.x == 2 {
		if r.head.coord.y-r.tail.coord.y == 0 {
			r.tail.coord.x = r.tail.coord.x + 1
		} else if r.head.coord.y-r.tail.coord.y == 1 {
			r.tail.coord.x = r.tail.coord.x + 1
			r.tail.coord.y = r.tail.coord.y + 1
		} else if r.head.coord.y-r.tail.coord.y == -1 {
			r.tail.coord.x = r.tail.coord.x + 1
			r.tail.coord.y = r.tail.coord.y - 1
		}
	} else if r.head.coord.x-r.tail.coord.x == -2 {
		if r.head.coord.y-r.tail.coord.y == 0 {
			r.tail.coord.x = r.tail.coord.x - 1
		} else if r.head.coord.y-r.tail.coord.y == 1 {
			r.tail.coord.x = r.tail.coord.x - 1
			r.tail.coord.y = r.tail.coord.y + 1
		} else if r.head.coord.y-r.tail.coord.y == -1 {
			r.tail.coord.x = r.tail.coord.x - 1
			r.tail.coord.y = r.tail.coord.y - 1
		}
	} else if r.head.coord.y-r.tail.coord.y == 2 {
		if r.head.coord.x-r.tail.coord.x == 0 {
			r.tail.coord.y = r.tail.coord.y + 1
		} else if r.head.coord.x-r.tail.coord.x == 1 {
			r.tail.coord.y = r.tail.coord.y + 1
			r.tail.coord.x = r.tail.coord.x + 1
		} else if r.head.coord.x-r.tail.coord.x == -1 {
			r.tail.coord.y = r.tail.coord.y + 1
			r.tail.coord.x = r.tail.coord.x - 1
		}
	} else if r.head.coord.y-r.tail.coord.y == -2 {
		if r.head.coord.x-r.tail.coord.x == 0 {
			r.tail.coord.y = r.tail.coord.y - 1
		} else if r.head.coord.x-r.tail.coord.x == 1 {
			r.tail.coord.y = r.tail.coord.y - 1
			r.tail.coord.x = r.tail.coord.x + 1
		} else if r.head.coord.x-r.tail.coord.x == -1 {
			r.tail.coord.y = r.tail.coord.y - 1
			r.tail.coord.x = r.tail.coord.x - 1
		}
	} else {
		fmt.Println("!!!!!Shouldn't go in here?????")
	}
	r.noteTailsPosition()
}

func (r *rope) moveLeft() {
	r.head.coord.x = r.head.coord.x - 1
	r.moveTail()
}

func (r *rope) moveUp() {
	r.head.coord.y = r.head.coord.y - 1
	r.moveTail()
}

func (r *rope) moveDown() {
	r.head.coord.y = r.head.coord.y + 1
	r.moveTail()
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

func partOne(lines []string) (string, error) {

	p := point{
		x: 0,
		y: 0,
	}
	m := map[point]bool{
		p: true,
	}

	rope := &rope{
		head: head{
			coord: p,
		},
		tail: tail{
			coord: p,
		},
		tailPositions: m,
	}

	for _, line := range lines {
		motion := strings.Split(line, " ")

		steps, err := strconv.Atoi(motion[1])
		if err != nil {
			return "", err
		}

		var f func()
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

		for i := 1; i <= steps; i++ {
			f()
		}
	}

	count := rope.calculateTailPositions()

	return strconv.Itoa(count), nil
}
