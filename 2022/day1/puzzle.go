package day1

import (
	"advent-of-code-go/util"
	"sort"
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

	lines = append(lines, "")

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

	// -> part 1 <-
	// return strconv.FormatInt(maxCaloriesElf.sumCalories, 10), nil

	// -> part 2 <-
	sort.SliceStable(inventoryList, func(i, j int) bool { return inventoryList[i].sumCalories > inventoryList[j].sumCalories })

	return strconv.FormatInt(inventoryList[0].sumCalories+inventoryList[1].sumCalories+inventoryList[2].sumCalories, 10), nil
}
