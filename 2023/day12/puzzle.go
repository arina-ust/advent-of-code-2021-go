package day12

import (
	"advent-of-code-go/util"
	"fmt"
	"github.com/mowshon/iterium"
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
	isFull     bool
}

func countArrangements(s string, damaged []int) int {
	knownBrokens := map[int]*broken{} // length to object
	b := &broken{}
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

			knownBrokens[b.length] = b
			b = &broken{}
		}
	}
	fmt.Println(knownBrokens)

	potentialChunks := strings.Split(s, ".")

	damagedObjects := map[int]*broken{}
	for i, d := range damaged {
		damagedObjects[d] = &broken{length: d, startIndex: i}
	}

	count := 0
	for _, pc := range potentialChunks { //???.### -> [???, ###]
		if len(pc) == 0 {
			continue
		}
		if len(pc) == strings.Count(pc, "#") {
			kb := knownBrokens[len(pc)]
			kb.isFull = true
			pc = "" // does it work with non-pointers?
			do := damagedObjects[len(pc)]
			do.isFull = true
			continue
		}
		chars := []string{".", "#"}
		permsI := iterium.Permutations(chars, len(pc))
		permutations, _ := permsI.Slice() // [[] [] []]
		for _, d := range damaged {
			do := damagedObjects[d]
			if do.isFull {
				continue
			}
			for _, p := range permutations {
				if strings.Count(strings.Join(p, ""), "#") == d {
					count++
				}
			}
		}
	}

	fmt.Println(count)
	return 1
}
