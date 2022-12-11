package day10

import (
	"advent-of-code-go/util"
	"strconv"
	"strings"
)

const day = "day10"

var inputFile string

func Solve(easy bool) (name string, res string, err error) {
	name = day
	setInput(easy)
	lines, err := util.ReadStringList(inputFile)
	if err != nil {
		return
	}

	res, err = partOne(lines)
	// res, err = partOne(lines, 10)

	return
}

func setInput(easy bool) {
	if easy {
		inputFile = day + "/" + util.InputFileEasy
	} else {
		inputFile = day + "/" + util.InputFileFull
	}
}

type cpu struct {
	register                register
	clock                   int
	signalStrengthsPerCycle map[int]int
}

type register int

func (cpu *cpu) performOp(s string) error {
	if s == "noop" {
		cpu.noOp()
		return nil
	}

	addOp := strings.Split(s, " ")
	x, err := strconv.Atoi(addOp[1])
	if err != nil {
		return err
	}

	cpu.add(x)

	return nil
}

func (cpu *cpu) noOp() {
	cpu.clock++
	cpu.checkSignalStrength()
}

func (cpu *cpu) add(x int) {
	for i := 0; i < 2; i++ {
		cpu.clock++
		cpu.checkSignalStrength()
	}
	cpu.register += register(x)
}

func (cpu *cpu) checkSignalStrength() {
	if cpu.clock == 20 || cpu.clock == 60 || cpu.clock == 100 || cpu.clock == 140 ||
		cpu.clock == 180 || cpu.clock == 220 {
		signalStrength := cpu.register * register(cpu.clock)
		cpu.signalStrengthsPerCycle[cpu.clock] = int(signalStrength)
	}
}

func partOne(lines []string) (string, error) {
	cpu := &cpu{
		register:                1,
		clock:                   0,
		signalStrengthsPerCycle: map[int]int{},
	}

	for _, line := range lines {
		cpu.performOp(line)
	}

	sum := 0
	for _, v := range cpu.signalStrengthsPerCycle {
		sum += v
	}

	return strconv.Itoa(sum), nil
}
