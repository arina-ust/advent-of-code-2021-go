package main

import (
	"advent-of-code-go/2023/day11"
	"fmt"
)

func main() {

	fmt.Println("Calculating solution for...")

	day, res, err := day11.Solve(true)

	fmt.Println("---> ", day, " <---")
	fmt.Println()

	if err != nil {
		fmt.Printf("Something went wrong: %v", err)
	}
	fmt.Println(res)
}
