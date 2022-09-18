package day16

import (
	"advent-of-code-2021/util"
	"encoding/hex"
	"fmt"
	"strconv"
)

const day = "day16"

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

func partOne(lines []string) (string, error) {
	line := lines[0]

	bytes, err := hex.DecodeString(line)
	if err != nil {
		return "", err
	}

	var signal string
	for _, b := range bytes {
		signal += fmt.Sprintf("%08b", b)
	}

	res, err := parse(signal, 0, 0, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(int64(res), 10), nil
}

// return version sum
func parse(signal string, versionSum int64, lengthSubP int64, numSubP int64) (int64, error) {
	if len(signal) < 11 {
		return versionSum, nil
	}

	version, err := parseVersion(signal)
	if err != nil {
		return -1, err
	}
	versionSum += version

	typeID, err := parseTypeID(signal)
	if err != nil {
		return -1, err
	}

	if isLiteral(typeID) {
		_, nextIndex, err := parseLiteral(signal[6:])
		if err != nil {
			return -1, err
		}

		if lengthSubP == 0 && numSubP == 0 {
			for nextIndex%4 != 0 {
				nextIndex++
			}
		} else if lengthSubP != 0 && numSubP == 0 {
			lengthSubP -= int64(nextIndex)
		} else if lengthSubP == 0 && numSubP != 0 {
			numSubP--
		} else {
			// ignore
		}

		versionSum, _ = parse(signal[nextIndex:], versionSum, lengthSubP, numSubP)
	} else {
		if isLengthOfSubpackets(rune(signal[6])) {
			calcLengthSubp, _ := getLengthOfSubpackets(signal[7:])
			versionSum, _ = parse(signal[7+15:], versionSum, calcLengthSubp, numSubP)
		} else {
			calcNumSubP, _ := getNumberOfSubpackets(signal[7:])
			versionSum, _ = parse(signal[7+11:], versionSum, lengthSubP, calcNumSubP)
		}
	}
	return versionSum, nil
}

func parseVersion(signal string) (int64, error) {
	return strconv.ParseInt(signal[0:3], 2, 64)
}

func parseTypeID(signal string) (int64, error) {
	return strconv.ParseInt(signal[3:6], 2, 64)
}

func parseLiteral(signalPart string) (string, int, error) {
	var literal string
	var exitIndex int
	for i := 0; i < len(signalPart); i += 5 {
		val, err := strconv.Atoi(string(signalPart[i]))
		if err != nil {
			return "", 0, err
		}

		literal += signalPart[i+1 : i+5]

		if val == 0 {
			exitIndex = 6 + i + 5 // vesion x 3 + typeID x 3 + next 5 bits
			break
		}

	}

	// l, _ := convertToInt(literal)
	// fmt.Printf("Exit index in literal %v was %v \n", l, exitIndex)

	return literal, exitIndex, nil
}

func isLiteral(typeID int64) bool {
	return typeID == 4
}

func convertToInt(val string) (int64, error) {
	return strconv.ParseInt(val, 2, 64)
}

func isLengthOfSubpackets(lengthTypeID rune) bool {
	if lengthTypeID == '0' {
		return true
	}
	return false // ie number of subpackets
}

func getLengthOfSubpackets(signalPart string) (int64, error) {
	return strconv.ParseInt(signalPart[0:15], 2, 64)
}

func getNumberOfSubpackets(signalPart string) (int64, error) {
	return strconv.ParseInt(signalPart[0:11], 2, 64)
}
