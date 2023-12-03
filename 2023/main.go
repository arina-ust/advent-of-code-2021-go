package main

import (
	"advent-of-code-go/2023/day3"
	"fmt"
)

func main() {

	fmt.Println("Calculating solution for...")

	day, res, err := day3.Solve(false) //512903 incorrect

	fmt.Println("---> ", day, " <---")
	fmt.Println()

	if err != nil {
		fmt.Printf("Something went wrong: %v", err)
	}
	fmt.Println(res)
}
