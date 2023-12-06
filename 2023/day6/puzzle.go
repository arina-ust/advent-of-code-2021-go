package day6

import (
	"advent-of-code-go/util"
	"strconv"
	"strings"
)

const day = "day6"

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

type race struct {
	time     int
	distance int
}

func partOne(lines []string) (int, error) {

	var races = make([]*race, 4) // 4 for full, 3 for easy

	times := strings.Split(strings.TrimSpace(strings.Split(lines[0], ": ")[1]), " ")
	j := 0
	for _, v := range times {
		s := strings.TrimSpace(v)
		if len(s) == 0 {
			continue
		}
		t, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		races[j] = &race{
			time: t,
		}
		j++
	}

	distances := strings.Split(strings.TrimSpace(strings.Split(lines[1], ": ")[1]), " ")
	j = 0
	for _, v := range distances {
		s := strings.TrimSpace(v)
		if len(s) == 0 {
			continue
		}
		d, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		races[j].distance = d
		j++
	}

	return calculate(races)
}

func calculate(races []*race) (int, error) {
	res := 1
	for _, r := range races {
		// from 1 to (r.time)
		// hold 1 ms -> speed 1 mm/ms => r.time - hold = distance * 1 ms * 1 mm/ms (7-1 = x * 1) x= 6
		// hold 2 ms -> speed 2 mm/ms => r.time - hold = distance * 1 ms * 2 mm/ms (7-2 = x * 2) x= 10
		// hold 3 ms -> speed 3 mm/ms => r.time - hold = distance * 1 ms * 3 mm/ms (7-3 = x * 3) x= 12
		// hold 4 ms -> speed 4 mm/ms => r.time - hold = distance * 1 ms * 4 mm/ms (7-4 = x * 4) x= 12
		// hold 5 ms -> speed 5 mm/ms => r.time - hold = distance * 1 ms * 5 mm/ms (7-5 = x * 5) x= 10
		// hold 6 ms -> speed 6 mm/ms => r.time - hold = distance * 1 ms * 6 mm/ms (7-6 = x * 6) x= 6
		record := r.distance
		wins := 0
		for i := 1; i < r.time; i++ {
			x := i * (r.time - i)
			if x > record {
				wins++
			}
		}

		res *= wins
	}
	return res, nil
}

func partTwo(lines []string) (int, error) {

	var races = make([]*race, 1)

	times := strings.Split(strings.TrimSpace(strings.Split(lines[0], ": ")[1]), " ")
	time := ""
	for _, v := range times {
		time += strings.TrimSpace(v)
	}
	t, err := strconv.Atoi(time)
	if err != nil {
		return 0, err
	}
	races[0] = &race{
		time: t,
	}

	distances := strings.Split(strings.TrimSpace(strings.Split(lines[1], ": ")[1]), " ")
	dist := ""
	for _, v := range distances {
		dist += strings.TrimSpace(v)
	}
	d, err := strconv.Atoi(dist)
	if err != nil {
		return 0, err
	}
	races[0].distance = d

	return calculate(races)
}
