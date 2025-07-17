package main

import (
	"2015/utils"
	"fmt"
)

var m = map[int32]int{
	'(': 1,
	')': -1,
}

func part1(input string) int {
	floor := 0
	for _, char := range input {
		floor += m[char]
	}
	return floor
}

func part2(input string) int {
	floor := 0
	for position, char := range input {
		floor += m[char]
		if floor == -1 {
			return position + 1
		}
	}

	return -1
}

func main() {
	input := utils.ReadInput(1)[0]

	fmt.Println("-- Day 01 --")
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}
