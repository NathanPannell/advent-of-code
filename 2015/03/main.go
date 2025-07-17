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
	'^': {0, 1},
	'v': {0, -1},
	'<': {-1, 0},
	'>': {1, 0},
}

func part1(input string) int {
	visits := map[coordinates]bool{
		coordinates{0, 0}: true,
	}
	coords := coordinates{0, 0}

	for _, move := range input {
		moveCoords := m[move]
		coords.X += moveCoords.X
		coords.Y += moveCoords.Y

		visits[coords] = true
	}
	return len(visits)
}

func part2(input string) int {
	visits := map[coordinates]bool{
		coordinates{0, 0}: true,
	}
	santaCoords := coordinates{0, 0}
	roboSantaCoords := coordinates{0, 0}

	for i, move := range input {
		// Santa goes first then alternates (even turns)
		var currentSanta *coordinates
		if i%2 == 0 {
			currentSanta = &santaCoords
		} else {
			currentSanta = &roboSantaCoords
		}

		moveCoords := m[move]
		currentSanta.X += moveCoords.X
		currentSanta.Y += moveCoords.Y

		visits[*currentSanta] = true
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
