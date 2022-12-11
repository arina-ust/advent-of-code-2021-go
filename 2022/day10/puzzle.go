package day10

import (
	"advent-of-code-go/util"
	"fmt"
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

	// res, err = partOne(lines)
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

type cpu struct {
	register                register
	clock                   int
	signalStrengthsPerCycle map[int]int

	crt    [][]string
	sprite *sprite
}

type register int

type sprite struct {
	i, j, k int
	y       int
}

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
	cpu.drawPixel()
}

func (cpu *cpu) add(x int) {
	for i := 0; i < 2; i++ {
		cpu.clock++
		cpu.checkSignalStrength()
		cpu.drawPixel()
	}
	cpu.register += register(x)
	cpu.moveSprite()
}

func (cpu *cpu) checkSignalStrength() {
	if cpu.clock == 20 || cpu.clock == 60 || cpu.clock == 100 || cpu.clock == 140 ||
		cpu.clock == 180 || cpu.clock == 220 {
		signalStrength := cpu.register * register(cpu.clock)
		cpu.signalStrengthsPerCycle[cpu.clock] = int(signalStrength)
	}
}

func (cpu *cpu) drawPixel() {
	row, col := 0, 0
	if cpu.clock >= 1 && cpu.clock <= 40 {
		row = 0
		col = cpu.clock - 1
	} else if cpu.clock >= 41 && cpu.clock <= 80 {
		row = 1
		col = cpu.clock - 41
	} else if cpu.clock >= 81 && cpu.clock <= 120 {
		row = 2
		col = cpu.clock - 81
	} else if cpu.clock >= 121 && cpu.clock <= 160 {
		row = 3
		col = cpu.clock - 121
	} else if cpu.clock >= 161 && cpu.clock <= 200 {
		row = 4
		col = cpu.clock - 161
	} else if cpu.clock >= 201 && cpu.clock <= 240 {
		row = 5
		col = cpu.clock - 201
	}

	shouldLight := (cpu.sprite.i == col || cpu.sprite.j == col || cpu.sprite.k == col)
	if shouldLight {
		cpu.crt[row][col] = "#"
	} else {
		cpu.crt[row][col] = "."
	}
}

func (cpu *cpu) moveSprite() {
	cpu.sprite.i = int(cpu.register) - 1
	cpu.sprite.j = int(cpu.register)
	cpu.sprite.k = int(cpu.register) + 1
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

func partTwo(lines []string) (string, error) {
	crt := make([][]string, 6)
	for i := 0; i < 6; i++ {
		row := make([]string, 40)
		crt[i] = row
	}

	cpu := &cpu{
		register:                1,
		clock:                   0,
		signalStrengthsPerCycle: map[int]int{},
		crt:                     crt,
		sprite:                  &sprite{i: 0, j: 1, k: 2, y: 0},
	}

	for _, line := range lines {
		cpu.performOp(line)
	}

	for _, row := range cpu.crt {
		fmt.Println(row)
	}

	return "", nil
}
