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

	if easy {
		res, err = partOne(lines)
	} else {
		return day, "", fmt.Errorf("not solved yet")
	}

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
	fmt.Println(signal)
	fmt.Println(len(signal))

	fmt.Println(parseVersion(signal))
	typeID, _ := parseTypeID(signal)
	fmt.Println(typeID)

	if isLiteral(typeID) {
		literal, _ := parseLiteral(signal[6:])
		fmt.Println(literal)
		fmt.Println(convertToInt(literal))
	} else {
		if isLengthOfSubpackets(rune(signal[6])) {
			lengthSubp, _ := getLengthOfSubpackets(signal[7:])
			fmt.Println(lengthSubp)
			firstLiteral, lengthLiteral, _ := parseLiteral(signal[22+6:]) // 6 is version and type ID of this subpacket
			fmt.Println(firstLiteral)
			fmt.Println(convertToInt(firstLiteral))
			lengthSubp = lengthSubp - lengthLiteral
			secondLiteral, _ := parseLiteral(signal[22+6+lengthLiteral+6:]) // 6 is version and type ID of this subpacket
			fmt.Println(secondLiteral)
			fmt.Println(convertToInt(secondLiteral))
		}
	}

	return "", nil
}

func parseVersion(signal string) (int64, error) {
	return strconv.ParseInt(signal[0:3], 2, 64)
}

func parseTypeID(signal string) (int64, error) {
	return strconv.ParseInt(signal[3:6], 2, 64)
}

func parseLiteral(signalPart string) (string, error) {
	var literal string
	for i := 0; i < len(signalPart); i += 5 {
		val, err := strconv.Atoi(string(signalPart[i]))
		if err != nil {
			return "", err
		}

		literal += signalPart[i+1 : i+5]

		if val == 0 {
			break
		}

	}
	// TODO return also length of literal = i  + necesarry padding accounted for
	return literal, nil
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
