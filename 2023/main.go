package main

import (
	"advent-of-code-go/2023/day7"
	"fmt"
)

func main() {

	fmt.Println("Calculating solution for...")

	day, res, err := day7.Solve(false) // 248717914 too low, 249623465 

	fmt.Println("---> ", day, " <---")
	fmt.Println()

	if err != nil {
		fmt.Printf("Something went wrong: %v", err)
	}
	fmt.Println(res)
}
