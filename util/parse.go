package util

import (
	"bufio"
	"os"
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
