package main

import (
	"advent-of-code-go/2023/day10"
	"fmt"
)

func main() {

	fmt.Println("Calculating solution for...")

	day, res, err := day10.Solve(true) // 329 too high

	fmt.Println("---> ", day, " <---")
	fmt.Println()

	if err != nil {
		fmt.Printf("Something went wrong: %v", err)
	}
	fmt.Println(res)
}
