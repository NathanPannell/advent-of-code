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

func apply(lights *[][1000]bool, dynamicLights *[][1000]int, start coords, end coords, command string) {
	// Use the first 8 characters of the command to determine the type of operation
	if command == "turn on " {
		for x := start.X; x <= end.X; x++ {
			for y := start.Y; y <= end.Y; y++ {
				(*lights)[x][y] = true
				(*dynamicLights)[x][y]++
			}
		}
	} else if command == "turn off" {
		for x := start.X; x <= end.X; x++ {
			for y := start.Y; y <= end.Y; y++ {
				(*lights)[x][y] = false
				if (*dynamicLights)[x][y] > 0 {
					(*dynamicLights)[x][y]--
				}
			}
		}
	} else {
		for x := start.X; x <= end.X; x++ {
			for y := start.Y; y <= end.Y; y++ {
				(*lights)[x][y] = !(*lights)[x][y]
				(*dynamicLights)[x][y] += 2
			}
		}
	}
}

func countLightsOn(lights [][1000]bool, dynamicLights [][1000]int) (int, int) {
	count, brightness := 0, 0
	for x := 0; x < len(lights); x++ {
		for y := 0; y < len(lights[x]); y++ {
			if lights[x][y] {
				count++
			}
			brightness += dynamicLights[x][y]
		}
	}
	return count, brightness
}

func solve(input []string) (int, int) {
	part1 := 0
	part2 := 0

	lights := make([][1000]bool, 1000)
	dynamicLights := make([][1000]int, 1000)
	r := regexp.MustCompile(`(\d*),(\d*) through (\d*),(\d*)$`)

	for _, command := range input {
		start, end := coords{}, coords{}

		matches := r.FindStringSubmatch(command)
		start.X, _ = strconv.Atoi(matches[1])
		start.Y, _ = strconv.Atoi(matches[2])
		end.X, _ = strconv.Atoi(matches[3])
		end.Y, _ = strconv.Atoi(matches[4])

		apply(&lights, &dynamicLights, start, end, command[:8])
	}

	part1, part2 = countLightsOn(lights, dynamicLights)
	return part1, part2
}

func main() {
	input := utils.ReadInput(6)
	part1, part2 := solve(input)

	fmt.Println("-- Day 06 --")
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
