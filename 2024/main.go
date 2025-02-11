package main

import (
	"advent-of-code-go/2024/day3"
	"fmt"
)

func main() {
	fmt.Println("Calculating solution for...")

	day, res, err := day3.Solve(false)

	fmt.Println("---> ", day, " <---")
	fmt.Println()

	if err != nil {
		fmt.Printf("Something went wrong: %v", err)
	}
	fmt.Println(res)
}
