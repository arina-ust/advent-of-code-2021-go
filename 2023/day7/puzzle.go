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
	if len(m) == 5 { // hand size
		return 1
	}

	for _, v := range m {
		if v == 5 {
			return 7
		} else if v == 4 {
			return 6
		} else if v == 3 {
			if len(m) == 2 {
				return 5
			}
			if len(m) == 3 {
				return 4
			}
		} else if v == 2 {
			if len(m) == 3 {
				return 3
			} else if len(m) == 4 {
				return 2
			}
		}
	}
	return 0 // error
}

var cardStrengths = map[rune]int{'A': 13, 'K': 12, 'Q': 11, 'J': 10, 'T': 9, '9': 8,
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
		h.rank = j+1
		res += h.rank * h.bid
	}
	
	return res, nil
}
