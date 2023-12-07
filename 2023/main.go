package main

import (
	"advent-of-code-go/2023/day7"
	"fmt"
)

func main() {

	fmt.Println("Calculating solution for...")

	day, res, err := day7.Solve(true) // 248717914 too low

	fmt.Println("---> ", day, " <---")
	fmt.Println()

	if err != nil {
		fmt.Printf("Something went wrong: %v", err)
	}
	fmt.Println(res)
}
