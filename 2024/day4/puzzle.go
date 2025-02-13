package day4

import (
	"advent-of-code-go/util"
	"fmt"
)

const day = "day4"

var inputFile string

// Directions represent all possible search directions (including diagonals)
var directions = [][]int{
	{0, 1},  // right
	{1, 0},  // down
	{1, 1},  // down-right
	{1, -1}, // down-left
}

func Solve(easy bool) (name string, res int, err error) {
	name = day
	setInput(easy)
	lines, err := util.ReadMatrixString(inputFile)
	if err != nil {
		return
	}
	if len(lines) == 0 {
		return name, 0, nil
	}

	res, err = partOne(lines)
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

func partOne(lines [][]string) (int, error) {
	xmasCount := 0
	uniqueLocations := make(map[string]bool)

	for i := range lines {
		for j := range lines[i] {
			locations := findXMAS(lines, i, j)
			for _, loc := range locations {
				if !uniqueLocations[loc] {
					uniqueLocations[loc] = true
					xmasCount++
				}
			}
		}
	}

	// fmt.Printf("Total XMAS count: %d\n", xmasCount)
	// fmt.Println("Unique XMAS/SAMX Locations:")
	// for loc := range uniqueLocations {
	// 	fmt.Println(loc)
	// }
	return xmasCount, nil
}

func findXMAS(grid [][]string, startRow, startCol int) []string {
	// Possible words to find
	words := [][]string{
		{"X", "M", "A", "S"}, // XMAS
		{"S", "A", "M", "X"}, // SAMX
	}

	foundLocations := make([]string, 0)

	for _, word := range words {
		for _, dir := range directions {
			if locations := checkWord(grid, startRow, startCol, dir[0], dir[1], word); len(locations) > 0 {
				// Create a unique key based on the letter indices and direction
				uniqueLoc := fmt.Sprintf("%s-%s-%s-%s-%d-%d",
					locations[0], locations[1], locations[2], locations[3],
					dir[0], dir[1])
				foundLocations = append(foundLocations, uniqueLoc)
			}
		}
	}

	return foundLocations
}

func checkWord(grid [][]string, row, col, dRow, dCol int, target []string) []string {
	// Check if starting point is within grid
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
		return nil
	}

	// Check if first letter matches
	if grid[row][col] != target[0] {
		return nil
	}

	// Track current position
	currentRow, currentCol := row, col
	matchedLetters := 1
	matchedPositions := []string{fmt.Sprintf("(%d,%d)", row, col)}

	// Continue searching in the direction
	for matchedLetters < len(target) {
		// Move to next position
		currentRow += dRow
		currentCol += dCol

		// Check if new position is within grid
		if currentRow < 0 || currentRow >= len(grid) ||
			currentCol < 0 || currentCol >= len(grid[0]) {
			return nil
		}

		// Check if next letter matches
		if grid[currentRow][currentCol] == target[matchedLetters] {
			matchedLetters++
			matchedPositions = append(matchedPositions, fmt.Sprintf("(%d,%d)", currentRow, currentCol))
		} else {
			return nil
		}
	}

	// If we've matched all letters, return the matched positions
	return matchedPositions
}

func partTwo(lines [][]string) (int, error) {
	// Find all MAS words
	masWords := findMASWords(lines)

	// Count X-MAS formations
	xmasCount := 0
	uniqueXMAS := make(map[string]bool)

	// fmt.Println("Total MAS words found:", len(masWords))

	// Compare each pair of MAS words
	for i := 0; i < len(masWords); i++ {
		for j := i + 1; j < len(masWords); j++ {
			if isValidXMAS(masWords[i], masWords[j]) {
				// Create a unique key to prevent double-counting
				key := fmt.Sprintf("%s-%s", masWords[i].key, masWords[j].key)
				if !uniqueXMAS[key] {
					uniqueXMAS[key] = true
					xmasCount++

					// Debug print for each valid X-MAS formation
					// fmt.Printf("X-MAS Formation %d:\n", xmasCount)
					// fmt.Printf("  Word 1: M=%v, A=%v, S=%v\n", masWords[i].m, masWords[i].a, masWords[i].s)
					// fmt.Printf("  Word 2: M=%v, A=%v, S=%v\n", masWords[j].m, masWords[j].a, masWords[j].s)
					// fmt.Printf("  Intersection at: %v\n", masWords[i].a)
				}
			}
		}
	}

	return xmasCount, nil
}

// MASWord represents a found MAS word with its positions
type MASWord struct {
	m, a, s [2]int // Integer coordinates
	key     string // Unique identifier for the word
}

func findMASWords(grid [][]string) []MASWord {
	words := []string{"MAS"}
	foundWords := make([]MASWord, 0)

	// All possible search directions
	allDirections := [][]int{
		{0, 1},   // right
		{0, -1},  // left
		{1, 0},   // down
		{-1, 0},  // up
		{1, 1},   // down-right
		{1, -1},  // down-left
		{-1, 1},  // up-right
		{-1, -1}, // up-left
	}

	for i := range grid {
		for j := range grid[i] {
			for _, dir := range allDirections {
				if word := findWord(grid, i, j, dir, words[0]); word != nil {
					foundWords = append(foundWords, *word)
				}
			}
		}
	}

	return foundWords
}

func findWord(grid [][]string, row, col int, dir []int, target string) *MASWord {
	// Ensure starting point is within grid
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
		return nil
	}

	// Track positions of each letter
	positions := make([][2]int, len(target))
	currentRow, currentCol := row, col

	// Check each letter in the word
	for idx, letter := range target {
		// Check grid bounds
		if currentRow < 0 || currentRow >= len(grid) ||
			currentCol < 0 || currentCol >= len(grid[0]) {
			return nil
		}

		// Check letter match
		if string(grid[currentRow][currentCol][0]) != string(letter) {
			return nil
		}

		// Store position
		positions[idx] = [2]int{currentRow, currentCol}

		// Move to next position
		currentRow += dir[0]
		currentCol += dir[1]
	}

	// Create MAS word
	return &MASWord{
		m:   positions[0],
		a:   positions[1],
		s:   positions[2],
		key: fmt.Sprintf("%v-%v-%v-%v", positions[0], positions[1], positions[2], dir),
	}
}

func isValidXMAS(word1, word2 MASWord) bool {
	// Must intersect only at A
	if word1.a != word2.a {
		return false
	}

	// Prevent using the same word twice
	if word1.key == word2.key {
		return false
	}

	// Ensure M and S are on opposite sides of the A
	return (word1.m[0] != word2.m[0] || word1.m[1] != word2.m[1]) &&
		(word1.s[0] != word2.s[0] || word1.s[1] != word2.s[1]) &&
		abs(word1.m[0]-word1.a[0]) == abs(word2.s[0]-word1.a[0]) &&
		abs(word1.m[1]-word1.a[1]) == abs(word2.s[1]-word1.a[1])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
