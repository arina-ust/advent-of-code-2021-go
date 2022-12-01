package day1

import (
	"advent-of-code-go/util"
	"strconv"
)

const day = "day1"

var inputFile string

func Solve(easy bool) (name string, res string, err error) {
	name = day
	setInput(easy)
	lines, err := util.ReadStringList(inputFile)
	if err != nil {
		return
	}

	res, err = partOne(lines)
	// res, err = partTwo(lines)

	return
}

func setInput(easy bool) {
	if easy {
		inputFile = day + "/" + util.InputFileEasy
	} else {
		inputFile = day + "/" + util.InputFileFull
	}
}

type elfInventory struct {
	elf           int
	itemsCalories []string
	sumCalories   int64
}

func partOne(lines []string) (string, error) {
	var inventoryList []elfInventory
	currentElf := elfInventory{
		elf: 1,
	}
	maxCaloriesElf := elfInventory{
		elf:         -1,
		sumCalories: 0,
	}

	for _, line := range lines {
		if len(line) == 0 {
			if currentElf.sumCalories > maxCaloriesElf.sumCalories {
				maxCaloriesElf.elf = currentElf.elf
				maxCaloriesElf.sumCalories = currentElf.sumCalories
				maxCaloriesElf.itemsCalories = currentElf.itemsCalories
			}

			inventoryList = append(inventoryList, currentElf)
			currentElf = elfInventory{
				elf: currentElf.elf + 1,
			}

			continue
		}

		currentElf.itemsCalories = append(currentElf.itemsCalories, line)

		v, err := strconv.Atoi(line)
		if err != nil {
			return "", err
		}
		currentElf.sumCalories += int64(v)
	}

	return strconv.FormatInt(maxCaloriesElf.sumCalories, 10), nil
}
