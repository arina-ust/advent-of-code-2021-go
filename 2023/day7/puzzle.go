package day7

import (
	"advent-of-code-go/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const day = "day7"

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

type hand struct {
	cards        string
	bid          int
	rank         int
	typeStrength int
}

func (h *hand) getStrength() int {
	m := map[rune]int{}
	for _, card := range h.cards {
		m[card] += 1
	}

	// five of kind = 7
	// four of kind = 6
	// full house = 5
	// three of kind = 4
	// two pair = 3
	// one pair = 2
	// high card = 1

	if m['J'] == 5 {
		// JJJJJ
		return 7 // five of kind
	}
	if m['J'] == 4 {
		// JJJJK
		return 7 // five of kind
	}
	if m['J'] == 3 {
		if len(m) == 2 {
			// JJJKK
			return 7 // five of kind
		} else if len(m) == 3 {
			// JJJKQ
			return 6 // four of kind
		}
	}
	if m['J'] == 2 {
		if len(m) == 2 {
			// JJKKK
			return 7 // five of kind
		} else if len(m) == 3 {
			// JJKKQ
			return 6 // four of kind
		} else if len(m) == 4 {
			// JJKTQ
			return 4 // three of kind
		}
	}
	if m['J'] == 1 {
		if len(m) == 2 {
			// JKKKK
			return 7 // five of kind
		} else if len(m) == 3 {
			foundThree := false
			for _, v := range m {
				if v == 3 {
					foundThree = true
				}
			}
			if foundThree {
				// JKKKQ
				return 6 // four of kind
			} else {
				// JKKQQ
				return 5 // full house
			}
		} else if len(m) == 4 {
			// JKKTQ
			return 4 // three of kind
		} else if len(m) == 5 {
			// J2KTQ
			return 2 // one pair
		}
	}

	if len(m) == 5 { // hand size
		// 23456
		return 1 // high card
	}

	for _, v := range m {
		if v == 5 {
			// AAAAA
			return 7 // five of kind
		} else if v == 4 {
			// AA8AA
			return 6 // four of kind
		} else if v == 3 {
			if len(m) == 2 {
				// 23332
				return 5 // full house
			}
			if len(m) == 3 {
				// TTT98
				return 4 // three of kind
			}
		} else if v == 2 {
			if len(m) == 3 {
				// 23432
				return 3 // two pair
			} else if len(m) == 4 {
				// A23A4
				return 2 // one pair
			}
		}
	}
	return 0 // error
}

var cardStrengths = map[rune]int{'A': 13, 'K': 12, 'Q': 11, 'J': 0, 'T': 9, '9': 8,
	'8': 7, '7': 6, '6': 5, '5': 4, '4': 3, '3': 2, '2': 1}

func (h *hand) compareLess(other *hand) bool {
	if h.typeStrength > other.typeStrength {
		return false
	} else if h.typeStrength < other.typeStrength {
		return true
	}
	// equal strength, compare by card
	for i := 0; i < 5; i++ {
		rh := rune(h.cards[i])
		ro := rune(other.cards[i])
		if cardStrengths[rh] > cardStrengths[ro] {
			//			fmt.Printf("rh %s turned out bigger than ro %s with values %v %v\n", string(rh), string(ro), cardStrengths[rh], cardStrengths[ro])
			return false
		} else if cardStrengths[rh] < cardStrengths[ro] {
			//			fmt.Printf("rh %s turned out smaller than ro %s with values %v %v\n", string(rh), string(ro), cardStrengths[rh], cardStrengths[ro])
			return true
		}
		//		fmt.Printf("rh %s was the same as ro %s with values %v %v\n", string(rh), string(ro), cardStrengths[rh], cardStrengths[ro])
	}
	fmt.Println("shouldn't be here")
	return false
}

func partOne(lines []string) (int, error) {
	hands := make([]*hand, len(lines))
	for i, line := range lines {
		input := strings.Split(line, " ")
		bid, _ := strconv.Atoi(input[1])
		hand := &hand{
			cards: input[0],
			bid:   bid,
		}
		strgth := hand.getStrength()
		hand.typeStrength = strgth
		hands[i] = hand
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].compareLess(hands[j])
	})

	res := 0
	for j, h := range hands {
		h.rank = j + 1
		res += h.rank * h.bid
	}

	return res, nil
}
