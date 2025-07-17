package main

import (
	"2015/utils"
	"fmt"
)

type coordinates struct {
	X int
	Y int
}

var m = map[int32]coordinates{
	'^': coordinates{0, 1},
	'v': coordinates{0, -1},
	'<': coordinates{-1, 0},
	'>': coordinates{1, 0},
}

func part1(input string) int {
	visits := map[string]int{"0,0": 1}
	coords := coordinates{0, 0}

	for _, move := range input {
		moveCoords := m[move]
		coords.X += moveCoords.X
		coords.Y += moveCoords.Y

		position := fmt.Sprintf("%d,%d", coords.X, coords.Y)
		visits[position] += 1
	}
	return len(visits)
}

func part2(input string) int {
	visits := map[string]int{"0,0": 1}
	santaCoords := coordinates{0, 0}
	roboSantaCoords := coordinates{0, 0}

	for i, move := range input {
		moveCoords := m[move]
		var position string

		// Santa goes first then alternates (even turns)
		if i%2 == 0 {
			santaCoords.X += moveCoords.X
			santaCoords.Y += moveCoords.Y
			position = fmt.Sprintf("%d,%d", santaCoords.X, santaCoords.Y)
		} else {
			roboSantaCoords.X += moveCoords.X
			roboSantaCoords.Y += moveCoords.Y
			position = fmt.Sprintf("%d,%d", roboSantaCoords.X, roboSantaCoords.Y)
		}

		visits[position] += 1
	}
	return len(visits)
}

func main() {
	input := utils.ReadInput(3)[0]
	part1 := part1(input)
	part2 := part2(input)

	fmt.Println("-- Day 03 --")
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
