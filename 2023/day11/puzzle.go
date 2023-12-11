package day11

import (
	"advent-of-code-go/util"
	"fmt"
	"github.com/mowshon/iterium"
	"math"
	"slices"
)

const day = "day11"

var inputFile string

func Solve(easy bool) (name string, res int, err error) {
	name = day
	setInput(easy)
	lines, err := util.ReadMatrixString(inputFile)
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

type pair struct {
	x int
	y int
}

func partOne(lines [][]string) (int, error) {
	//	for _, line := range lines {
	//		fmt.Println(line)
	//	}
	// detect empty lines
	var emptyLineIndices []int
	for j := 0; j < len(lines); j++ {
		countEmpty := 0
		for i := 0; i < len(lines[0]); i++ {
			if lines[j][i] == "." {
				countEmpty++
			}
		}
		if countEmpty == len(lines[0]) {
			emptyLineIndices = append(emptyLineIndices, j)
		}
	}
	fmt.Println("=====", len(lines), " x ", len(lines[0]))
	// add empty lines
	countInserted := 0
	for _, eli := range emptyLineIndices {
		//		[[.#..][....][#...][..#.]]
		empty := make([]string, len(lines[0]))
		for i := 0; i < len(empty); i++ {
			empty[i] = "."
		}
		//		[[.#..][....][....][#...][..#.]]
		lines = slices.Insert(lines, eli+1+countInserted, empty)
		countInserted++
	}
	//	for _, line := range lines {
	//		fmt.Println(line)
	//	}
	fmt.Println("=====", len(lines), " x ", len(lines[0]))
	// detect empty columns
	var emptyColIndices []int
	for i := 0; i < len(lines[0]); i++ {
		countEmpty := 0
		for j := 0; j < len(lines); j++ {
			if lines[j][i] == "." {
				countEmpty++
			}
		}
		if countEmpty == len(lines) {
			emptyColIndices = append(emptyColIndices, i)
		}
	}
	// add empty columns
	countInserted = 0
	for _, eci := range emptyColIndices {
		//		[[.#..][....][#...][..#.]]
		for j := 0; j < len(lines); j++ {
			//		[[.#...][.....][#....][..#..]]
			lines[j] = slices.Insert(lines[j], eci+1+countInserted, ".")
		}
		countInserted++
	}
	//	for _, line := range lines {
	//		fmt.Println(line)
	//	}
	fmt.Println("=====", len(lines), " x ", len(lines[0]))

	// find all pairs
	pairsMap := map[pair]bool{}
	for j := 0; j < len(lines); j++ {
		for i := 0; i < len(lines[0]); i++ {
			if lines[j][i] == "#" {
				p := pair{x: i, y: j}
				if _, ok := pairsMap[p]; !ok {
					pairsMap[p] = true
				}
			}
		}
	}
	//	fmt.Println(pairsMap)

	keys := make([]pair, len(pairsMap))
	i := 0
	for k := range pairsMap {
		keys[i] = k
		i++
	}
	fmt.Println(len(keys))

	pairs, _ := iterium.Combinations(keys, 2).Slice() // [[{} {}] [{}{}] [{}{}]]
	fmt.Println(len(pairs))

	// calculate paths
	// {1,6} -> {5,11} = 9
	// {4,0} -> {9,10} = 15
	// {0,2} -> {12,7} = 17
	// {0,11} -> {5,11} = 5
	sum := 0
	for _, p := range pairs {
		v1 := p[0]
		v2 := p[1]
		dx := math.Abs(float64(v1.x - v2.x))
		dy := math.Abs(float64(v1.y - v2.y))
		sum += int(dx) + int(dy)
	}

	return sum, nil
}

func partTwo(lines [][]string) (int, error) {
	// detect empty lines
	emptyLineIndices := map[int]int{} // index to previous count of empty
	countEmptyLines := 0
	for j := 0; j < len(lines); j++ {
		countEmpty := 0
		for i := 0; i < len(lines[0]); i++ {
			if lines[j][i] == "." {
				countEmpty++
			}
		}
		emptyLineIndices[j] = countEmptyLines
		if countEmpty == len(lines[0]) {
			countEmptyLines++
		}
	}

	// detect empty columns
	emptyColIndices := map[int]int{} // index to previous count of empty
	countEmptyCols := 0
	for i := 0; i < len(lines[0]); i++ {
		countEmpty := 0
		for j := 0; j < len(lines); j++ {
			if lines[j][i] == "." {
				countEmpty++
			}
		}
		emptyColIndices[i] = countEmptyCols
		if countEmpty == len(lines) {
			countEmptyCols++
		}
	}

	// find all pairs
	pairsMap := map[pair]bool{}
	for j := 0; j < len(lines); j++ {
		for i := 0; i < len(lines[0]); i++ {
			if lines[j][i] == "#" {
				p := pair{x: i, y: j}
				if _, ok := pairsMap[p]; !ok {
					pairsMap[p] = true
				}
			}
		}
	}

	keys := make([]pair, len(pairsMap))
	i := 0
	for k := range pairsMap {
		keys[i] = k
		i++
	}
	fmt.Println(len(keys))

	pairs, _ := iterium.Combinations(keys, 2).Slice() // [[{} {}] [{}{}] [{}{}]]
	fmt.Println(len(pairs))

	// calculate paths
	// {1,6} -> {5,11} = 9   --- {1,5} -> {4,9} = 9
	// {4,0} -> {9,10} = 15
	// {0,2} -> {12,7} = 17
	// {0,11} -> {5,11} = 5
	sum := 0
	for _, p := range pairs {
		v1 := p[0]
		v2 := p[1]
		dx := math.Abs(float64(v1.x + emptyColIndices[v1.x] - (v2.x + emptyColIndices[v2.x])))
		dy := math.Abs(float64(v1.y + emptyLineIndices[v1.y] - (v2.y + emptyLineIndices[v2.y])))
		sum += int(dx) + int(dy)
	}

	return sum, nil
}
