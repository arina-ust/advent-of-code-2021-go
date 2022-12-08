package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const InputFileEasy = "easy.txt"
const InputFileFull = "full.txt"

func ReadStringList(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func ReadMatrix(filepath string) ([][]int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var matrix [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := []int{}
		for _, n := range strings.Split(scanner.Text(), "") {
			num, err := strconv.Atoi(n)
			if err != nil {
				return [][]int{}, err
			}
			row = append(row, num)
		}
		matrix = append(matrix, row)
	}

	return matrix, scanner.Err()
}
