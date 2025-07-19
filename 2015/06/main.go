package main

import (
	"advent/utils"
	"fmt"
	"regexp"
	"strconv"
)

type coords struct {
	X int
	Y int
}

func apply(lights *[][1000]bool, start coords, end coords, command string) {
	// Use the first 8 characters of the command to determine the type of operation
	if command == "turn on " {
		for x := start.X; x <= end.X; x++ {
			for y := start.Y; y <= end.Y; y++ {
				(*lights)[x][y] = true
			}
		}
	} else if command == "turn off" {
		for x := start.X; x <= end.X; x++ {
			for y := start.Y; y <= end.Y; y++ {
				(*lights)[x][y] = false
			}
		}
	} else {
		for x := start.X; x <= end.X; x++ {
			for y := start.Y; y <= end.Y; y++ {
				(*lights)[x][y] = !(*lights)[x][y]
			}
		}
	}
}

func countLightsOn(lights [][1000]bool) int {
	count := 0
	for x := 0; x < len(lights); x++ {
		for y := 0; y < len(lights[x]); y++ {
			if lights[x][y] {
				count++
			}
		}
	}
	return count
}

func solve(input []string) (int, int) {
	part1 := 0
	part2 := 0

	lights := make([][1000]bool, 1000)
	r := regexp.MustCompile(`(\d*),(\d*) through (\d*),(\d*)$`)

	for _, command := range input {
		start, end := coords{}, coords{}

		matches := r.FindStringSubmatch(command)
		start.X, _ = strconv.Atoi(matches[1])
		start.Y, _ = strconv.Atoi(matches[2])
		end.X, _ = strconv.Atoi(matches[3])
		end.Y, _ = strconv.Atoi(matches[4])

		apply(&lights, start, end, command[:8])
	}

	part1 = countLightsOn(lights)
	return part1, part2
}

func main() {
	input := utils.ReadInput(6)
	part1, part2 := solve(input)

	fmt.Println("-- Day 06 --")
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
